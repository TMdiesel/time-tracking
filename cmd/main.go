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
	timeEntryRepo := repositories.NewTimeEntryRepository(db)
	timeEntryPres := presenter.NewTimeEntryPresenter()
	timeEntryService := service.NewTimeEntryService(taskRepo)

	taskUsecase := usecase.NewTaskUsecase(taskRepo, timeEntryRepo, taskService, timeEntryService)
	taskController := controller.NewTaskController(taskUsecase, projectUsecase, taskPres)
	timeEntryUsecase := usecase.NewTimeEntryUsecase(timeEntryRepo)
	timeEntryController := controller.NewTimeEntryController(timeEntryUsecase, timeEntryPres)

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

	// task start
	startTaskCmd := &cobra.Command{
		Use:   "start",
		Short: "start a task",
		Run: func(cmd *cobra.Command, args []string) {
			taskController.StartTask()
		},
	}

	// task end
	endTaskCmd := &cobra.Command{
		Use:   "end",
		Short: "end a task",
		Run: func(cmd *cobra.Command, args []string) {
			taskController.EndTask()
		},
	}

	// --- TimeEntry Commands ---
	timeEntryCmd := &cobra.Command{
		Use:   "entry",
		Short: "Manage time entries",
	}

	// time entry list
	listTimeEntryCmd := &cobra.Command{
		Use:   "list",
		Short: "List time entries",
		Run: func(cmd *cobra.Command, args []string) {
			from, _ := cmd.Flags().GetString("from")
			to, _ := cmd.Flags().GetString("to")
			timeEntryController.ListTimeEntries(from, to)
		},
	}
	listTimeEntryCmd.Flags().StringP("from", "f", "", "Start date (YYYY-MM-DD) (optional)")
	listTimeEntryCmd.Flags().StringP("to", "t", "", "End date (YYYY-MM-DD) (optional)")
	listTimeEntryCmd.Flags().SortFlags = false

	// --- コマンド登録 ---
	projectCmd.AddCommand(createProjectCmd)
	projectCmd.AddCommand(listProjectCmd)
	taskCmd.AddCommand(createTaskCmd)
	taskCmd.AddCommand(listTaskCmd)
	taskCmd.AddCommand(startTaskCmd)
	taskCmd.AddCommand(endTaskCmd)
	timeEntryCmd.AddCommand(listTimeEntryCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(timeEntryCmd)

	// 実行
	rootCmd.Execute()
}
