package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/models"
)

type Handlers struct {
	taskList *models.TaskListResponse
}

func NewHandler() *Handlers {
	return &Handlers{taskList: &models.TaskListResponse{
		Tasks: []models.TaskResponse{},
		Count: 0,
	}}
}

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {

	// проверка на метод
	// if r.Method != "POST" {
	// 	json.NewEncoder(w).Encode("The HTTP method is specified incorrectly")
	// 	return
	// }

	var task models.TaskResponse // изменить на CreateTaskResponse

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

	h.taskList.Tasks = append(h.taskList.Tasks, task)
	h.taskList.Count++

	// временный ответ на успешную отправуц
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&task)
	fmt.Println("Request was received successfully!")

}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {

	// проверка на метод
	// if r.Method != "GET" {
	// 	json.NewEncoder(w).Encode("The HTTP method is specified incorrectly")
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.taskList)

}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %s, URL: %s\n", r.Method, r.URL.Path)
	// проверка на метод
	// if r.Method != "DELETE" {
	// 	json.NewEncoder(w).Encode("The HTTP method is specified incorrectly")
	// 	return
	// }

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

	// временный ответ на успешную отправуц
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&task)
	fmt.Println("Request was received successfully!")

}

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {

	// проверка на метод
	// if r.Method != "PUT" {
	// 	json.NewEncoder(w).Encode("The HTTP method is specified incorrectly")
	// 	return
	// }

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

	// временный ответ на успешную отправуц
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&task)
	fmt.Println("Request was received successfully!")

}
