package entities

import (
	"time"

	"github.com/google/uuid"
)

type TimeEntry struct {
	ID        uuid.UUID
	TaskID    uuid.UUID
	StartedAt time.Time
	EndedAt   *time.Time
	Duration  *time.Duration
}
