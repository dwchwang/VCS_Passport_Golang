package models

import "time"

type BorrowRecord struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	BorrowerID uint      `json:"borrower_id"`
	BookID     uint      `json:"book_id"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"` // Thời gian trả sách (có thể null nếu chưa trả)
	Status     string    `gorm:"type:varchar(20);default:'BORROWED'" json:"status"` // BORROWED hoặc RETURNED
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}