package grpc

import (
	"context"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/proto/checklist/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client pb.CheckListServiceClient
}

func NewGRPCClient(serverAddress string) (*GRPCClient, error) {

	// Установка соединение с gRPC сервером
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

// CreateTask отправляет запрос на создание задачи на gRPC сервер
func (c *GRPCClient) CreateTask(ctx context.Context, title, description string) (*pb.TaskResponse, error) {
	req := &pb.CreateTaskRequest{
		Title:       title,
		Description: description,
	}

	return c.client.CreateTask(ctx, req)
}

// GetTasks получает список задач с gRPC сервера
func (c *GRPCClient) GetTasks(ctx context.Context) (*pb.TaskListResponse, error) {
	return c.client.GetTasks(ctx, &emptypb.Empty{})
}

// DeleteTask отправляет запрос на удаление задачи на gRPC сервер
func (c *GRPCClient) DeleteTask(ctx context.Context, id int32) error {
	req := &pb.DeleteTaskRequest{
		Id: id,
	}

	_, err := c.client.DeleteTask(ctx, req)
	return err
}

// UpdateTaskStatus отправляет запрос на обновление стутуса задачи
func (c *GRPCClient) UpdateTaskStatus(ctx context.Context, id int32) (*pb.TaskResponse, error) {
	req := &pb.UpdateTaskStatusRequest{
		Id: id,
	}

	return c.client.UpdateTaskStatus(ctx, req)
}
