
package repositories

import (
    "database/sql"
    "time-tracker/domain/entities"
    "time-tracker/domain/interfaces"
)

type ProjectRepository struct {
    db *sql.DB
}

func NewProjectRepository(db *sql.DB) interfaces.IProjectRepository {
    return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Save(project *entities.Project) error {
    query := `INSERT INTO projects (name, description, created_at, updated_at) VALUES (?, ?, ?, ?)`
    _, err := r.db.Exec(query, project.Name, project.Description, project.CreatedAt, project.UpdatedAt)
    return err
}
