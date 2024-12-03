package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"tictactoe/config"
	"tictactoe/pkg/go-api/handlers"
	"tictactoe/pkg/logger"
)

type server struct {
	handlers.IExampleHandler
}

func initServer() *server {
	err := config.SetupConfig()
	myLogger := logger.Logger(logrus.ErrorLevel)
	if err == nil {
		config.DbConfiguration()
		myLogger.Error(config.DbConfiguration())
	}
	example := handlers.NewExampleHandler()
	return &server{example}
}

func main() {
	gin.SetMode("debug")
	myServer := initServer()
	r := gin.Default()
	r.GET("/health", handlers.HealthCheck)
	r.GET("/hello/:name", myServer.IExampleHandler.HelloName)

	r.Run()
}
