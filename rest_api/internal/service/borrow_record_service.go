package service

import (
	"github.com/dwchwang/rest_api_golang/internal/models"
	"github.com/dwchwang/rest_api_golang/internal/repository"
)

type BorrowRecordService interface {
	BorrowBook(borrowerID, bookID uint) error
	ReturnBook(recordID uint) error
}

type borrowRecordService struct {
	repo repository.BorrowRecordRepository
}

func NewBorrowRecordService(repo repository.BorrowRecordRepository) BorrowRecordService {
	return &borrowRecordService{repo: repo}
}

func (s *borrowRecordService) BorrowBook(borrowerID, bookID uint) error {
	record := &models.BorrowRecord{
		BorrowerID: borrowerID,
		BookID:     bookID,
	}
	return s.repo.BorrowBook(record)
}

func (s *borrowRecordService) ReturnBook(recordID uint) error {
	return s.repo.ReturnBook(recordID)
}