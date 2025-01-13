package usecase

import (
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"
	"time-tracker/domain/service"
)

type ProjectUsecase struct {
	repo     interfaces.IProjectRepository
	taskRepo interfaces.ITaskRepository
	service  *service.ProjectService
}

func NewProjectUsecase(repo interfaces.IProjectRepository, taskRepo interfaces.ITaskRepository, service *service.ProjectService) *ProjectUsecase {
	return &ProjectUsecase{repo: repo, taskRepo: taskRepo, service: service}
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

// 指定された project と関連 task を archive する
func (u *ProjectUsecase) ArchiveProjectAndRelatedTasks(project entities.Project) error {
	project.Archive()
	err := u.repo.Update(&project)
	if err != nil {
		return err
	}

	// 関連するタスクもアーカイブ
	tasks, err := u.taskRepo.GetByProjectID(project.ID)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		task.Archive()
		if err := u.taskRepo.Update(&task); err != nil {
			return err
		}
	}
	return nil
}
