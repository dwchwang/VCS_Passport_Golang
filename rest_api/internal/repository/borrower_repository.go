package repository

import (
	"errors"

	"github.com/dwchwang/rest_api_golang/internal/models"
	"gorm.io/gorm"
)

type BorrowerRepository interface {
	Create(borrower *models.Borrower) error
	FindAll() ([]models.Borrower, error)
	FindById(id uint) (*models.Borrower, error)
	Update(borrower *models.Borrower) error
	Delete(id uint) error
}

type borrowerRepository struct {
	db *gorm.DB
}

func NewBorrowerRepository(db *gorm.DB) BorrowerRepository {
	return &borrowerRepository{db: db}
}

func (r *borrowerRepository) Create(borrower *models.Borrower) error {
	return r.db.Create(borrower).Error
}

func (r *borrowerRepository) FindAll() ([]models.Borrower, error) {
	var borrowers []models.Borrower
	err := r.db.Find(&borrowers).Error
	return borrowers, err
}

func (r *borrowerRepository) FindById(id uint) (*models.Borrower, error) {
	var borrower models.Borrower
	err := r.db.First(&borrower, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("không tìm thấy người mượn với ID này")
		}
		return nil, err
	}
	return &borrower, nil
}

func (r *borrowerRepository) Update(borrower *models.Borrower) error {
	return r.db.Save(borrower).Error
}

func (r *borrowerRepository) Delete(id uint) error {
	return r.db.Delete(&models.Borrower{}, id).Error
}