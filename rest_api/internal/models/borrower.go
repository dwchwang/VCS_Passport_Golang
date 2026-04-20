package models

import "time"

type Borrower struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(150);not null" json:"name"`
	Phone     string    `gorm:"type:varchar(20);uniqueIndex" json:"phone"`
	Gender    string    `gorm:"type:varchar(10)" json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
