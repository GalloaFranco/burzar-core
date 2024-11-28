package main

import (
	"os"

	"github.com/spf13/viper"
)

func LoadConfigs(env string) error {
	viper.SetConfigName(env)
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	return viper.ReadInConfig()
}

func GetEnvironment() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return env
}
