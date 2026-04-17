package repository

import (
	"github.com/dwchwang/rest_api_golang/internal/models"
	"gorm.io/gorm"
)

type BorrowerRepository interface {
	Create(borrower *models.Borrower) error
	FindAll() ([]models.Borrower, error)
	FindByID(id uint) (*models.Borrower, error)
	Update(borrower *models.Borrower) error
	Delete(id uint) error
}

type borrowerRepository struct {
	db *gorm.DB
}

func NewBorrowerRepository(db *gorm.DB) BookRepository {
	return &BorrowerRepository{
		db: db,
	}
}


