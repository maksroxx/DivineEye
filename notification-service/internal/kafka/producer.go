package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/maksroxx/DivineEye/notification-service/internal/logger"
	"go.uber.org/zap"
)

type Producerer interface {
	PublishAlertTriggered(alertID string) error
	Close() error
}

type Producer struct {
	producer sarama.SyncProducer
	topic    string
	log      logger.Logger
}

func NewProducer(brokers []string, topic string, log logger.Logger) (*Producer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = 3
	cfg.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka producer: %w", err)
	}
	return &Producer{producer: producer, topic: topic, log: log}, nil
}

func (p *Producer) PublishAlertTriggered(alertID string) error {
	payload := map[string]any{
		"alert_id": alertID,
		"event":    "triggered",
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(data),
	}
	_, _, err = p.producer.SendMessage(msg)
	p.log.Info("[Producer.go] Publish alert triggered", zap.Any("alert_id", alertID))
	return err
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
