package routes

import (
	"task-management-api/internal/handler"
	"task-management-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// Route interface — mỗi group routes implement cái này
type Route interface {
	Register(rg *gin.RouterGroup)
}

// Đăng ký tất cả routes
func Setup(r *gin.Engine, h *handler.Handlers) {
	r.Use(middleware.ErrorHandler())

	api := r.Group("/api/v1")

	// Public routes
	newPublicRoutes(h).Register(api)

	// Private routes
	auth := api.Group("/")
	auth.Use(middleware.Auth())
	newPrivateRoutes(h).Register(auth)
}