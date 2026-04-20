package service

import (
	"context"
	"task-management-api/internal/models"
	"task-management-api/internal/repository"
	"task-management-api/pkg/apperror"
)

type TaskService interface {
	CreateTask(ctx context.Context, ownerID, projectID uint, req CreateTaskInput) (*models.Task, error)
	GetTask(ctx context.Context, ownerID, projectID, taskID uint) (*models.Task, error)
	ListTasks(ctx context.Context, ownerID, projectID uint, page, limit int, status, priority string) ([]models.Task, int64, error)
	UpdateTask(ctx context.Context, ownerID, projectID, taskID uint, req UpdateTaskInput) (*models.Task, error)
	DeleteTask(ctx context.Context, ownerID, projectID, taskID uint) error
}

// Input structs — dùng trong service thay vì request struct của handler
type CreateTaskInput struct {
	Title       string
	Description string
	Priority    models.TaskPriority
	AssigneeID  *uint
	DueDate     *string
}

type UpdateTaskInput struct {
	Title       string
	Description string
	Status      models.TaskStatus
	Priority    models.TaskPriority
	AssigneeID  *uint
	DueDate     *string
}

type taskService struct {
	taskRepo    repository.TaskRepository
	projectRepo repository.ProjectRepository
}

func NewTaskService(taskRepo repository.TaskRepository, projectRepo repository.ProjectRepository) TaskService {
	return &taskService{taskRepo, projectRepo}
}

func (s *taskService) checkOwner(ctx context.Context, ownerID, projectID uint) error {
	project, err := s.projectRepo.FindByID(ctx, projectID)
	if err != nil {
		return err
	}
	if project.OwnerID != ownerID {
		return apperror.ErrForbidden
	}
	return nil
}

func (s *taskService) CreateTask(ctx context.Context, ownerID, projectID uint, req CreateTaskInput) (*models.Task, error) {
	if err := s.checkOwner(ctx, ownerID, projectID); err != nil {
		return nil, err
	}
	priority := req.Priority
	if priority == "" {
		priority = models.PriorityMedium
	}
	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      models.StatusTodo,
		Priority:    priority,
		ProjectID:   projectID,
		AssigneeID:  req.AssigneeID,
	}
	if err := s.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskService) GetTask(ctx context.Context, ownerID, projectID, taskID uint) (*models.Task, error) {
	if err := s.checkOwner(ctx, ownerID, projectID); err != nil {
		return nil, err
	}
	task, err := s.taskRepo.FindByID(ctx, taskID)
	if err != nil {
		return nil, err
	}
	if task.ProjectID != projectID {
		return nil, apperror.ErrNotFound
	}
	return task, nil
}

func (s *taskService) ListTasks(ctx context.Context, ownerID, projectID uint, page, limit int, status, priority string) ([]models.Task, int64, error) {
	if err := s.checkOwner(ctx, ownerID, projectID); err != nil {
		return nil, 0, err
	}
	return s.taskRepo.FindByProjectID(ctx, projectID, page, limit, status, priority)
}

func (s *taskService) UpdateTask(ctx context.Context, ownerID, projectID, taskID uint, req UpdateTaskInput) (*models.Task, error) {
	task, err := s.GetTask(ctx, ownerID, projectID, taskID)
	if err != nil {
		return nil, err
	}
	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if req.Status != "" {
		task.Status = req.Status
	}
	if req.Priority != "" {
		task.Priority = req.Priority
	}
	if req.AssigneeID != nil {
		task.AssigneeID = req.AssigneeID
	}
	if err := s.taskRepo.Update(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskService) DeleteTask(ctx context.Context, ownerID, projectID, taskID uint) error {
	if _, err := s.GetTask(ctx, ownerID, projectID, taskID); err != nil {
		return err
	}
	return s.taskRepo.Delete(ctx, taskID)
}