package repositories

import (
	"fmt"
	"time"
	"time-tracker/domain/dto"
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

func (r *TimeEntryRepository) Update(timeEntry *entities.TimeEntry) error {
	updatedModel := model.TimeEntry{
		ID:        timeEntry.ID,
		TaskID:    timeEntry.TaskID,
		StartedAt: timeEntry.StartedAt,
		EndedAt:   timeEntry.EndedAt,
		Duration:  timeEntry.Duration,
	}
	if err := r.db.Save(&updatedModel).Error; err != nil {
		return fmt.Errorf("failed to update time entry: %w", err)
	}

	return nil
}

func (r *TimeEntryRepository) ListTimeEntriesWithProjectAndTaskName(from, to *time.Time) ([]dto.TimeEntryWithProjectAndTaskName, error) {
	var timeEntries []model.TimeEntry

	query := r.db.Preload("Task.Project")
	if from != nil {
		query = query.Where("started_at >= ?", *from)
	}
	if to != nil {
		query = query.Where("started_at <= ?", *to)
	}
	if err := query.Find(&timeEntries).Error; err != nil {
		return nil, err
	}

	// DTOへの変換
	var results []dto.TimeEntryWithProjectAndTaskName
	for _, entry := range timeEntries {
		results = append(results, dto.TimeEntryWithProjectAndTaskName{
			ID:          entry.ID,
			TaskID:      entry.TaskID,
			StartedAt:   entry.StartedAt,
			EndedAt:     entry.EndedAt,
			Duration:    entry.Duration,
			ProjectName: entry.Task.Project.Name,
			TaskName:    entry.Task.Name,
		})
	}

	return results, nil
}
