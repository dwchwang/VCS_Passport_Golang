package service

import (
	"context"
	"task-management-api/internal/models"
	"task-management-api/internal/repository"
	"task-management-api/pkg/apperror"
)

type ProjectService interface {
	CreateProject(ctx context.Context, ownerID uint, name, description string) (*models.Project, error)
	GetProject(ctx context.Context, ownerID, projectID uint) (*models.Project, error)
	ListProjects(ctx context.Context, ownerID uint) ([]models.Project, int64, error)
	UpdateProject(ctx context.Context, ownerID, projectID uint, name, description string) (*models.Project, error)
	DeleteProject(ctx context.Context, ownerID, projectID uint) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &projectService{projectRepo}
}

func (s *projectService) CreateProject(ctx context.Context, ownerID uint, name, description string) (*models.Project, error) {
	project := &models.Project{
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}
	if err := s.projectRepo.Create(ctx, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (s *projectService) GetProject(ctx context.Context, ownerID, projectID uint) (*models.Project, error) {
	project, err := s.projectRepo.FindByID(ctx, projectID)
	if err != nil {
		return nil, err
	}
	if project.OwnerID != ownerID {
		return nil, apperror.ErrForbidden
	}
	return project, nil
}

func (s *projectService) ListProjects(ctx context.Context, ownerID uint) ([]models.Project, int64, error) {
	return s.projectRepo.FindByOwnerID(ctx, ownerID)
}

func (s *projectService) UpdateProject(ctx context.Context, ownerID, projectID uint, name, description string) (*models.Project, error) {
	project, err := s.GetProject(ctx, ownerID, projectID)
	if err != nil {
		return nil, err
	}
	if name != "" {
		project.Name = name
	}
	if description != "" {
		project.Description = description
	}
	if err := s.projectRepo.Update(ctx, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (s *projectService) DeleteProject(ctx context.Context, ownerID, projectID uint) error {
	if _, err := s.GetProject(ctx, ownerID, projectID); err != nil {
		return err
	}
	return s.projectRepo.Delete(ctx, projectID)
}
