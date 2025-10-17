package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/grpc"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/kafka"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/models"
)

type Handlers struct {
	grpcClient    *grpc.GRPCClient
	kafkaProducer *kafka.Producer
}

func NewHandler(grpcClient *grpc.GRPCClient, kafkaProducer *kafka.Producer) *Handlers {
	return &Handlers{
		grpcClient:    grpcClient,
		kafkaProducer: kafkaProducer,
	}
}

func (h *Handlers) getUserIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return forwarded
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown"
	}

	return ip
}

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {

	var task models.CreateTaskRequest

	// парсим запрос на создание
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// валидация данных
	if task.Title == "" {
		http.Error(w, "Title should not be empty!", http.StatusBadRequest)
		return
	}
	if task.Description == "" {
		http.Error(w, "Description should not be empty!", http.StatusBadRequest)
		return
	}

	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Отправка запроса на gRPC сервер
	resp, err := h.grpcClient.CreateTask(ctx, task.Title, task.Description)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create task: %v", err), http.StatusInternalServerError)
		return
	}

	//Отправка события в Kafka
	userIP:= h.getUserIP(r)
	if err:= h.kafkaProducer.SendCreateAction(ctx,userIP,int(resp.Id),task.Title);err!=nil{
		fmt.Printf("Failed to send Kafka event:%v\n",err)
	}


	// КОнвертация gRPC запроса
	taskResponse := models.TaskResponse{
		ID:          int(resp.Id),
		Title:       resp.Title,
		Description: resp.Description,
		IsCompleted: resp.IsCompleted,
		CreatedAt:   resp.CreatedAt.AsTime(),
		UpdatedAt:   resp.UpdatedAt.AsTime(),
	}

	// ответ на успешную отправуц
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&taskResponse)
	fmt.Println("Request was received successfully!")

}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {

	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := h.grpcClient.GetTasks(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get tasks: %v", err), http.StatusInternalServerError)
		return
	}

	//Отправка события в Kafka
	userIP:= h.getUserIP(r)
	if err:= h.kafkaProducer.SendGetAction(ctx,userIP);err!=nil{
		fmt.Printf("Failed to send Kafka event:%v\n",err)
	}

	// КОнвертация gRPC запроса
	tasks := make([]models.TaskResponse, 0, len(resp.Tasks))
	for _, task := range resp.Tasks {
		tasks = append(tasks, models.TaskResponse{
			ID:          int(task.Id),
			Title:       task.Title,
			Description: task.Description,
			IsCompleted: task.IsCompleted,
			CreatedAt:   task.CreatedAt.AsTime(),
			UpdatedAt:   task.UpdatedAt.AsTime(),
		})
	}

	taskListResponse := models.TaskListResponse{
		Tasks: tasks,
		Count: int(resp.Count),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskListResponse)

}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {

	var task models.DeleteTaskRequest

	// парсим запрос на удаление
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// валидация данных
	if task.ID == 0 {
		http.Error(w, "ID should not be empty!", http.StatusBadRequest)
		return
	}

	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Отправка запроса на gRPC сервер
	err := h.grpcClient.DeleteTask(ctx, int32(task.ID))

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete task: %v", err), http.StatusInternalServerError)
		return
	}


	//Отправка события в Kafka
	userIP:= h.getUserIP(r)
	if err:= h.kafkaProducer.SendDeleteAction(ctx,userIP,task.ID);err!=nil{
		fmt.Printf("Failed to send Kafka event:%v\n",err)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Task deleted successfully",
		"id":      task.ID,
	})
	fmt.Println("Request was received successfully!")

}

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {

	var task models.UpdateTaskStatusRequest

	// парсим запрос на изменение
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// валидация данных
	if task.ID == 0 {
		http.Error(w, "ID should not be empty!", http.StatusBadRequest)
		return
	}

	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Отправка запроса на gRPC сервер
	resp, err := h.grpcClient.UpdateTaskStatus(ctx, int32(task.ID))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update task: %v", err), http.StatusInternalServerError)
		return
	}

	//Отправка события в Kafka
	userIP:= h.getUserIP(r)
	if err:= h.kafkaProducer.SendUpdateAction(ctx,userIP,task.ID);err!=nil{
		fmt.Printf("Failed to send Kafka event:%v\n",err)
	}

	// КОнвертация gRPC запроса
	taskResponse := models.TaskResponse{
		ID:          int(resp.Id),
		Title:       resp.Title,
		Description: resp.Description,
		IsCompleted: resp.IsCompleted,
		CreatedAt:   resp.CreatedAt.AsTime(),
		UpdatedAt:   resp.UpdatedAt.AsTime(),
	}

	// временный ответ на успешную отправуц
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&taskResponse)
	fmt.Println("Request was received successfully!")

}
