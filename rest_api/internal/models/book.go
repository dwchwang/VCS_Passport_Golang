package models

import "time"

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null;index" json:"title"` //index for faster search by title
	Author    string    `gorm:"type:varchar(255);not null;index" json:"author"` //index for faster search by author
	Status    string    `gorm:"type:varchar(50);default:'AVAILABLE'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}