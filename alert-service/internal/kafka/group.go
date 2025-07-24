package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/maksroxx/DivineEye/alert-service/internal/logger"
	"go.uber.org/zap"
)

func StartConsumerGroup(ctx context.Context, brokers []string, groupID string, topics []string, handler sarama.ConsumerGroupHandler, log logger.Logger) {
	config := NewConsumerConfig()

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Error("[Group.go] Failed to create consumer group: %v", zap.Error(err))
	}

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, topics, handler); err != nil {
				log.Info("[Group.go] Error consuming kafka: %v", zap.Error(err))
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
}
