
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time-tracker/adapter/controller"
    "time-tracker/infrastructure/repositories"
    "time-tracker/adapter/presenter"
    "time-tracker/usecase"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    if len(os.Args) < 3 {
        log.Fatal("Usage: go run main.go <ProjectName> <Description>")
    }

    projectName := os.Args[1]
    projectDescription := os.Args[2]

    db, err := sql.Open("sqlite3", "./time_tracker.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    repo := repositories.NewProjectRepository(db)
    successPresenter := presenter.NewProjectPresenter()
    errorPresenter := presenter.NewErrorPresenter()

    projectUsecase := usecase.NewProjectUseCase(repo)
    projectController := controller.NewProjectController(projectUsecase, successPresenter, errorPresenter)

    projectController.CreateProject(projectName, projectDescription)

    fmt.Println("🚀 Project creation completed!")
}
