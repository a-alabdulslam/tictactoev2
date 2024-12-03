package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
}

func DbConfiguration() string {
	dbName := viper.GetString("DB_NAME")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")

	dbDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	return dbDSN
}
