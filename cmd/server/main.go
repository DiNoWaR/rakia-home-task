package main

import (
	"context"
	"fmt"
	"github.com/dinowar/rakia-home-task/internal/pkg/config"
	"github.com/dinowar/rakia-home-task/internal/pkg/server"
	"github.com/dinowar/rakia-home-task/internal/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ctx := context.Background()
	serviceConfig := &config.ServiceConfig{}
	if configErr := envconfig.Process(ctx, serviceConfig); configErr != nil {
		logger.Fatal("failed to init config", zap.Error(configErr))
	}

	logService := service.NewLogService(logger)
	repService := service.NewRepositoryService(logService)
	appServer := server.NewAppServer(repService, logService)

	router := gin.Default()

	router.POST("/posts", appServer.CreatePost)
	router.GET("/posts/:id", appServer.GetPost)
	router.GET("/posts", appServer.GetPosts)
	router.PUT("/posts/:id", appServer.UpdatePost)
	router.DELETE("/posts/:id", appServer.DeletePost)

	logger.Info(fmt.Sprintf("service starting on port: %s", serviceConfig.ServicePort))
	serverStartErr := router.Run(fmt.Sprintf("%s:%s", serviceConfig.ServiceHost, serviceConfig.ServicePort))
	if serverStartErr != nil {
		logger.Fatal("failed to start server", zap.Error(serverStartErr))
	}
}
