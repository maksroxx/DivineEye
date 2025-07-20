package kafka

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

type Producerer interface {
	PublishAlertCreated(alertId, userId string) error
	PublishAlertDeleted(alertId, userId string) error
	Close() error
}

type Producer struct {
	producer sarama.SyncProducer
	topics   []string
}

func NewProducer(cfg ConfigProducer) (*Producer, error) {
	producer, err := sarama.NewSyncProducer(cfg.Brokers, NewProducerConfig())
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer: producer,
		topics:   cfg.Topics,
	}, nil
}

func (p *Producer) PublishAlertCreated(alertId, userId string) error {
	payload := map[string]string{
		"alert_id": alertId,
		"user_id":  userId,
		"type":     "alert_created",
	}
	data, _ := json.Marshal(payload)
	headers := []sarama.RecordHeader{
		{Key: []byte("User-ID"), Value: []byte(userId)},
	}

	msg := &sarama.ProducerMessage{
		Headers: headers,
		Topic:   p.topics[0],
		Value:   sarama.ByteEncoder(data),
	}

	_, _, err := p.producer.SendMessage(msg)
	return err
}

func (p *Producer) PublishAlertDeleted(alertId, userId string) error {
	payload := map[string]string{
		"alert_id": alertId,
		"user_id":  userId,
		"type":     "alert_deleted",
	}

	data, _ := json.Marshal(payload)
	headers := []sarama.RecordHeader{
		{Key: []byte("User-ID"), Value: []byte(userId)},
	}

	msg := &sarama.ProducerMessage{
		Headers: headers,
		Topic:   p.topics[0],
		Value:   sarama.ByteEncoder(data),
	}

	_, _, err := p.producer.SendMessage(msg)
	return err
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
