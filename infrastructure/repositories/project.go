package repositories

import (
	"fmt"
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"
	"time-tracker/infrastructure/database/model"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) interfaces.IProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(project *entities.Project) error {
	return r.db.Create(&model.Project{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}).Error
}

func (r *ProjectRepository) FindAllActive() ([]entities.Project, error) {
	var projectModels []model.Project
	err := r.db.Where("archived_at IS NULL").Find(&projectModels).Error
	if err != nil {
		return nil, err
	}

	var projects []entities.Project
	for _, model := range projectModels {
		projects = append(projects, entities.Project{
			ID:          model.ID,
			Name:        model.Name,
			Description: model.Description,
			CreatedAt:   model.CreatedAt,
			UpdatedAt:   model.UpdatedAt,
		})
	}
	return projects, nil
}

func (r *ProjectRepository) IsNameDuplicated(name string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Project{}).
		Where("name = ?", name).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *ProjectRepository) Update(project *entities.Project) error {
	updatedModel := model.Project{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		ArchivedAt:  project.ArchivedAt,
	}
	if err := r.db.Save(&updatedModel).Error; err != nil {
		return fmt.Errorf("failed to update time entry: %w", err)
	}

	return nil
}
