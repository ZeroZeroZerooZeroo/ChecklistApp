package service

import (
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/models"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(title, description string) (*models.Task, error) {
	return s.repo.CreateTask(title, description)
}

func (s *Service) GetAllTasks() ([]*models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *Service) DeleteTask(id int32) error {
	return s.repo.DeleteTask(id)
}

func (s *Service) UpdateTaskStatus(id int32) (*models.Task, error) {
	return s.repo.UpdateTaskStatus(id)
}
