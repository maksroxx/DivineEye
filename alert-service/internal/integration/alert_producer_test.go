package integration

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/maksroxx/DivineEye/alert-service/internal/alert"
	"github.com/maksroxx/DivineEye/alert-service/internal/kafka"
	"github.com/maksroxx/DivineEye/alert-service/internal/logger"
	"github.com/maksroxx/DivineEye/alert-service/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAlertService_KafkaPublish(t *testing.T) {
	brokers := []string{"localhost:9092"}
	topic := "alerts_test"

	producer, err := kafka.NewProducer(kafka.ConfigProducer{
		Brokers: brokers,
		Topics:  []string{topic},
	})
	require.NoError(t, err)
	defer producer.Close()

	db := setupTestDB(t)
	repo := repository.NewPostgresRepo(db)

	log := logger.NewZapLogger()
	defer log.Sync()

	svc := alert.NewService(repo, producer, log)

	userID := "11111111-1111-1111-1111-111111111111"
	_, err = svc.Create(context.Background(), userID, "BTC", 67000)
	require.NoError(t, err)

	// consumer
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
	//

	timeout := time.After(10 * time.Second)
	var found bool

	for !found {
		select {
		case msg := <-partitionConsumer.Messages():
			val := string(msg.Value)
			log.Info("[TEST KAFKA]", zap.String("received", val))
			if strings.Contains(val, userID) && strings.Contains(val, "alert_id") {
				assert.Contains(t, val, "alert_id")
				assert.Contains(t, val, "user_id")
				found = true
			}
		case <-timeout:
			t.Fatal("timeout waiting for alert_created in kafka")
		}
	}
}
