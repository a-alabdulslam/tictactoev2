package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"tictactoe/config"
	"tictactoe/pkg/game"
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
	hub := game.NewHub()
	go hub.Run()

	r := gin.Default()
	r.GET("/health", handlers.HealthCheck)
	r.GET("/hello/:name", myServer.IExampleHandler.HelloName)

	r.GET("/ws/", func(c *gin.Context) {
		roomId := c.Param("roomId")
		name := c.Request.Header.Get("name")
		// validate name/roomId and return 400 if invalid or empty

		game.ServeWS(c, name, roomId, hub)

	})
	r.Run()
}
