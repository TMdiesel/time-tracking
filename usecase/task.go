package usecase

import (
	"time-tracker/domain/dto"
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"
	"time-tracker/domain/service"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	repo    interfaces.ITaskRepository
	service *service.TaskService
}

func NewTaskUsecase(repo interfaces.ITaskRepository, service *service.TaskService) *TaskUsecase {
	return &TaskUsecase{repo: repo, service: service}
}

func (u *TaskUsecase) CreateTask(projectID uuid.UUID, name string, description *string) (*entities.Task, error) {
	task, err := u.service.NewTask(projectID, name, description)
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(task)
	return task, err
}

func (u *TaskUsecase) ListTasks() ([]dto.TaskWithProjectDTO, error) {
	return u.repo.FindAllWithProjectName()
}
