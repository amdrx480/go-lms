package utils

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading configuration file: %s\n", err)
	}

	return viper.GetString(key)
}

func GetConfigDuration(key string) time.Duration {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading configuration file: %s\n", err)
	}

	value := viper.GetString(key)
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Fatalf("error parsing duration for key %s: %v", key, err)
	}
	return duration
}
