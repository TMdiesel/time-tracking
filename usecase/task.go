package usecase

import (
	"time-tracker/domain/dto"
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"
	"time-tracker/domain/service"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	taskRepo         interfaces.ITaskRepository
	timeEntryRepo    interfaces.ITimeEntryRepository
	taskService      *service.TaskService
	timeEntryService *service.TimeEntryService
}

func NewTaskUsecase(taskRepo interfaces.ITaskRepository, timeEntryRepo interfaces.ITimeEntryRepository, taskService *service.TaskService, timeEntryService *service.TimeEntryService) *TaskUsecase {
	return &TaskUsecase{taskRepo: taskRepo, timeEntryRepo: timeEntryRepo, taskService: taskService, timeEntryService: timeEntryService}
}

func (u *TaskUsecase) CreateTask(projectID uuid.UUID, name string, description *string) (*entities.Task, error) {
	task, err := u.taskService.NewTask(projectID, name, description)
	if err != nil {
		return nil, err
	}

	err = u.taskRepo.Create(task)
	return task, err
}

func (u *TaskUsecase) ListTasks() ([]dto.TaskWithProjectDTO, error) {
	return u.taskRepo.FindAllWithProjectName()
}

func (u *TaskUsecase) StartTask(taskID uuid.UUID) error {
	timeEntry, err := u.timeEntryService.NewTimeEntry(taskID)
	if err != nil {
		return err
	}
	err = u.timeEntryRepo.Create(timeEntry)
	return err
}
