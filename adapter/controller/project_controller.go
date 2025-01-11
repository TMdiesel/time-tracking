
package controller

import (
    "time-tracker/domain/interfaces"
    "time-tracker/usecase"
)

type ProjectController struct {
    useCase      usecase.IProjectUseCase
    successPres  interfaces.IProjectPresenter
    errorPres    interfaces.IErrorPresenter
}

func NewProjectController(useCase usecase.IProjectUseCase, success interfaces.IProjectPresenter, error interfaces.IErrorPresenter) *ProjectController {
    return &ProjectController{
        useCase:     useCase,
        successPres: success,
        errorPres:   error,
    }
}

func (c *ProjectController) CreateProject(name, description string) {
    err := c.useCase.AddProject(name, description)
    if err != nil {
        c.errorPres.ShowError(err)
        return
    }
    c.successPres.ShowCreateSuccess(name)
}
