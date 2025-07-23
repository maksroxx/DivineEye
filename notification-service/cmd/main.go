package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/maksroxx/DivineEye/notification-service/internal/config"
	"github.com/maksroxx/DivineEye/notification-service/internal/consumer"
	"github.com/maksroxx/DivineEye/notification-service/internal/fcm"
	"github.com/maksroxx/DivineEye/notification-service/internal/kafka"
	"github.com/maksroxx/DivineEye/notification-service/internal/logger"
	"github.com/maksroxx/DivineEye/notification-service/internal/repository"
	"github.com/maksroxx/DivineEye/notification-service/internal/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logg := logger.NewZapLogger()
	defer logg.Sync()

	cfg, err := config.Load("./notification-service/config/config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", cfg.Database.Dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewPstgresRepo(db)

	fcmSender, err := fcm.NewSender(ctx, cfg.Fcm.ServerKey)
	if err != nil {
		panic(err)
	}

	notifier := service.NewNotifi(repo, fcmSender, logg)

	kafkaTopics := []string{cfg.Kafka.AlertTopic, cfg.Kafka.PricesTopic}
	consumerGroup := consumer.NewHandler(repo, notifier, logg)

	go kafka.StartConsumerGroup(ctx, cfg.Kafka.Brokers, "notification-consumer", kafkaTopics, consumerGroup, logg)

	logg.Info("notification-service started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	logg.Info("🧹 shutting down gracefully...")
	time.Sleep(2 * time.Second)
}
