package controller

import (
	"time-tracker/adapter/presenter"
	"time-tracker/usecase"
)

type ProjectController struct {
	usecase   *usecase.ProjectUsecase
	presenter *presenter.ProjectPresenter
}

func NewProjectController(u *usecase.ProjectUsecase, p *presenter.ProjectPresenter) *ProjectController {
	return &ProjectController{
		usecase:   u,
		presenter: p,
	}
}

func (c *ProjectController) CreateProject(name string, description *string) {
	project, err := c.usecase.CreateProject(name, description)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	c.presenter.ShowCreateSuccess(project)
}

func (c *ProjectController) ListProjects() {
	projects, err := c.usecase.ListProjects()
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	c.presenter.ShowProjects(projects)
}
