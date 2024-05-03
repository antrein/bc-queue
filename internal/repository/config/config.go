package config

import (
	"antrein/bc-queue/model/config"
	"context"

	pb "github.com/antrein/proto-repository/pb/bc"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Repository struct {
	cfg         *config.Config
	redisClient *redis.Client
	grpcClient  *grpc.ClientConn
}

func New(cfg *config.Config, rc *redis.Client, gc *grpc.ClientConn) *Repository {
	return &Repository{
		cfg:         cfg,
		redisClient: rc,
		grpcClient:  gc,
	}
}

func (r *Repository) GetProjectConfig(ctx context.Context, host string) (*pb.ProjectConfigResponse, error) {
	svc := pb.NewProjectConfigServiceClient(r.grpcClient)
	req := &pb.ConfigRequest{Host: host}
	config, err := svc.GetProjectConfig(ctx, req)
	if err != nil {
		return nil, err
	}
	return config, nil
}
