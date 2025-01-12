package repositories

import (
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"
	"time-tracker/infrastructure/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) interfaces.ITaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *entities.Task) error {
	return r.db.Create(&model.Task{
		ID:          task.ID,
		ProjectID:   task.ProjectID,
		Name:        task.Name,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}).Error
}

func (r *TaskRepository) IsNameDuplicated(projectID uuid.UUID, name string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Task{}).Where("project_id = ? AND name = ?", projectID, name).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *TaskRepository) FindAll() ([]entities.Task, error) {
	var taskModels []model.Task
	err := r.db.Find(&taskModels).Error
	if err != nil {
		return nil, err
	}

	var tasks []entities.Task
	for _, model := range taskModels {
		tasks = append(tasks, entities.Task{
			ID:          model.ID,
			ProjectID:   model.ProjectID,
			Name:        model.Name,
			Description: model.Description,
			CreatedAt:   model.CreatedAt,
			UpdatedAt:   model.UpdatedAt,
		})
	}
	return tasks, nil
}
