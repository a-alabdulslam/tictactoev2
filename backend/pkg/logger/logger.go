package logger

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *logrus.Logger
)

func initializeLogger(level logrus.Level) *logrus.Logger {

	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: level,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02T15:04:05",
			LogFormat:       "[%lvl%] %time% - %msg%\n",
		},
	}

	return log
}

func Logger(level logrus.Level) *logrus.Logger {
	once.Do(func() {
		logger = initializeLogger(level)
	})
	return logger
}
