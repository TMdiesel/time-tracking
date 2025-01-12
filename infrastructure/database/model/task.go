package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProjectID   uuid.UUID `gorm:"type:uuid;not null"`                            // 外部キー（Project.ID）
	Project     Project   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // リレーション
	Name        string    `gorm:"not null"`
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
