package repository

import (
	"errors"
	"time"

	"github.com/dwchwang/rest_api_golang/internal/models"
	"gorm.io/gorm"
)

type BorrowRecordRepository interface {
	BorrowBook(record *models.BorrowRecord) error
	ReturnBook(recordID uint) error
}

type borrowRecordRepository struct {
	db *gorm.DB
}

func NewBorrowRecordRepository(db *gorm.DB) BorrowRecordRepository {
	return &borrowRecordRepository{db: db}
}

func (r *borrowRecordRepository) BorrowBook(record *models.BorrowRecord) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var book models.Book
		if err := tx.First(&book, record.BookID).Error; err != nil {
			return errors.New("Book not found")
		}
		if book.Status == "BORROWED" {
			return errors.New("Book is already borrowed")
		}
		book.Status = "BORROWED"
		if err := tx.Save(&book).Error; err != nil {
			return err
		}

		record.BorrowDate = time.Now()
		record.Status = "BORROWED"
		if err := tx.Create(record).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *borrowRecordRepository) ReturnBook(recordID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Tim phieu muon
		var record models.BorrowRecord
		if err := tx.First(&record, recordID).Error; err != nil {
			return errors.New("Borrow record not found")
		}
		if record.Status == "RETURNED" {
			return errors.New("Book is already returned")
		}

		// cap nhat phieu muon
		record.Status = "RETURNED"
		record.ReturnDate = time.Now()
		if err := tx.Save(&record).Error; err != nil {
			return err
		}

		// cap nhat lai trang thai sach
		var book models.Book
		if err := tx.First(&book, record.BookID).Error; err != nil {
			return err
		}
		book.Status = "AVAILABLE"
		if err := tx.Save(&book).Error; err != nil {
			return err
		}
		return nil
	})
}
