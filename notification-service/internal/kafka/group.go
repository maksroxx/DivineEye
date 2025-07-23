package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/maksroxx/DivineEye/notification-service/internal/logger"
	"go.uber.org/zap"
)

func StartConsumerGroup(ctx context.Context, brokers []string, groupID string, topics []string, handler sarama.ConsumerGroupHandler, log logger.Logger) {
	config := NewConsumerConfig()

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Error("failed to create consumer group: %v", zap.Error(err))
	}

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, topics, handler); err != nil {
				log.Info("error consuming kafka: %v", zap.Error(err))
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
}
