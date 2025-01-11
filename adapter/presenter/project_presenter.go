package presenter

import (
	"fmt"
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
	// 作成日順でソート
	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].CreatedAt.Before(projects[j].CreatedAt)
	})

	fmt.Println("📋 Project List:")
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
}
