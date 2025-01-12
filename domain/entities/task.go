package entities

import (
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
