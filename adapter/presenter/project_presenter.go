
package presenter

import (
    "fmt"
    "time-tracker/domain/entities"
)

type ProjectPresenter struct{}

func NewProjectPresenter() *ProjectPresenter {
    return &ProjectPresenter{}
}

func (p *ProjectPresenter) ShowCreateSuccess(project *entities.Project) {
    fmt.Printf("✅ プロジェクト '%s' が作成されました。（作成日: %s）\n", project.Name, project.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (p *ProjectPresenter) ShowProjects(projects []entities.Project) {
    fmt.Println("📋 プロジェクト一覧:")
    for _, project := range projects {
        fmt.Printf("- %s: %s（作成日: %s, 更新日: %s）\n",
            project.Name,
            project.Description,
            project.CreatedAt.Format("2006-01-02 15:04:05"),
            project.UpdatedAt.Format("2006-01-02 15:04:05"))
    }
}

func (p *ProjectPresenter) ShowError(err error) {
    fmt.Printf("❌ エラー: %v\n", err)
}
