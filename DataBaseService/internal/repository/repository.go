package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateTask(title, description string) (*models.Task, error) {
	query := `INSERT INTO tasks (title,description,is_completed,created_at,updated_at)
	VALUES($1,$2,$3,$4,$5) RETURNING id,title,description,is_completed,created_at,updated_at`

	now := time.Now()
	var task models.Task

	err := r.db.QueryRow(query, title, description, false, now, now).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.IsCompleted,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)

		return &task, nil
	}
}
