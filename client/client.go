package client

import (
	"context"
	"time"

	"antrein/bc-queue/internal/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Call(name string) (string, error) {
	conn, err := grpc.Dial("bc-dashboard.production.svc.cluster.local:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}
	return r.GetMessage(), nil
}
