package routes

import "github.com/gin-gonic/gin"

// Định nghĩa Interface: Bất kỳ struct nào có hàm Register(r *gin.RouterGroup) đều được coi là một Route
type Route interface {
	Register(r *gin.RouterGroup)
}

// Hàm này nhận vào Engine của Gin và một "danh sách vô hạn" các Route (Variadic parameter)
func RegisterRoutes(r *gin.Engine, routes ...Route) {
	// Định nghĩa prefix chung cho toàn bộ API
	api := r.Group("/api/v1")

	// Lặp qua từng route được truyền vào và yêu cầu chúng tự đăng ký vào cái /api/v1 này
	for _, route := range routes {
		route.Register(api)
	}
}