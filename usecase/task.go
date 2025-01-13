package usecase

import (
	"fmt"
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

// archive されていない task 一覧を取得する
func (u *TaskUsecase) ListTasks() ([]dto.TaskWithProjectDTO, error) {
	return u.taskRepo.FindAllActiveWithProjectName()
}

func (u *TaskUsecase) StartTask(taskID uuid.UUID) error {
	// 実行中のタスクが存在する場合は開始できない
	runningTimeEntries, err := u.timeEntryRepo.GetAllRunning()
	if err != nil {
		return err
	}
	if len(runningTimeEntries) > 0 {
		return fmt.Errorf("Another task is already running. Please end it before starting a new one.")
	}

	timeEntry, err := u.timeEntryService.NewTimeEntry(taskID)
	if err != nil {
		return err
	}
	err = u.timeEntryRepo.Create(timeEntry)
	return err
}

func (u *TaskUsecase) EndTask() error {
	runningTimeEntries, err := u.timeEntryRepo.GetAllRunning()
	if err != nil {
		return err
	}
	if len(runningTimeEntries) == 0 {
		return fmt.Errorf("No running task to end.")
	}

	// タスク停止処理
	if len(runningTimeEntries) >= 2 {
		fmt.Println("⚠️ Multiple running tasks detected. All will be ended.")

	}
	for _, entry := range runningTimeEntries {
		entry.End()
		if err := u.timeEntryRepo.Update(&entry); err != nil {
			return err
		}
	}

	return nil
}
