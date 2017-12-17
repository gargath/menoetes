package main

import (
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"

	"github.com/gargath/menoetes/server"

)

func main() {
	err := LoadConfig()
  if err != nil {
		log.WithFields(log.Fields{"error":err,}).Warn("Error reading config file")
	}

	s := server.New(viper.Get("tls.certfile").(string), viper.Get("tls.keyfile").(string))
	s.Run()

}
