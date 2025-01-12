package usecase

import (
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	repo interfaces.ITaskRepository
}

func NewTaskUsecase(repo interfaces.ITaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

func (u *TaskUsecase) CreateTask(projectID uuid.UUID, name string, description *string) (*entities.Task, error) {
	task, err := entities.NewTask(projectID, name, description)
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(task)
	return task, err
}
