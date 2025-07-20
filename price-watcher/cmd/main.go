package main

import (
	"context"

	"github.com/maksroxx/DivineEye/price-watcher/internal/config"
	"github.com/maksroxx/DivineEye/price-watcher/internal/kafka"
	"github.com/maksroxx/DivineEye/price-watcher/internal/logger"
	"github.com/maksroxx/DivineEye/price-watcher/internal/watcher"
)

func main() {
	log := logger.NewZapLogger()
	defer log.Sync()

	cfg, err := config.Load("./price-watcher/config/config.yaml")
	if err != nil {
		panic(err)
	}

	producer, err := kafka.NewProducer(kafka.ConfigProducer{Brokers: cfg.Kafka.Brokers, Topics: cfg.Kafka.Topic})
	if err != nil {
		panic(err)
	}

	w := watcher.NewWatcher(producer, cfg.Binance.Symbols, log)
	log.Info("Price Watcher running...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = w.Run(ctx)
	if err != nil {
		panic(err)
	}
}
