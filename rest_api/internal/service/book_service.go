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
		return err // Sẽ trả về lỗi "không tìm thấy sách" từ Repo
	}

	// 2. Cập nhật dữ liệu mới vào bản ghi cũ
	existingBook.Title = updatedData.Title
	existingBook.Author = updatedData.Author
	// Lưu ý: Không cập nhật lại ID hoặc Trạng thái mượn/trả ở hàm này

	// 3. Lưu xuống DB
	return s.repo.Update(existingBook)
}

func (s *bookService) DeleteBook(id uint) error {
	// Bạn có thể thêm logic: Kiểm tra xem sách có đang bị ai mượn không trước khi xóa
	// (Tạm thời chúng ta cứ gọi hàm xóa thẳng)
	return s.repo.Delete(id)
}