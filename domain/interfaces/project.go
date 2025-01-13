package interfaces

import "time-tracker/domain/entities"

type IProjectRepository interface {
	Create(project *entities.Project) error
	FindAllActive() ([]entities.Project, error)
	IsNameDuplicated(name string) (bool, error)
	Update(project *entities.Project) error
}
