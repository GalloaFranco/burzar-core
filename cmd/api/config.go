package main

import "github.com/spf13/viper"

func LoadConfigs() error {
	viper.SetConfigName("local")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	return viper.ReadInConfig()
}
