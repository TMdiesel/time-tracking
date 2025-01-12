package repositories

import (
	"time-tracker/domain/dto"
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

func (r *TaskRepository) FindAllWithProjectName() ([]dto.TaskWithProjectDTO, error) {
	var tasks []model.Task

	err := r.db.Preload("Project").Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	// DTOに変換
	var result []dto.TaskWithProjectDTO
	for _, task := range tasks {
		result = append(result, dto.TaskWithProjectDTO{
			ID:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			ProjectID:   task.ProjectID,
			ProjectName: task.Project.Name,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return result, nil
}
