package interfaces

import "time-tracker/domain/entities"

type IProjectRepository interface {
	Create(project *entities.Project) error
	FindAll() ([]entities.Project, error)
	IsNameDuplicated(name string) (bool, error)
}
