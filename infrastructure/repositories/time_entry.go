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

func (r *TimeEntryRepository) GetAllRunning() ([]entities.TimeEntry, error) {
	var timeEntryModels []model.TimeEntry
	err := r.db.Where("ended_at IS NULL").Find(&timeEntryModels).Error
	if err != nil {
		return nil, err
	}

	var timeEntries []entities.TimeEntry
	for _, entry := range timeEntryModels {
		timeEntries = append(timeEntries, entities.TimeEntry{
			ID:        entry.ID,
			TaskID:    entry.TaskID,
			StartedAt: entry.StartedAt,
			EndedAt:   entry.EndedAt,
			Duration:  entry.Duration,
		})
	}

	return timeEntries, nil
}
