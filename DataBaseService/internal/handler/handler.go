package handler

import (
	"context"
	"log"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/models"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/service"
	pb "github.com/ZeroZeroZerooZeroo/ChecklistApp/proto/checklist/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCHandler struct {
	pb.UnimplementedCheckListServiceServer
	taskService *service.Service
}

func NewGRPCHandler(taskService *service.Service) *GRPCHandler {
	return &GRPCHandler{
		taskService: taskService,
	}
}

func (h *GRPCHandler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("Creating task:%s", req.Title)

	task, err := h.taskService.CreateTask(req.Title, req.Description)
	if err != nil {
		log.Printf("Failed to create task:%v", err)
		return nil, status.Error(codes.Internal, "failed to create task")

	}

	log.Printf("Task created with ID:%d", task.ID)
	return h.taskToProto(task), nil

}

func (h *GRPCHandler) taskToProto(task *models.Task) *pb.TaskResponse {
	return &pb.TaskResponse{
		Id:          int32(task.ID),
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CreatedAt:   timestamppb.New(task.CreatedAt),
		UpdatedAt:   timestamppb.New(task.UpdatedAt),
	}
}

func (h *GRPCHandler) tasksToProto(tasks []*models.Task) *pb.TaskListResponse {
	protoTask := make([]*pb.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		protoTask = append(protoTask, h.taskToProto(task))
	}

	return &pb.TaskListResponse{
		Tasks: protoTask,
		Count: int32(len(protoTask)),
	}
}
