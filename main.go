package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/gargath/menoetes/pkg/server"
)

func main() {
	err := LoadConfig()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warn("Error reading config file")
		os.Exit(1)
	}

	s := server.New(viper.Get("tls.certfile").(string), viper.Get("tls.keyfile").(string), viper.GetBool("debug"))
	s.Run()

}
