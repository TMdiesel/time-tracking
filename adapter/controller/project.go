package controller

import (
	"fmt"
	"sort"
	"time-tracker/adapter/presenter"
	"time-tracker/domain/entities"
	"time-tracker/usecase"

	"github.com/manifoldco/promptui"
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

func (c *ProjectController) ArchiveProject() {
	// プロジェクト一覧の取得
	projects, err := c.usecase.ListProjects()
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	if len(projects) == 0 {
		c.presenter.ShowError(fmt.Errorf("No projects available. Please create a project first"))
	}
	// 名前順でソート
	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].Name < projects[j].Name
	})

	// インタラクティブにプロジェクト選択
	selectedProject, err := c.selectProjectInteractive(projects)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}

	err = c.usecase.ArchiveProject(selectedProject)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}

	c.presenter.ShowCreateSuccess(&selectedProject)

}

// インタラクティブにプロジェクト選択
func (c *ProjectController) selectProjectInteractive(projects []entities.Project) (entities.Project, error) {
	projectNames := []string{}
	for _, project := range projects {
		projectNames = append(projectNames, project.Name)
	}

	// インタラクティブ選択
	prompt := promptui.Select{
		Label: "Select a Project",
		Items: projectNames,
	}
	index, _, err := prompt.Run()
	if err != nil {
		return entities.Project{}, fmt.Errorf("Project selection failed: %v", err)
	}

	return projects[index], nil
}
