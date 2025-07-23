package kafka

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/maksroxx/DivineEye/alert-service/internal/logger"
	"go.uber.org/zap"
)

type Producerer interface {
	PublishAlertCreated(alertId, userId, coin, direction string, price float64) error
	PublishAlertDeleted(alertId, userId string) error
	Close() error
}

type Producer struct {
	producer sarama.SyncProducer
	topics   string
	log      logger.Logger
}

func NewProducer(cfg ConfigProducer, log logger.Logger) (*Producer, error) {
	producer, err := sarama.NewSyncProducer(cfg.Brokers, NewProducerConfig())
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer: producer,
		topics:   cfg.Topics[0],
		log:      log,
	}, nil
}

func (p *Producer) PublishAlertCreated(alertId, userId, coin, direction string, price float64) error {
	p.log.Info("PublishAlertCreated", zap.Any("direction", direction))
	payload := map[string]any{
		"alert_id":  alertId,
		"user_id":   userId,
		"coin":      coin,
		"direction": direction,
		"price":     price,
		"type":      "alert_created",
	}
	data, _ := json.Marshal(payload)
	headers := []sarama.RecordHeader{
		{Key: []byte("User-ID"), Value: []byte(userId)},
	}

	msg := &sarama.ProducerMessage{
		Headers: headers,
		Topic:   p.topics,
		Value:   sarama.ByteEncoder(data),
	}

	p.log.Info("publishing alert", zap.String("number", alertId))
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
		Topic:   p.topics,
		Value:   sarama.ByteEncoder(data),
	}

	_, _, err := p.producer.SendMessage(msg)
	return err
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
