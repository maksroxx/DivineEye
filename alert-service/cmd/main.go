package main

import (
	"database/sql"
	"fmt"
	"net"

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
	cfg, err := config.Load("./alert-service/config/config.yaml")
	if err != nil {
		panic(err)
	}
	log := logger.NewZapLogger()
	defer log.Sync()

	db, _ := sql.Open("postgres", cfg.Database.Dsn)
	producer, err := kafka.NewProducer(kafka.ConfigProducer{Brokers: cfg.Kafka.Brokers, Topics: cfg.Kafka.Topic})
	if err != nil {
		panic(err)
	}
	repo := repository.NewPostgresRepo(db)
	svc := alert.NewService(repo, producer, log)

	grpcServcer := grpc.NewServer()
	pb.RegisterAlertServiceServer(grpcServcer, alert.NewGRPCServer(svc))

	lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GRPCPort))
	log.Info("Alert Serivce running", zap.Int("port", cfg.Server.GRPCPort))
	grpcServcer.Serve(lis)
}
