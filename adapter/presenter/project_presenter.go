
package presenter

import "fmt"

type ProjectPresenter struct{}

func NewProjectPresenter() *ProjectPresenter {
    return &ProjectPresenter{}
}

func (p *ProjectPresenter) ShowCreateSuccess(name string) {
    fmt.Printf("✅ Project '%s' was successfully created!\n", name)
}
