package app

import (
	"log"
	"task-management-api/internal/handler"
	"task-management-api/internal/repository"
	"task-management-api/internal/routes"
	"task-management-api/internal/service"
	"task-management-api/pkg/config"
	"task-management-api/pkg/database"
	"task-management-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	config *config.Config
}

func New() *App {
	// 1. Load config
	cfg := config.Load()

	// 2. Kết nối DB
	db := database.NewPostgresDB(cfg)

	// 3. Set JWT secret cho middleware
	middleware.JWTSecret = cfg.JWTSecret

	// 4. Repositories
	userRepo    := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	taskRepo    := repository.NewTaskRepository(db)

	// 5. Services
	userService    := service.NewUserService(userRepo, cfg.JWTSecret)
	projectService := service.NewProjectService(projectRepo)
	taskService    := service.NewTaskService(taskRepo, projectRepo)

	// 6. Handlers
	handlers := &handler.Handlers{
		User:    handler.NewUserHandler(userService),
		Project: handler.NewProjectHandler(projectService),
		Task:    handler.NewTaskHandler(taskService),
	}

	// 7. Setup Gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 8. Setup routes
	routes.Setup(r, handlers)

	return &App{router: r, config: cfg}
}

func (a *App) Run() {
	log.Printf("Server running on http://localhost:%s", a.config.ServerPort)
	if err := a.router.Run(":" + a.config.ServerPort); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}