package kafka

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

type Producerer interface {
	PublishPrice(symbol string, price float64) error
	Close() error
}

type Producer struct {
	Producer sarama.SyncProducer
	Topic    string
}

func NewProducer(cfg ConfigProducer) (*Producer, error) {
	producer, err := sarama.NewSyncProducer(cfg.Brokers, NewProducerConfig())
	if err != nil {
		return nil, err
	}
	return &Producer{
		Producer: producer,
		Topic:    cfg.Topics[0],
	}, nil
}

func (p *Producer) PublishPrice(symbol string, price float64) error {
	msgData := map[string]any{
		"symbol": symbol,
		"price":  price,
	}

	data, _ := json.Marshal(msgData)

	msg := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err := p.Producer.SendMessage(msg)
	return err
}

func (p *Producer) Close() error {
	return p.Producer.Close()
}
