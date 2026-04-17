package config

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)


func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: Không tìm thấy file .env, sẽ dùng biến môi trường hệ thống.")
	}
}

// GetEnv giúp lấy giá trị theo Key, nếu không có thì trả về giá trị mặc định (fallback)
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}