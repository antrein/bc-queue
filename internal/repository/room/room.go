package room

import (
	"antrein/bc-queue/model/config"
	"antrein/bc-queue/model/entity"
	"context"
	"encoding/json"
	"fmt"

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

func (r *Repository) AddUserToRoom(ctx context.Context, key string, session entity.Session) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	_, err = r.redisClient.LPush(ctx, key, data).Result()
	return err
}

func (r *Repository) AddUserToWaitingRoom(ctx context.Context, projectID string, session entity.Session) error {
	return r.AddUserToRoom(ctx, fmt.Sprintf("%s:waiting", projectID), session)
}

func (r *Repository) AddUserToMainRoom(ctx context.Context, projectID string, session entity.Session) error {
	return r.AddUserToRoom(ctx, fmt.Sprintf("%s:main", projectID), session)
}

func (r *Repository) RemoveUserFromRoom(ctx context.Context, projectID string, roomType string, sessionID string) error {
	key := fmt.Sprintf("%s:%s", projectID, roomType)
	sessions, err := r.redisClient.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return err
	}

	for _, sessionStr := range sessions {
		var session entity.Session
		json.Unmarshal([]byte(sessionStr), &session)
		if session.SessionID == sessionID {
			_, err := r.redisClient.LRem(ctx, key, 1, sessionStr).Result()
			return err
		}
	}

	return fmt.Errorf("User not found in %s room", roomType)
}

func (r *Repository) CountUserInRoom(ctx context.Context, projectID string, roomType string) (int64, error) {
	key := fmt.Sprintf("%s:%s", projectID, roomType)
	return r.redisClient.LLen(ctx, key).Result()
}

func (r *Repository) GetUserFromRoom(ctx context.Context, projectID string, roomType string) ([]entity.Session, error) {
	key := fmt.Sprintf("%s:%s", projectID, roomType)
	sessionsStr, err := r.redisClient.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var sessions []entity.Session
	for _, s := range sessionsStr {
		var session entity.Session
		if err := json.Unmarshal([]byte(s), &session); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}
