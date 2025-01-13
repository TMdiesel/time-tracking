package interfaces

import (
	"time-tracker/domain/entities"
)

type ITimeEntryRepository interface {
	Create(timeEntry *entities.TimeEntry) error
	GetAllRunning() ([]entities.TimeEntry, error)
	Update(timeEntry *entities.TimeEntry) error
}
