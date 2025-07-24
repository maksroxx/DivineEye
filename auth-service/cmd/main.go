package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

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
		panic(err)
	}

	db, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		panic(err)
	}

	jwt.Init(cfg.JWT.Secret)

	userRepo := repository.NewPostgresRepo(db)
	authSvc := auth.NewAuthService(userRepo, log)
	grpcServer := auth.NewGRPCServer(authSvc)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GRPCPort))
	if err != nil {
		log.Error("Failed to listen", zap.Error(err))
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcSrv, grpcServer)
	log.Info("Auth Service running", zap.Int("port", cfg.Server.GRPCPort))
	grpcSrv.Serve(lis)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	log.Info("🧹 Shutting down gracefully...")
	time.Sleep(2 * time.Second)
}
