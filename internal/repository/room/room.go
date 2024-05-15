package room

import (
	"antrein/bc-queue/model/config"
	"context"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	cfg         *config.Config
	redisClient *redis.Client
}

func New(cfg *config.Config, rc *redis.Client) *Repository {
	return &Repository{
		cfg:         cfg,
		redisClient: rc,
	}
}

func (r *Repository) AddUserToWaitingRoom(ctx context.Context, sessionID string) error {
	return nil
}

func (r *Repository) AddUserToMainRoom(ctx context.Context, sessionID string) error {
	return nil
}

func (r *Repository) CountUserInWaitingRoom(ctx context.Context) (int, error) {
	return 0, nil
}

func (r *Repository) CountUserInMainRoom(ctx context.Context) (int, error) {
	return 0, nil
}
