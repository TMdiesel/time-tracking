package presenter

import (
	"fmt"
	"os"
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

func (p *TaskPresenter) ShowTasks(tasks []entities.Task) {
	for _, task := range tasks {
		desc := "(No description)"
		if task.Description != nil {
			desc = *task.Description
		}
		fmt.Printf("- %s: %s\n", task.Name, desc)

	}
}
