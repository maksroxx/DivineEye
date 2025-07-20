package integration

import (
	"context"

	"github.com/maksroxx/DivineEye/price-watcher/internal/kafka"
)

type MockWatcher struct {
	producer kafka.Producerer
}

func (m *MockWatcher) Run(ctx context.Context) error {
	_ = m.producer.PublishPrice("btcusdt", 67890.0)
	return nil
}
