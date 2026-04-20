package handler

import (
	"net/http"
	"task-management-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{ service service.UserService }

func NewUserHandler(s service.UserService) *UserHandler { return &UserHandler{s} }

type registerRequest struct {
	Name     string `json:"name"     binding:"required,min=2,max=100"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.service.Register(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, user, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"user":         user,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetUint("userID")
	user, err := h.service.GetProfile(c.Request.Context(), userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user)
}