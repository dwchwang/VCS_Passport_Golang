package routes

import (
	"task-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type publicRoutes struct{ h *handler.Handlers }

func newPublicRoutes(h *handler.Handlers) Route {
	return &publicRoutes{h}
}

func (r *publicRoutes) Register(rg *gin.RouterGroup) {
	rg.POST("/register", r.h.User.Register)
	rg.POST("/login",    r.h.User.Login)
}