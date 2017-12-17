package main

import (
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

func LoadConfig() (error) {
	log.SetLevel(log.InfoLevel)
	viper.SetDefault("debug", false)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		switch t := err.(type) {
		default:
			return t
		case viper.ConfigFileNotFoundError:
			log.Info("No config file found, using defaults")
		}
	}
	return nil
}
