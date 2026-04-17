package repository

import (
	"errors"
	"github.com/dwchwang/rest_api_golang/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	FindAll() ([]models.Book, error)
	FindById(id uint) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository{
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) Create(book *models.Book) error{
	return r.db.Create(book).Error
}

func (r *bookRepository) FindAll() ([]models.Book, error){
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) FindById(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error

	if err != nil {
		// Bắt đích danh lỗi không tìm thấy của GORM
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("không tìm thấy sách với ID này")
		}
		return nil, err
	}
	return &book, nil
}

// Cập nhật thông tin sách
func (r *bookRepository) Update(book *models.Book) error {
	// GORM hàm Save() sẽ cập nhật toàn bộ các trường của model.
	// Nếu bản ghi chưa có ID, nó sẽ Insert. Nếu có ID rồi, nó sẽ Update.
	return r.db.Save(book).Error
}

// Xóa sách theo ID
func (r *bookRepository) Delete(id uint) error {
	// GORM cần biết bạn muốn xóa trên bảng nào, nên ta truyền &models.Book{} vào làm mẫu
	return r.db.Delete(&models.Book{}, id).Error
}