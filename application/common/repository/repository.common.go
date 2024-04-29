package repository

import (
	"antrein/bc-queue/application/common/resource"
	"antrein/bc-queue/internal/repository/config"
	cfg "antrein/bc-queue/model/config"
)

type CommonRepository struct {
	ConfigRepo *config.Repository
}

func NewCommonRepository(cfg *cfg.Config, rsc *resource.CommonResource) (*CommonRepository, error) {
	configRepo := config.New(cfg, rsc.Redis, rsc.GRPC)

	commonRepo := CommonRepository{
		ConfigRepo: configRepo,
	}
	return &commonRepo, nil
}
