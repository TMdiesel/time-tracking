package service

import (
	"errors"
	"time"
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"

	"github.com/google/uuid"
)

type TimeEntryService struct {
	taskRepo interfaces.ITaskRepository
}

func NewTimeEntryService(taskRepo interfaces.ITaskRepository) *TimeEntryService {
	return &TimeEntryService{
		taskRepo: taskRepo,
	}
}

func (s *TimeEntryService) NewTimeEntry(taskID uuid.UUID) (*entities.TimeEntry, error) {
	// taskIDが存在するかを確認する
	task, err := s.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		err = errors.New("task not found")
		return nil, err
	}

	return &entities.TimeEntry{
		ID:        uuid.New(),
		TaskID:    taskID,
		StartedAt: time.Now(),
	}, nil
}
