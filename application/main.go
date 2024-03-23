package main

import (
	"antrein/bc-queue/application/rest"
	"antrein/bc-queue/model/config"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	rest_app, err := rest.ApplicationDelegate(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err = rest.StartServer(cfg, rest_app); err != nil {
		log.Fatal(err)
	}

}
