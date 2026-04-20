package service

import (
	"context"
	"task-management-api/internal/models"
	"task-management-api/internal/repository"
	"task-management-api/pkg/apperror"
	pkgjwt "task-management-api/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, name, email, password string) (*models.User, error)
	Login(ctx context.Context, email, password string) (string, *models.User, error)
	GetProfile(ctx context.Context, id uint) (*models.User, error)
}

type userService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo repository.UserRepository, jwtSecret string) UserService {
	return &userService{userRepo, jwtSecret}
}

func (s *userService) Register(ctx context.Context, name, email, password string) (*models.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, apperror.ErrInternalServer
	}
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(ctx context.Context, email, password string) (string, *models.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", nil, apperror.ErrInvalidCreds
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, apperror.ErrInvalidCreds
	}
	token, err := pkgjwt.Generate(user.ID, s.jwtSecret)
	if err != nil {
		return "", nil, apperror.ErrInternalServer
	}
	return token, user, nil
}

func (s *userService) GetProfile(ctx context.Context, id uint) (*models.User, error) {
	return s.userRepo.FindByID(ctx, id)
}