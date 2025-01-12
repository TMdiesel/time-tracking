package dto

import (
	"time"

	"github.com/google/uuid"
)

type TaskWithProjectDTO struct {
	ID          uuid.UUID
	Name        string
	Description *string
	ProjectID   uuid.UUID
	ProjectName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
