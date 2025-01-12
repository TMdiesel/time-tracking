package service

import (
	"errors"
	"fmt"
	"time"
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"

	"github.com/google/uuid"
)

type TaskService struct {
	repo interfaces.ITaskRepository
}

func NewTaskService(repo interfaces.ITaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) NewTask(projectID uuid.UUID, name string, description *string) (*entities.Task, error) {
	if name == "" {
		return nil, errors.New("task name not specified")
	}
	// 同一 project の中に同名の task が存在するかを確認する
	isDuplicated, err := s.repo.IsNameDuplicated(projectID, name)
	if err != nil {
		return nil, err
	}
	if isDuplicated {
		return nil, fmt.Errorf("task with the name %s already exists", name)
	}

	return &entities.Task{
		ID:          uuid.New(),
		ProjectID:   projectID,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
