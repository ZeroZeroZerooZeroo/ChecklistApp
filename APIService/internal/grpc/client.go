package grpc

import (
	"context"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/proto/checklist/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client pb.CheckListServiceClient
}

func NewGRPCClient(serverAddress string) (*GRPCClient, error) {

	conn, err := grpc.NewClient(
		serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err

	}

	client := pb.NewCheckListServiceClient(conn)

	return &GRPCClient{
		conn:   conn,
		client: client,
	}, nil
}

func (c *GRPCClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *GRPCClient) CreateTask(ctx context.Context, title, description string) (*pb.TaskResponse, error) {
	req := &pb.CreateTaskRequest{
		Title:       title,
		Description: description,
	}

	return c.client.CreateTask(ctx, req)
}


