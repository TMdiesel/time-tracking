package interfaces

import "time-tracker/domain/entities"

type ITaskRepository interface {
	Create(task *entities.Task) error
}
