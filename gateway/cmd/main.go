package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maksroxx/DivineEye/gateway/internal/auth"
	"github.com/maksroxx/DivineEye/gateway/internal/config"
	"github.com/maksroxx/DivineEye/gateway/internal/grpcclient"
	"github.com/maksroxx/DivineEye/gateway/internal/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load("./gateway/config/config.yaml")
	if err != nil {
		panic(err)
	}
	log := logger.NewZapLogger()
	defer log.Sync()

	authGrpc := grpcclient.NewAuthClient(cfg.Services.Auth)
	authService := auth.NewAuthService(authGrpc)
	authHandler := auth.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/auth/login", authHandler.Login)
	r.POST("/auth/register", authHandler.Register)

	protected := r.Group("/api", auth.JWTAuth(authService))
	protected.GET("/ping", func(ctx *gin.Context) {
		userId := ctx.GetString("user_id")
		ctx.JSON(200, gin.H{"user_id": userId})
	})

	addr := fmt.Sprintf(":%d", cfg.Server.Address)
	log.Info("Api Gateway running", zap.String("addr", addr))
	r.Run(addr)
}
