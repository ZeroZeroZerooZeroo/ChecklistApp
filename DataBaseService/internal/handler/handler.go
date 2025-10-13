package handler

import (
	"context"
	"log"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/models"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/service"
	pb "github.com/ZeroZeroZerooZeroo/ChecklistApp/proto/checklist/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (h *GRPCHandler) GetTasks(ctx context.Context, req *emptypb.Empty) (*pb.TaskListResponse, error) {

	log.Printf("Getting all tasks")

	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		log.Printf("Failed to get tasks: %v", err)
		return nil, status.Error(codes.Internal, "failed to get tasks")
	}

	log.Printf("Returning %d tasks", len(tasks))
	return h.tasksToProto(tasks), nil
}

func (h *GRPCHandler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*emptypb.Empty, error) {
	log.Printf("Deleting task with ID: %d", req.Id)

	err := h.taskService.DeleteTask(req.Id)

	if err != nil {
		log.Printf("Failed to delete task: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Task deleted with ID: %d", req.Id)
	return &emptypb.Empty{}, nil
}

func (h *GRPCHandler) UpdateTaskStatus(ctx context.Context, req *pb.UpdateTaskStatusRequest) (*pb.TaskResponse, error) {
	log.Printf("Updating task status for ID: %d", req.Id)

	task, err := h.taskService.UpdateTaskStatus(req.Id)
	if err != nil {
		log.Printf("Failed to update task: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Printf("Task updated with ID: %d", task.ID)
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
