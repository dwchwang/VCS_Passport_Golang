package app

import (
	"github.com/dwchwang/rest_api_golang/internal/handler"
	"github.com/dwchwang/rest_api_golang/internal/repository"
	"github.com/dwchwang/rest_api_golang/internal/routes"
	"github.com/dwchwang/rest_api_golang/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Khởi tạo toàn bộ ứng dụng và trả về Engine của Gin
func SetupApp(db *gorm.DB) *gin.Engine {
	// module book
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)
	bookRoutes := routes.NewBookRoutes(bookHandler)

	// module borrower
	borrowerRepo := repository.NewBorrowerRepository(db)
	borrowerService := service.NewBorrowerService(borrowerRepo)
	borrowerHandler := handler.NewBorrowerHandler(borrowerService)
	borrowerRoutes := routes.NewBorrowerRoutes(borrowerHandler)

	// module borrow record
	borrowRecordRepo := repository.NewBorrowRecordRepository(db)
	borrowRecordService := service.NewBorrowRecordService(borrowRecordRepo)
	borrowRecordHandler := handler.NewBorrowRecordHandler(borrowRecordService)
	borrowRecordRoutes := routes.NewBorrowRecordRoutes(borrowRecordHandler)


	router := gin.Default()
	// Đăng ký toàn bộ vào router
	routes.RegisterRoutes(router, bookRoutes, borrowerRoutes, borrowRecordRoutes) // phẩy thêm userRoutes, v.v.

	return router
}