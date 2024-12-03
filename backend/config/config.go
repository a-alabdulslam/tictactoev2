package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"tictactoe/pkg/logger"
)

type Configuration struct {
	Database DatabaseConfiguration
}

// SetupConfig configuration
func SetupConfig() error {
	var configuration *Configuration
	newLogger := logger.Logger(logrus.ErrorLevel)

	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		newLogger.Errorf("error to read config, %v", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		newLogger.Errorf("error to decode, %v", err)
		return err
	}

	return nil
}
