
package repositories

import (
    "gorm.io/gorm"
    "time-tracker/domain/entities"
    "time-tracker/domain/interfaces"
    "time-tracker/infrastructure/database/model"
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

func (r *ProjectRepository) FindAll() ([]entities.Project, error) {
    var projectModels []model.Project
    err := r.db.Find(&projectModels).Error
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
