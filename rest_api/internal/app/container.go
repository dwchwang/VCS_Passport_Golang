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
	// 1. Khởi tạo tất cả Repositories
	bookRepo := repository.NewBookRepository(db)
	// userRepo := repository.NewUserRepository(db)
	// borrowRepo := repository.NewBorrowRepository(db)

	// 2. Khởi tạo tất cả Services
	bookService := service.NewBookService(bookRepo)
	// userService := service.NewUserService(userRepo)
	// borrowService := service.NewBorrowService(borrowRepo, bookRepo, userRepo)

	// 3. Khởi tạo tất cả Handlers
	bookHandler := handler.NewBookHandler(bookService)
	// userHandler := handler.NewUserHandler(userService)

	// 4. Lắp vào Routes
	bookRoutes := routes.NewBookRoutes(bookHandler)
	// userRoutes := routes.NewUserRoutes(userHandler)

	router := gin.Default()
	
	// Đăng ký toàn bộ vào router
	routes.RegisterRoutes(router, bookRoutes) // phẩy thêm userRoutes, v.v.

	return router
}