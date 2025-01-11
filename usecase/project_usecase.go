
package usecase

import (
    "github.com/google/uuid"
    "time-tracker/domain/entities"
    "time-tracker/domain/interfaces"
    "time"
)

type ProjectUsecase struct {
    repo interfaces.IProjectRepository
}

func NewProjectUsecase(repo interfaces.IProjectRepository) *ProjectUsecase {
    return &ProjectUsecase{repo: repo}
}

func (u *ProjectUsecase) CreateProject(name, description string) (*entities.Project, error) {
    project := &entities.Project{
        ID:          uuid.New(),
        Name:        name,
        Description: description,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    err := u.repo.Create(project)
    return project, err
}

func (u *ProjectUsecase) ListProjects() ([]entities.Project, error) {
    return u.repo.FindAll()
}
