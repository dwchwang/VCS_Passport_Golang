package database

import (
	"context"
	"log"

	"time"
	"task-management-api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(cfg *config.Config) *gorm.DB {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Đổi thành logger.Silent khi lên Production
	}
	db, err := gorm.Open(postgres.Open(cfg.DSN()), gormConfig)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB: ", err)
	}
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

	log.Println("database connected successfully")
	return db
}
