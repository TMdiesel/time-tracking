
package main

import (
    "time-tracker/adapter/controller"
    "time-tracker/adapter/presenter"
    "time-tracker/infrastructure/database"
    "time-tracker/infrastructure/repositories"
    "time-tracker/usecase"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "time-tracker",
    Short: "A CLI for managing projects",
}

func main() {
    db := database.NewDatabase()
    repo := repositories.NewProjectRepository(db)
    pres := presenter.NewProjectPresenter()
    usecase := usecase.NewProjectUsecase(repo)
    projectController := controller.NewProjectController(usecase, pres)

    rootCmd.AddCommand(&cobra.Command{
        Use:   "create [name] [description]",
        Short: "Create a new project",
        Args:  cobra.ExactArgs(2),
        Run: func(cmd *cobra.Command, args []string) {
            projectController.CreateProject(args[0], args[1])
        },
    })

    rootCmd.AddCommand(&cobra.Command{
        Use:   "list",
        Short: "List all projects",
        Run: func(cmd *cobra.Command, args []string) {
            projectController.ListProjects()
        },
    })

    rootCmd.Execute()
}
