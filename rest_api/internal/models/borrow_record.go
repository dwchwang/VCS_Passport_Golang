package models

import "time"

type BorrowRecord struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	BorrowerID uint       `gorm:"not null" json:"borrower_id"`
	BookID     uint       `gorm:"not null" json:"book_id"`
	BorrowDate time.Time  `gorm:"autoCreateTime" json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date"`
	Status     string     `gorm:"type:varchar(50);default:'BORROWING'" json:"status"`
}