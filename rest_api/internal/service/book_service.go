package service

import (
	"errors"

	"github.com/dwchwang/rest_api_golang/internal/models"
	"github.com/dwchwang/rest_api_golang/internal/repository"
)

type BookService interface {
	CreateBook(book *models.Book) error
	GetAllBooks() ([]models.Book, error)
	GetBookById(id uint) (*models.Book, error)
	UpdateBook(id uint, book *models.Book) error
	DeleteBook(id uint) error
	GetUnborrowedBooks() ([]models.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService{
	return &bookService{
		repo: repo,
	}
}


func (s *bookService) CreateBook(book *models.Book) error{
	if book.Title == "" {
		return errors.New("Ten sach khong duoc de trong")
	}
	if book.Author == "" {
		return errors.New("Ten tacgia khong duoc de trong")
	}
	return s.repo.Create(book)
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookById(id uint) (*models.Book, error) {
	return s.repo.FindById(id)
}

func (s *bookService) UpdateBook(id uint, updatedData *models.Book) error {
	// 1. Kiểm tra xem sách có tồn tại không trước khi update
	existingBook, err := s.repo.FindById(id)
	if err != nil {
		return err 
	}

	// 2. Cập nhật dữ liệu mới vào bản ghi cũ
	existingBook.Title = updatedData.Title
	existingBook.Author = updatedData.Author

	// 3. Lưu xuống DB
	return s.repo.Update(existingBook)
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}

func (s *bookService) GetUnborrowedBooks() ([]models.Book, error) {
	return s.repo.FindUnborrowedBooks()
}