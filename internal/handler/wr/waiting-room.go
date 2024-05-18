package wr

import (
	"antrein/bc-queue/application/common/repository"
	guard "antrein/bc-queue/application/middleware"
	"antrein/bc-queue/internal/utils"
	"antrein/bc-queue/model/config"
	"antrein/bc-queue/model/dto"
	"antrein/bc-queue/model/entity"
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	app.HandleFunc("/bc/queue/register", guard.DefaultGuard(h.RegisterQueue))
}

func (h *Handler) RegisterQueue(g *guard.GuardContext) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	host := g.Request.Referer()
	projectID, err := utils.ExtractProjectID(host)
	if err != nil {
		return g.ReturnError(500, err.Error())
	}
	config, err := h.repo.ConfigRepo.GetProjectConfig(ctx, projectID)
	if err != nil {
		return g.ReturnError(500, err.Error())
	}
	currentUser, err := h.repo.RoomRepo.CountUserInRoom(ctx, projectID, "main")
	if err != nil {
		return g.ReturnError(500, err.Error())
	}
	sessionID := uuid.New()
	session := entity.Session{
		SessionID:  sessionID.String(),
		EnqueuedAt: time.Now(),
	}
	roomClaim := entity.JWTClaim{
		SessionID: sessionID.String(),
		ProjectID: projectID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    projectID,
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	waitingRoomToken, err := utils.GenerateJWTToken(h.cfg.Secrets.WaitingRoomSecret, roomClaim)
	if err != nil {
		return g.ReturnError(500, err.Error())
	}
	if currentUser < int64(config.Threshold) {
		err = h.repo.RoomRepo.AddUserToMainRoom(ctx, projectID, session)
		if err != nil {
			return g.ReturnError(500, err.Error())
		}
		mainRoomToken, err := utils.GenerateJWTToken(h.cfg.Secrets.MainRoomSecret, roomClaim)
		if err != nil {
			return g.ReturnError(500, err.Error())
		}
		tokens := dto.RegisterQueueResponse{
			WaitingRoomToken: waitingRoomToken,
			MainRoomToken:    mainRoomToken,
		}
		return g.ReturnSuccess(tokens)
	}
	err = h.repo.RoomRepo.AddUserToWaitingRoom(ctx, projectID, session)
	if err != nil {
		return g.ReturnError(500, err.Error())
	}
	tokens := dto.RegisterQueueResponse{
		WaitingRoomToken: waitingRoomToken,
		MainRoomToken:    "",
	}
	return g.ReturnSuccess(tokens)
}
