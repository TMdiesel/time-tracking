package repositories

import (
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"
	"time-tracker/infrastructure/database/model"

	"gorm.io/gorm"
)

type TimeEntryRepository struct {
	db *gorm.DB
}

func NewTimeEntryRepository(db *gorm.DB) interfaces.ITimeEntryRepository {
	return &TimeEntryRepository{db: db}
}

func (r *TimeEntryRepository) Create(timeEntry *entities.TimeEntry) error {
	return r.db.Create(&model.TimeEntry{
		ID:        timeEntry.ID,
		TaskID:    timeEntry.TaskID,
		StartedAt: timeEntry.StartedAt,
		EndedAt:   timeEntry.EndedAt,
		Duration:  timeEntry.Duration,
	}).Error
}
