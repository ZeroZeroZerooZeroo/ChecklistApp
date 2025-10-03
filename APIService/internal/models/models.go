package models

import "time"

// основная структура задачи
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// структура для парсинга тела запроса при создании задачи
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// структура для парсинга тела запроса при удалении задачи
type DeleteTaskRequest struct {
	ID int `json:"id"`
}

// структура для парсинга тела запроса при изменении задачи
type UpdateTaskStatusRequest struct {
	ID int `json:"id"`
}

// структура для формирования ответа с 1 задачей
type TaskResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// структура для формирования ответа со списком задач
type TaskListResponse struct {
	Tasks []TaskResponse `json:"tasks"`
	Count int            `json:"count"`
}
