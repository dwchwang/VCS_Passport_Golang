package handler

import (
	"net/http"
	"strconv"
	"task-management-api/internal/service"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct{ service service.ProjectService }

func NewProjectHandler(s service.ProjectService) *ProjectHandler { return &ProjectHandler{s} }

type createProjectRequest struct {
	Name        string `json:"name"        binding:"required,min=2,max=255"`
	Description string `json:"description" binding:"omitempty,max=1000"`
}

type updateProjectRequest struct {
	Name        string `json:"name"        binding:"omitempty,min=2,max=255"`
	Description string `json:"description" binding:"omitempty,max=1000"`
}

func (h *ProjectHandler) Create(c *gin.Context) {
	var req createProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project, err := h.service.CreateProject(c.Request.Context(), c.GetUint("userID"), req.Name, req.Description)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, project)
}

func (h *ProjectHandler) List(c *gin.Context) {
	projects, total, err := h.service.ListProjects(c.Request.Context(), c.GetUint("userID"))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
		"total":    total,
	})
}

func (h *ProjectHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("projectId"), 10, 32)
	project, err := h.service.GetProject(c.Request.Context(), c.GetUint("userID"), uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) Update(c *gin.Context) {
	var req updateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.ParseUint(c.Param("projectId"), 10, 32)
	project, err := h.service.UpdateProject(c.Request.Context(), c.GetUint("userID"), uint(id), req.Name, req.Description)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("projectId"), 10, 32)
	if err := h.service.DeleteProject(c.Request.Context(), c.GetUint("userID"), uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "project deleted successfully"})
}