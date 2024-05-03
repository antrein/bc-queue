package wr

import (
	"antrein/bc-queue/application/common/repository"
	guard "antrein/bc-queue/application/middleware"
	"antrein/bc-queue/model/config"
	"context"
	"fmt"
	"net/http"
	"time"
)

type Handler struct {
	cfg  *config.Config
	repo *repository.CommonRepository
}

func New(cfg *config.Config, repo *repository.CommonRepository) *Handler {
	return &Handler{
		cfg:  cfg,
		repo: repo,
	}
}

func (h *Handler) RegisterHandler(app *http.ServeMux) {
	app.HandleFunc("/bc/queue/gg", guard.DefaultGuard(h.RegisterQueue))
}

func (h *Handler) RegisterQueue(g *guard.GuardContext) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	projectID := g.Request.Host
	fmt.Println(projectID)
	config, err := h.repo.ConfigRepo.GetProjectConfig(ctx, projectID)
	if err != nil {
		return g.ReturnError(500, err.Error())
	}
	return g.ReturnSuccess(config)
}
