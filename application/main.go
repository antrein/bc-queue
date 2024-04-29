package main

import (
	"antrein/bc-queue/application/common/repository"
	"antrein/bc-queue/application/common/resource"
	"antrein/bc-queue/application/rest"
	"antrein/bc-queue/model/config"
	"context"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	rsc, err := resource.NewCommonResource(cfg, ctx)
	if err != nil {
		log.Fatal(err)
	}
	repository, err := repository.NewCommonRepository(cfg, rsc)
	if err != nil {
		log.Fatal(err)
	}
	rest_app, err := rest.ApplicationDelegate(cfg, repository)
	if err != nil {
		log.Fatal(err)
	}

	if err = rest.StartServer(cfg, rest_app); err != nil {
		log.Fatal(err)
	}

}
