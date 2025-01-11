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
		Use:   "create_project [name] [[description]]",
		Short: "Create a new project",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			var description *string
			if len(args) == 2 {
				description = &args[1]
			}
			projectController.CreateProject(args[0], description)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "list_projects",
		Short: "List all projects",
		Run: func(cmd *cobra.Command, args []string) {
			projectController.ListProjects()
		},
	})

	rootCmd.Execute()
}
