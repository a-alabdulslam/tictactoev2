package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"tictactoe/pkg/logger"
)

func HealthCheck(c *gin.Context) {
	mylogger := logger.Logger(logrus.ErrorLevel)
	mylogger.Warnf("error to read config, %v", "fdfd")
	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})

}
