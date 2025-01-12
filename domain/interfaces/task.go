package interfaces

import (
	"time-tracker/domain/entities"

	"github.com/google/uuid"
)

type ITaskRepository interface {
	Create(task *entities.Task) error
	IsNameDuplicated(projectID uuid.UUID, name string) (bool, error)
	FindAll() ([]entities.Task, error)
}
