package routes

import (
	"task-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type privateRoutes struct{ h *handler.Handlers }

func newPrivateRoutes(h *handler.Handlers) Route {
	return &privateRoutes{h}
}

func (r *privateRoutes) Register(rg *gin.RouterGroup) {
	// User
	rg.GET("/profile", r.h.User.GetProfile)

	// Projects — đổi :id thành :projectId
	projects := rg.Group("/projects")
	{
		projects.POST("",                r.h.Project.Create)
		projects.GET("",                 r.h.Project.List)
		projects.GET("/:projectId",      r.h.Project.Get)
		projects.PUT("/:projectId",      r.h.Project.Update)
		projects.DELETE("/:projectId",   r.h.Project.Delete)
	}

	// Tasks — giữ nguyên
	tasks := rg.Group("/projects/:projectId/tasks")
	{
		tasks.POST("",           r.h.Task.Create)
		tasks.GET("",            r.h.Task.List)
		tasks.GET("/:taskId",    r.h.Task.Get)
		tasks.PUT("/:taskId",    r.h.Task.Update)
		tasks.DELETE("/:taskId", r.h.Task.Delete)
	}
}