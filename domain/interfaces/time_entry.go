package interfaces

import (
	"time"
	"time-tracker/domain/dto"
	"time-tracker/domain/entities"
)

type ITimeEntryRepository interface {
	Create(timeEntry *entities.TimeEntry) error
	GetAllRunning() ([]entities.TimeEntry, error)
	Update(timeEntry *entities.TimeEntry) error
	ListTimeEntriesWithProjectAndTaskName(from, to *time.Time) ([]dto.TimeEntryWithProjectAndTaskName, error)
}
