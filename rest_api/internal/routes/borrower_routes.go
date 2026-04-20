package routes

import (
	"github.com/dwchwang/rest_api_golang/internal/handler"
	"github.com/gin-gonic/gin"
)

type BorrowerRoutes struct {
	handler *handler.BorrowerHandler
}

func NewBorrowerRoutes(handler *handler.BorrowerHandler) Route {
	return &BorrowerRoutes{
		handler: handler,
	}
}

func (br *BorrowerRoutes) Register(r *gin.RouterGroup) {
	borrowers := r.Group("/borrowers")
	{
		borrowers.POST("", br.handler.CreateBorrower)
		borrowers.GET("", br.handler.GetAllBorrowers)
		borrowers.GET("/:id", br.handler.GetBorrowerById)
		borrowers.PUT("/:id", br.handler.UpdateBorrower)
		borrowers.DELETE("/:id", br.handler.DeleteBorrower)
	}
}