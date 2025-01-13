package usecase

import (
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"
	"time-tracker/domain/service"
)

type ProjectUsecase struct {
	repo    interfaces.IProjectRepository
	service *service.ProjectService
}

func NewProjectUsecase(repo interfaces.IProjectRepository, service *service.ProjectService) *ProjectUsecase {
	return &ProjectUsecase{repo: repo, service: service}
}

func (u *ProjectUsecase) CreateProject(name string, description *string) (*entities.Project, error) {
	project, err := u.service.NewProject(name, description)
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(project)
	return project, err
}

// archive されていない project 一覧を取得する
func (u *ProjectUsecase) ListProjects() ([]entities.Project, error) {
	return u.repo.FindAllActive()
}

func (u *ProjectUsecase) ArchiveProject(project entities.Project) error {
	project.Archive()
	return u.repo.Update(&project)
}
