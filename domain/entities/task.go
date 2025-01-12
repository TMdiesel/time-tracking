package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	ProjectID   uuid.UUID
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(projectID uuid.UUID, name string, description *string) (*Task, error) {
	if name == "" {
		return nil, errors.New("task name not specified")
	}
	return &Task{
		ID:          uuid.New(),
		ProjectID:   projectID,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
