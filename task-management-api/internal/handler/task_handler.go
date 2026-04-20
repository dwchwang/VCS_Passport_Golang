package handler

import (
	"net/http"
	"strconv"
	"task-management-api/internal/models"
	"task-management-api/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct{ service service.TaskService }

func NewTaskHandler(s service.TaskService) *TaskHandler { return &TaskHandler{s} }

type createTaskRequest struct {
	Title       string              `json:"title"       binding:"required,min=1,max=255"`
	Description string              `json:"description" binding:"omitempty,max=2000"`
	Priority    models.TaskPriority `json:"priority"    binding:"omitempty,oneof=low medium high"`
	AssigneeID  *uint               `json:"assignee_id" binding:"omitempty"`
}

type updateTaskRequest struct {
	Title       string              `json:"title"       binding:"omitempty,min=1,max=255"`
	Description string              `json:"description" binding:"omitempty,max=2000"`
	Status      models.TaskStatus   `json:"status"      binding:"omitempty,oneof=todo in_progress done"`
	Priority    models.TaskPriority `json:"priority"    binding:"omitempty,oneof=low medium high"`
	AssigneeID  *uint               `json:"assignee_id" binding:"omitempty"`
}

func getProjectID(c *gin.Context) uint {
	id, _ := strconv.ParseUint(c.Param("projectId"), 10, 32)
	return uint(id)
}

func getTaskID(c *gin.Context) uint {
	id, _ := strconv.ParseUint(c.Param("taskId"), 10, 32)
	return uint(id)
}

func (h *TaskHandler) Create(c *gin.Context) {
	var req createTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := h.service.CreateTask(c.Request.Context(), c.GetUint("userID"), getProjectID(c), service.CreateTaskInput{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		AssigneeID:  req.AssigneeID,
	})
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) List(c *gin.Context) {
	page, _  := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	status   := c.Query("status")
	priority := c.Query("priority")

	tasks, total, err := h.service.ListTasks(c.Request.Context(), c.GetUint("userID"), getProjectID(c), page, limit, status, priority)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *TaskHandler) Get(c *gin.Context) {
	task, err := h.service.GetTask(c.Request.Context(), c.GetUint("userID"), getProjectID(c), getTaskID(c))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Update(c *gin.Context) {
	var req updateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := h.service.UpdateTask(c.Request.Context(), c.GetUint("userID"), getProjectID(c), getTaskID(c), service.UpdateTaskInput{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		AssigneeID:  req.AssigneeID,
	})
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Delete(c *gin.Context) {
	if err := h.service.DeleteTask(c.Request.Context(), c.GetUint("userID"), getProjectID(c), getTaskID(c)); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}