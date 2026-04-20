package routes

import (
	"github.com/dwchwang/rest_api_golang/internal/handler"
	"github.com/gin-gonic/gin"
)

type BorrowRecordRoutes struct {
	handler *handler.BorrowRecordHandler
}

func NewBorrowRecordRoutes(handler *handler.BorrowRecordHandler) Route {
	return &BorrowRecordRoutes{
		handler: handler,
	}
}

func (br *BorrowRecordRoutes) Register(r *gin.RouterGroup) {
	records := r.Group("/borrows")
	{
		records.POST("", br.handler.BorrowBook)
		records.POST("/return/:id", br.handler.ReturnBook)
	}
}
