package dto

import (
	"time"

	"github.com/google/uuid"
)

type TimeEntryWithProjectAndTaskName struct {
	ID          uuid.UUID
	TaskID      uuid.UUID
	StartedAt   time.Time
	EndedAt     *time.Time
	Duration    *time.Duration
	ProjectName string
	TaskName    string
}
