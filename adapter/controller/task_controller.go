package controller

import (
	"time-tracker/adapter/presenter"
	"time-tracker/usecase"

	"github.com/google/uuid"
)

type TaskController struct {
	usecase   *usecase.TaskUsecase
	presenter *presenter.TaskPresenter
}

func NewTaskController(u *usecase.TaskUsecase, p *presenter.TaskPresenter) *TaskController {
	return &TaskController{
		usecase:   u,
		presenter: p,
	}
}

func (c *TaskController) CreateTask(projectID string, name string, description *string) {
	pid, err := uuid.Parse(projectID)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	project, err := c.usecase.CreateTask(pid, name, description)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	c.presenter.ShowCreateSuccess(project)
}
