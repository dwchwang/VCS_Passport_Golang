package main

import (
	"fmt"
	"log"

	"github.com/dwchwang/rest_api_golang/config"
	"github.com/dwchwang/rest_api_golang/internal/app"
)
func main() {

	config.LoadEnv()
	db := config.ConnectDatabase()

	// Giao toàn bộ việc lắp ráp cho Container
	router := app.SetupApp(db) 

	port := config.GetEnv("PORT", "8080")
	fmt.Printf("🚀 Server đang chạy tại http://localhost:%s\n", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Lỗi khởi chạy server:", err)
	}
}
