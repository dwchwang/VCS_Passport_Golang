package repository

import (
	"context"
	"errors"
	"task-management-api/internal/models"
	"task-management-api/pkg/apperror"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *models.Project) error
	FindByID(ctx context.Context, id uint) (*models.Project, error)
	FindByOwnerID(ctx context.Context, ownerID uint) ([]models.Project, int64, error)
	Update(ctx context.Context, project *models.Project) error
	Delete(ctx context.Context, id uint) error
}

type projectRepository struct{ db *gorm.DB }

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(ctx context.Context, project *models.Project) error {
	err := r.db.WithContext(ctx).Create(project).Error
	if err != nil {
		return apperror.ErrInternalServer
	}
	return nil
}

func (r *projectRepository) FindByID(ctx context.Context, id uint) (*models.Project, error) {
	var project models.Project
	err := r.db.WithContext(ctx).First(&project, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return &project, err
}

func (r *projectRepository) FindByOwnerID(ctx context.Context, ownerID uint) ([]models.Project, int64, error) {
	var projects []models.Project
	var count int64
	r.db.WithContext(ctx).
		Model(&models.Project{}).
		Where("owner_id = ?", ownerID).
		Count(&count)
	err := r.db.WithContext(ctx).
		Where("owner_id = ?", ownerID).
		Order("create_at DESC").
		Find(&projects).Error
	return projects, count, err
}

func (r *projectRepository) Update(ctx context.Context, project *models.Project) error {
	err := r.db.WithContext(ctx).Save(project).Error
	if err != nil {
		return apperror.ErrInternalServer
	}
	return nil
}

func (r *projectRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).Delete(&models.Project{}, id).Error
	if err != nil {
		return apperror.ErrInternalServer
	}
	return nil
}