package service

import (
	"errors"
	"time"
	"time-tracker/domain/entities"
	"time-tracker/domain/interfaces"

	"github.com/google/uuid"
)

type ProjectService struct {
	repo interfaces.IProjectRepository
}

func NewProjectService(repo interfaces.IProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) NewProject(name string, description *string) (*entities.Project, error) {
	if name == "" {
		return nil, errors.New("project name not specified")
	}
	isDuplicated, err := s.repo.IsNameDuplicated(name)
	if err != nil {
		return nil, err
	}
	if isDuplicated {
		return nil, errors.New("project with the same name already exists")
	}

	return &entities.Project{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
