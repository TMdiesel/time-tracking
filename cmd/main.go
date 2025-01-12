package main

import (
	"time-tracker/adapter/controller"
	"time-tracker/adapter/presenter"
	"time-tracker/domain/service"
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
	projectRepo := repositories.NewProjectRepository(db)
	projectPres := presenter.NewProjectPresenter()
	projectService := service.NewProjectService(projectRepo)
	projectUsecase := usecase.NewProjectUsecase(projectRepo, projectService)
	projectController := controller.NewProjectController(projectUsecase, projectPres)
	taskRepo := repositories.NewTaskRepository(db)
	taskPres := presenter.NewTaskPresenter()
	taskService := service.NewTaskService(taskRepo)
	taskUsecase := usecase.NewTaskUsecase(taskRepo, taskService)
	taskController := controller.NewTaskController(taskUsecase, taskPres)

	// project
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

	// task
	rootCmd.AddCommand(&cobra.Command{
		Use:   "create_task [project_id] [name] [[description]]",
		Short: "Create a new task",
		Args:  cobra.RangeArgs(2, 3),
		Run: func(cmd *cobra.Command, args []string) {
			var description *string
			if len(args) == 3 {
				description = &args[2]
			}
			taskController.CreateTask(args[0], args[1], description)
		},
	})

	rootCmd.Execute()
}
