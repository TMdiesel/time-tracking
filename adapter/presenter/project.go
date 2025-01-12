package presenter

import (
	"fmt"
	"os"
	"sort"
	"time-tracker/domain/entities"
)

type ProjectPresenter struct{}

func NewProjectPresenter() *ProjectPresenter {
	return &ProjectPresenter{}
}

func (p *ProjectPresenter) ShowCreateSuccess(project *entities.Project) {
	fmt.Printf("✅ Project '%s' has been created.\n", project.Name)
}

func (p *ProjectPresenter) ShowProjects(projects []entities.Project) {
	if len(projects) == 0 {
		fmt.Print("No projects found. Please create a new project to get started.")
		return
	}
	// 名前順でソート
	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].Name < projects[j].Name
	})

	for _, project := range projects {
		desc := "(No description)"
		if project.Description != nil {
			desc = *project.Description
		}
		fmt.Printf("- %s: %s\n", project.Name, desc)
	}
}

func (p *ProjectPresenter) ShowError(err error) {
	fmt.Printf("❌ Error: %v\n", err)
	os.Exit(1)
}
