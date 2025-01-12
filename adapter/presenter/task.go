package presenter

import (
	"fmt"
	"os"
	"sort"
	"time-tracker/domain/dto"
	"time-tracker/domain/entities"
)

type TaskPresenter struct{}

func NewTaskPresenter() *TaskPresenter {
	return &TaskPresenter{}
}

func (p *TaskPresenter) ShowCreateSuccess(task *entities.Task) {
	fmt.Printf("✅ Task '%s' has been created.\n", task.Name)
}

func (p *TaskPresenter) ShowError(err error) {
	fmt.Printf("❌ Error: %v\n", err)
	os.Exit(1)
}

func (p *TaskPresenter) ShowTasks(tasks []dto.TaskWithProjectDTO) {
	if len(tasks) == 0 {
		fmt.Print("No tasks found. Please create a new task to get started.")
		return
	}

	//タスクのソート（ProjectName → Nameの順）
	sort.SliceStable(tasks, func(i, j int) bool {
		if tasks[i].ProjectName == tasks[j].ProjectName {
			return tasks[i].Name < tasks[j].Name
		}
		return tasks[i].ProjectName < tasks[j].ProjectName
	})

	for _, task := range tasks {
		desc := "(No description)"
		if task.Description != nil {
			desc = *task.Description
		}
		fmt.Printf("- [%s] %s: %s\n", task.ProjectName, task.Name, desc)
	}
}
