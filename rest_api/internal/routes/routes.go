package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Register(r *gin.RouterGroup)
}

// Hàm này nhận vào Engine của Gin và một "danh sách vô hạn" các Route (Variadic parameter)
func RegisterRoutes(r *gin.Engine, routes ...Route) {
	api := r.Group("/api/v1")
	
	for _, route := range routes {
		route.Register(api)
	}
}