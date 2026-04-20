package service

import (
	"errors"

	"github.com/dwchwang/rest_api_golang/internal/models"
	"github.com/dwchwang/rest_api_golang/internal/repository"
)

type BorrowerService interface {
	CreateBorrower(borrower *models.Borrower) error
	GetAllBorrowers() ([]models.Borrower, error)
	GetBorrowerById(id uint) (*models.Borrower, error)
	UpdateBorrower(id uint, borrower *models.Borrower) error
	DeleteBorrower(id uint) error
}

type borrowerService struct {
	repo repository.BorrowerRepository
}

func NewBorrowerService(repo repository.BorrowerRepository) BorrowerService {
	return &borrowerService{repo: repo}
}

func (s *borrowerService) CreateBorrower(borrower *models.Borrower) error {
	if borrower.Name == "" {
		return errors.New("tên người mượn không được để trống")
	}
	if borrower.Phone == "" {
		return errors.New("số điện thoại không được để trống")
	}

	return s.repo.Create(borrower)
}

func (s *borrowerService) GetAllBorrowers() ([]models.Borrower, error) {
	return s.repo.FindAll()
}

func (s *borrowerService) GetBorrowerById(id uint) (*models.Borrower, error) {
	return s.repo.FindById(id)
}

func (s *borrowerService) UpdateBorrower(id uint, updatedData *models.Borrower) error {
	existingBorrower, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	existingBorrower.Name = updatedData.Name
	existingBorrower.Phone = updatedData.Phone
	existingBorrower.Gender = updatedData.Gender

	return s.repo.Update(existingBorrower)
}

func (s *borrowerService) DeleteBorrower(id uint) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
