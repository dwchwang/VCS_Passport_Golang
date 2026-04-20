package repository

import (
	"context"
	"errors"
	"task-management-api/internal/models"
	"task-management-api/pkg/apperror"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	FindByID(ctx context.Context, id uint) (*models.Task, error)
	FindByProjectID(ctx context.Context, projectID uint, page, limit int, status, priority string) ([]models.Task, int64, error)
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id uint) error
}

type taskRepository struct{ db *gorm.DB }

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *taskRepository) FindByID(ctx context.Context, id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.WithContext(ctx).
		Preload("Assignee").
		First(&task, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return &task, err
}

func (r *taskRepository) FindByProjectID(ctx context.Context, projectID uint, page, limit int, status, priority string) ([]models.Task, int64, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	query := r.db.WithContext(ctx).
		Model(&models.Task{}).
		Where("project_id = ?", projectID)

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	var total int64
	query.Count(&total)

	var tasks []models.Task
	err := query.
		Preload("Assignee").
		Order("created_at DESC").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&tasks).Error

	return tasks, total, err
}

func (r *taskRepository) Update(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Save(task).Error
}

func (r *taskRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Task{}, id).Error
}