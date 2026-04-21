package routes

import (
	"github.com/dwchwang/rest_api_golang/internal/handler"
	"github.com/gin-gonic/gin"
)

// Struct chứa Handler của sách
type BookRoutes struct {
	handler *handler.BookHandler
}

// Khởi tạo BookRoutes
func NewBookRoutes(handler *handler.BookHandler) Route {
	return &BookRoutes{
		handler: handler,
	}
}

// Thực thi hàm Register của Interface Route
func (br *BookRoutes) Register(r *gin.RouterGroup) {
	// br.handler.GetAllBooks không có ngoặc tròn () ở cuối nhé!
	books := r.Group("/books")
	{
		books.POST("", br.handler.CreateBook)
		books.GET("", br.handler.GetAllBooks)
		books.GET("/unborrowed", br.handler.GetUnborrowedBooks)
		books.GET("/:id", br.handler.GetBookById)
		books.PUT("/:id", br.handler.UpdateBook)
		books.DELETE("/:id", br.handler.DeleteBook)
	}
}