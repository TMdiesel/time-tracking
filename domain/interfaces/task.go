package interfaces

import (
	"time-tracker/domain/dto"
	"time-tracker/domain/entities"

	"github.com/google/uuid"
)

type ITaskRepository interface {
	Create(task *entities.Task) error
	IsNameDuplicated(projectID uuid.UUID, name string) (bool, error)
	FindAllActiveWithProjectName() ([]dto.TaskWithProjectDTO, error)
	GetTaskByID(taskID uuid.UUID) (*entities.Task, error)
	GetByProjectID(projectID uuid.UUID) ([]entities.Task, error)
	Update(task *entities.Task) error
}
