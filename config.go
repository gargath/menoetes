package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig() error {
	log.SetLevel(log.InfoLevel)
	viper.SetDefault("debug", false)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
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
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
		log.WithFields(log.Fields{"config": viper.AllSettings()}).Debug("Configuration:")
	}
	return validateConfig()
}

func validateConfig() error {
	if !viper.IsSet("database.username") {
		return fmt.Errorf("No database user found")
	}
	if !viper.IsSet("database.dbname") {
		return fmt.Errorf("No database name found")
	}
	return nil
}
