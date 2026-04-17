package models

import "time"

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Author    string    `gorm:"type:varchar(255);not null" json:"author"`
	Status    string    `gorm:"type:varchar(50);default:'AVAILABLE'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}