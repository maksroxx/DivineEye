package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/maksroxx/DivineEye/auth-service/internal/auth"
	"github.com/maksroxx/DivineEye/auth-service/internal/config"
	"github.com/maksroxx/DivineEye/auth-service/internal/jwt"
	"github.com/maksroxx/DivineEye/auth-service/internal/logger"
	"github.com/maksroxx/DivineEye/auth-service/internal/repository"
	pb "github.com/maksroxx/DivineEye/auth-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	log := logger.NewZapLogger()
	defer log.Sync()
	cfg, err := config.Load("./auth-service/config/config.yaml")
	if err != nil {
		log.Error("failed to load config", zap.Error(err))
	}

	db, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		log.Error("failed to connect db", zap.Error(err))
	}

	jwt.Init(cfg.JWT.Secret)

	userRepo := repository.NewPostgresRepo(db)
	authSvc := auth.NewAuthService(userRepo, log)
	grpcServer := auth.NewGRPCServer(authSvc)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GRPCPort))
	if err != nil {
		log.Error("failed to listen", zap.Error(err))
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcSrv, grpcServer)
	log.Info("Auth Service running", zap.Int("port", cfg.Server.GRPCPort))
	grpcSrv.Serve(lis)
}
