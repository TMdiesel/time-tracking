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
	Short: "A CLI for managing projects and tasks",
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
	taskController := controller.NewTaskController(taskUsecase, projectUsecase, taskPres)

	// --- Project Commands ---
	projectCmd := &cobra.Command{
		Use:   "project",
		Short: "Manage projects",
	}

	// project create
	createProjectCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			description, _ := cmd.Flags().GetString("description")
			var descPtr *string
			if description != "" {
				descPtr = &description
			}
			projectController.CreateProject(name, descPtr)
		},
	}
	createProjectCmd.Flags().StringP("name", "n", "", "Project name (required)")
	createProjectCmd.Flags().StringP("description", "d", "", "Project description (optional)")
	createProjectCmd.MarkFlagRequired("name")
	createProjectCmd.Flags().SortFlags = false

	// project list
	listProjectCmd := &cobra.Command{
		Use:   "list",
		Short: "List all projects",
		Run: func(cmd *cobra.Command, args []string) {
			projectController.ListProjects()
		},
	}

	// --- Task Commands ---
	taskCmd := &cobra.Command{
		Use:   "task",
		Short: "Manage tasks",
	}

	// task create
	createTaskCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new task",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			description, _ := cmd.Flags().GetString("description")

			var descPtr *string
			if description != "" {
				descPtr = &description
			}

			taskController.CreateTask(name, descPtr)
		},
	}
	createTaskCmd.Flags().StringP("name", "n", "", "Task name (required)")
	createTaskCmd.Flags().StringP("description", "d", "", "Task description (optional)")
	createTaskCmd.MarkFlagRequired("project")
	createTaskCmd.MarkFlagRequired("name")
	createTaskCmd.Flags().SortFlags = false

	// task list
	listTaskCmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			taskController.ListTasks()
		},
	}

	// --- コマンド登録 ---
	projectCmd.AddCommand(createProjectCmd)
	projectCmd.AddCommand(listProjectCmd)
	taskCmd.AddCommand(createTaskCmd)
	taskCmd.AddCommand(listTaskCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(taskCmd)

	// 実行
	rootCmd.Execute()
}
