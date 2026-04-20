package repository

import (
	"context"
	"errors"
	"strings"
	"task-management-api/internal/models"
	"task-management-api/pkg/apperror"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id uint) (*models.User, error)
}

type userRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return apperror.ErrEmailExists
		}
		return apperror.ErrInternalServer
	}
	return nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return &user, err
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrNotFound
	}
	return &user, err
}