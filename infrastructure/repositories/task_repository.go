package repositories

import (
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"
	"time-tracker/infrastructure/database/model"

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
