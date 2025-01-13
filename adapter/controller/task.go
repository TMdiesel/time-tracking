package controller

import (
	"fmt"
	"time-tracker/adapter/presenter"
	"time-tracker/domain/dto"
	"time-tracker/domain/entities"
	"time-tracker/usecase"

	"github.com/manifoldco/promptui"
)

type TaskController struct {
	taskUsecase    *usecase.TaskUsecase
	projectUsecase *usecase.ProjectUsecase
	presenter      *presenter.TaskPresenter
}

func NewTaskController(tu *usecase.TaskUsecase, pu *usecase.ProjectUsecase, p *presenter.TaskPresenter) *TaskController {
	return &TaskController{
		taskUsecase:    tu,
		projectUsecase: pu,
		presenter:      p,
	}
}

func (c *TaskController) CreateTask(name string, description *string) {
	// プロジェクト一覧の取得
	projects, err := c.projectUsecase.ListProjects()
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	if len(projects) == 0 {
		c.presenter.ShowError(fmt.Errorf("No projects available. Please create a project first"))
	}

	// インタラクティブにプロジェクト選択
	selectedProject, err := c.selectProjectInteractive(projects)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}

	// タスク作成
	task, err := c.taskUsecase.CreateTask(selectedProject.ID, name, description)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}

	// 成功メッセージ表示
	c.presenter.ShowCreateSuccess(task)
}

func (c *TaskController) ListTasks() {
	tasks, err := c.taskUsecase.ListTasks()
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	c.presenter.ShowTasks(tasks)
}

// インタラクティブにプロジェクト選択
func (c *TaskController) selectProjectInteractive(projects []entities.Project) (entities.Project, error) {
	projectNames := []string{}
	for _, project := range projects {
		projectNames = append(projectNames, project.Name)
	}

	// インタラクティブ選択
	prompt := promptui.Select{
		Label: "Select a Project",
		Items: projectNames,
	}
	index, _, err := prompt.Run()
	if err != nil {
		return entities.Project{}, fmt.Errorf("Project selection failed: %v", err)
	}

	return projects[index], nil
}

func (c *TaskController) StartTask() {
	// タスク一覧の取得
	tasks, err := c.taskUsecase.ListTasks()
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	if len(tasks) == 0 {
		c.presenter.ShowError(fmt.Errorf("No tasks available. Please create a task first"))
		return
	}

	// インタラクティブにタスク選択
	selectedTask, err := c.selectTaskInteractive(tasks)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}

	// タスク開始
	err = c.taskUsecase.StartTask(selectedTask.ID)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}

	c.presenter.ShowStartSuccess(selectedTask)
}

// インタラクティブにタスク選択
func (c *TaskController) selectTaskInteractive(tasks []dto.TaskWithProjectDTO) (dto.TaskWithProjectDTO, error) {
	taskNames := []string{}
	for _, task := range tasks {
		taskNames = append(taskNames, task.Name)
	}

	prompt := promptui.Select{
		Label: "Select a Task to Start",
		Items: taskNames,
	}
	index, _, err := prompt.Run()
	if err != nil {
		return dto.TaskWithProjectDTO{}, fmt.Errorf("Task selection failed: %v", err)
	}

	return tasks[index], nil
}

func (c *TaskController) StopTask() {
	err := c.taskUsecase.StopTask()
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	c.presenter.ShowStopSuccess()
}
