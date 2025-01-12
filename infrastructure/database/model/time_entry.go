package model

import (
	"time"

	"github.com/google/uuid"
)

type TimeEntry struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	TaskID    uuid.UUID `gorm:"type:uuid;not null"`                            // 外部キー
	Task      Task      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // リレーション
	StartedAt time.Time
	EndedAt   *time.Time
	Duration  *time.Duration
}
