package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/models"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/repository"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/pkg/database"
)

type Service struct {
	repo  *repository.Repository
	redis *database.RedisClient
}

func NewService(repo *repository.Repository, redis *database.RedisClient) *Service {
	return &Service{
		repo:  repo,
		redis: redis,
	}
}

func (s *Service) generateTasksListKey() string {
	return "tasks:list"
}

func (s *Service) generateTaskKey(id int32) string {
	return fmt.Sprintf("task:%d", id)
}

func (s *Service) CreateTask(title, description string) (*models.Task, error) {

	task, err := s.repo.CreateTask(title, description)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	listKey := s.generateTasksListKey()
	if err := s.redis.DeleteTask(ctx, listKey); err != nil {
		log.Printf("Warning: Failed to invalidate tasks list cache: %v", err)
	}

	log.Printf("Task created with ID: %d (cache invalidated)", task.ID)
	return task, nil
}

func (s *Service) GetAllTasks() ([]*models.Task, error) {
	ctx := context.Background()

	listKey := s.generateTasksListKey()
	cachedTasks, err := s.redis.GetTask(ctx, listKey)
	if err == nil {
		var tasks []*models.Task
		if err := json.Unmarshal([]byte(cachedTasks), &tasks); err == nil {
			log.Printf("Tasks list retrieved from cache (%d tasks)", len(tasks))
			return tasks, nil
		}
		log.Printf("Warning: Failed to unmarshal cached tasks list: %v", err)
	}
	log.Printf("Cache miss: retrieving tasks from database")
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	tasksJSON, err := json.Marshal(tasks)
	if err == nil {
		if err := s.redis.SetTask(ctx, listKey, tasksJSON); err != nil {
			log.Printf("Warning: Failed to cache tasks list: %v", err)
		} else {
			log.Printf("Tasks list cached successfully (%d tasks)", len(tasks))
		}
	}

	return tasks, nil
}

func (s *Service) DeleteTask(id int32) error {
	err:=s.repo.DeleteTask(id)
	if err!=nil{
		return err
	}
	ctx:=context.Background()
	listKey:=s.generateTasksListKey()
	if err:=s.redis.DeleteTask(ctx,listKey);err!=nil{
		log.Printf("Warning: Failed to invalidate tasks list cache: %v", err)
	}
	log.Printf("Task deleted with ID: %d (cache invalidated)", id)
	return nil
}

func (s *Service) UpdateTaskStatus(id int32) (*models.Task, error) {
	
	task, err := s.repo.UpdateTaskStatus(id)
	if err != nil {
		return nil, err
	}

	
	ctx := context.Background()
	listKey := s.generateTasksListKey()
	if err := s.redis.DeleteTask(ctx, listKey); err != nil {
		log.Printf("Warning: Failed to invalidate tasks list cache: %v", err)
	}

	log.Printf("Task updated with ID: %d (cache invalidated)", id)
	return task, nil
}
