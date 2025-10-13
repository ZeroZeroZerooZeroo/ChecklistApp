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

	}
	return &task, nil
}

func (r *Repository) GetAllTasks() ([]*models.Task, error) {
	query := `SELECT id,title,description,is_completed,created_at,updated_at
	FROM tasks`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w:", err)
	}

	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.IsCompleted,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)

		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error inerating tasks: %w", err)

	}

	return tasks, nil
}

func (r *Repository) DeleteTask(id int32) error {
	query := `DELETE FROM tasks WHERE id=$1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)

	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}

	return nil
}

func (r *Repository) UpdateTaskStatus(id int32) (*models.Task, error) {
	query := `UPDATE tasks SET is_completed = true, updated_at = $1
	WHERE id = $2 RETURNING id, title, description, is_completed, created_at, updated_at`

	var task models.Task
	err := r.db.QueryRow(query, time.Now(), id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.IsCompleted,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return &task, nil
}
