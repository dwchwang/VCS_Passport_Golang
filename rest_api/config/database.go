package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := GetEnv("DB_DSN", "")
	if dsn == "" {
		log.Fatal("Error: Chưa cấu hình DB_DSN trong file .env")
	}

	// 1. Cấu hình Logger (Hiện SQL log trong lúc code)
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Đổi thành logger.Silent khi lên Production
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("Lỗi kết nối DB: ", err)
	}

	// 2. Lấy đối tượng sql.DB nguyên thủy để cấu hình Connection Pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Lỗi lấy đối tượng sql.DB: ", err)
	}

	// Cấu hình Pool
	sqlDB.SetMaxIdleConns(10)           // Số kết nối duy trì sẵn kể cả khi không ai dùng
	sqlDB.SetMaxOpenConns(100)          // Số kết nối tối đa được phép mở
	sqlDB.SetConnMaxLifetime(time.Hour) // Tuổi thọ tối đa của 1 kết nối (1 tiếng)

	// Kiểm tra Ping thử
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close()
		log.Fatal("Lỗi Ping DB: ", err)
	}

	fmt.Println("Connected to DB")

	// err = db.AutoMigrate(
	// 	&models.Book{},
	// 	&models.Borrower{},
	// 	&models.BorrowRecord{},
	// )
	// if err != nil {
	// 	log.Fatal("Lỗi Migrate: ", err)
	// }

	DB = db
	return db
}
