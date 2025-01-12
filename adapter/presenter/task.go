package presenter

import (
	"fmt"
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
}
