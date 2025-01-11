
package usecase

import (
    "time-tracker/domain/entities"
    "time-tracker/domain/interfaces"
)

type IProjectUseCase interface {
    AddProject(name, description string) error
}

type projectUseCase struct {
    repo interfaces.IProjectRepository
}

func NewProjectUseCase(repo interfaces.IProjectRepository) IProjectUseCase {
    return &projectUseCase{
        repo: repo,
    }
}

func (u *projectUseCase) AddProject(name, description string) error {
    project := entities.NewProject(name, description)
    return u.repo.Save(project)
}
