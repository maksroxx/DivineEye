package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/maksroxx/DivineEye/alert-service/internal/alert"
	"github.com/maksroxx/DivineEye/alert-service/internal/config"
	"github.com/maksroxx/DivineEye/alert-service/internal/kafka"
	"github.com/maksroxx/DivineEye/alert-service/internal/logger"
	"github.com/maksroxx/DivineEye/alert-service/internal/repository"
	pb "github.com/maksroxx/DivineEye/alert-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.Load("./alert-service/config/config.yaml")
	if err != nil {
		panic(err)
	}
	log := logger.NewZapLogger()
	defer log.Sync()

	db, _ := sql.Open("postgres", cfg.Database.Dsn)
	producer, err := kafka.NewProducer(kafka.ConfigProducer{Brokers: cfg.Kafka.Brokers, Topics: cfg.Kafka.Topic}, log)
	if err != nil {
		panic(err)
	}
	repo := repository.NewPostgresRepo(db)
	svc := alert.NewService(repo, producer, log)

	consumerGroup := kafka.NewTriggeredHandler(repo)
	kafka.StartConsumerGroup(ctx, cfg.Kafka.Brokers, cfg.Kafka.Group, cfg.Kafka.Topic[0:], consumerGroup, log)

	grpcServcer := grpc.NewServer()
	pb.RegisterAlertServiceServer(grpcServcer, alert.NewGRPCServer(svc))

	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GRPCPort))
	log.Info("Alert Serivce running", zap.Int("port", cfg.Server.GRPCPort))
	grpcServcer.Serve(lis)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	log.Info("🧹 Shutting down gracefully...")
	time.Sleep(2 * time.Second)
}
