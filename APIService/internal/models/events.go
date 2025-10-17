package models

import "time"

type UserActionEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Action    string    `json:"action"`
	UserIP    string    `json:"user_ip,omitempty"`
	TaskID    int       `json:"task_id,omitempty"`
	Details   string    `json:"details,omitempty"`
}
