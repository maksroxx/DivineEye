package integration

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/maksroxx/DivineEye/price-watcher/internal/kafka"
	"github.com/maksroxx/DivineEye/price-watcher/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestMockWatcher_PublishesToKafka(t *testing.T) {
	brokers := []string{"localhost:9092"}
	topic := "price_updates_test"

	producer, err := kafka.NewProducer(kafka.ConfigProducer{
		Brokers: brokers,
		Topics:  []string{topic},
	})
	require.NoError(t, err)
	defer producer.Close()

	log := logger.NewZapLogger()
	defer log.Sync()

	watcher := &MockWatcher{producer: producer}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_ = watcher.Run(ctx)
	}()

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_8_0_0

	client, err := sarama.NewClient(brokers, config)
	require.NoError(t, err)
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	require.NoError(t, err)
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	require.NoError(t, err)
	defer partitionConsumer.Close()

	timeout := time.After(10 * time.Second)
	var found bool

	for !found {
		select {
		case msg := <-partitionConsumer.Messages():
			val := string(msg.Value)
			if strings.Contains(val, "btc") {
				log.Info("[TEST KAFKA]", zap.String("received", val))
				assert.Contains(t, val, "symbol")
				assert.Contains(t, val, "price")
				found = true
			}
		case <-timeout:
			t.Fatal("timeout waiting for price update in kafka")
		}
	}
}
