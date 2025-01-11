package usecase

import (
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"
)

type ProjectUsecase struct {
	repo interfaces.IProjectRepository
}

func NewProjectUsecase(repo interfaces.IProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{repo: repo}
}

func (u *ProjectUsecase) CreateProject(name string, description *string) (*entities.Project, error) {
	project, err := entities.NewProject(name, description)
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(project)
	return project, err
}

func (u *ProjectUsecase) ListProjects() ([]entities.Project, error) {
	return u.repo.FindAll()
}
