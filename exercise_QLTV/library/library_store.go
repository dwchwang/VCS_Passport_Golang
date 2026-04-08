package library

import (
	"fmt"
	"strings"
	"time"

	"dwchwang.com/exercise_qltv/models"
)

type Library struct {
	Books        map[string]models.Book
	Borrowers    map[string]models.Borrower
	Transactions map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		Books:        make(map[string]models.Book),
		Borrowers:    make(map[string]models.Borrower),
		Transactions: make(map[string]models.Transaction),
	}
}

func (lib *Library) AddBookStore(id, title, author string) error {
	if _, exists := lib.Books[id]; exists {
		return fmt.Errorf("Sach voi ID %s da ton tai", id)
	}
	lib.Books[id] = models.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}

	return nil
}

func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.Books))
	for _, book := range lib.Books {
		books = append(books, book)
	}
	return books
}

func (lib *Library) AddBorrowerStore(id, name, email string) error {
	if _, exists := lib.Books[id]; exists {
		return fmt.Errorf("Nguoi muon voi ID %s da ton tai", id)
	}
	lib.Borrowers[id] = models.Borrower{
		ID:    id,
		Name:  name,
		Email: email,
	}
	return nil
}

func (lib *Library) ListBorrowersStore() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.Borrowers))
	for _, borrower := range lib.Borrowers {
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func (lib *Library) BorrowBookStore(id, bookID, borrowerID string) error {
	book, bookExists := lib.Books[bookID]
	if !bookExists {
		return fmt.Errorf("Sach voi ID %s khong ton tai", bookID)
	}
	_, borrowerExists := lib.Borrowers[borrowerID]
	if !borrowerExists {
		return fmt.Errorf("Nguoi muon voi ID %s khong ton tai", borrowerID)
	}
	if book.IsBorrowed {
		return fmt.Errorf("Sach ten: %s da duoc muon", book.Title)
	}
	if _, exists := lib.Transactions[id]; exists {
		return fmt.Errorf("Giao dich voi ID %s da ton tai", id)
	}

	book.IsBorrowed = true
	lib.Books[bookID] = book
	lib.Transactions[id] = models.Transaction{
		ID:         id,
		BookID:     bookID,
		BorrowerID: borrowerID,
		BorrowDate: time.Now(),
	}

	return nil
}

func (lib *Library) ListBorrowHistoryByBorrowerStore(borrowerID string) []models.Transaction {
	if _, borrowerExists := lib.Borrowers[borrowerID]; !borrowerExists {
		return nil
	}
	history := make([]models.Transaction, 0)
	for _, transaction := range lib.Transactions {
		if transaction.BorrowerID == borrowerID {
			history = append(history, transaction)
		}
	}
	return history
}

func (lib *Library) GetBookTitleStore(bookID string) string {
	book := lib.Books[bookID]
	return book.Title
}

func (lib *Library) ReturnBookStore(transactionID string) error {
	transaction, exists := lib.Transactions[transactionID]
	if !exists {
		return fmt.Errorf("Giao dich voi ID %s khong ton tai", transactionID)
	}

	if !transaction.ReturnDate.IsZero() {
		return fmt.Errorf("Sach trong giao dich %s da duoc tra truoc do", transactionID)
	}

	// cap nhat trang thai cua sach do
	book, bookExists := lib.Books[transaction.BookID]
	if !bookExists {
		return fmt.Errorf("Sach voi ID %s khong ton tai", transaction.BookID)
	}
	book.IsBorrowed = false
	lib.Books[transaction.BookID] = book

	// cap nhat ngay tra trong giao dich
	transaction.ReturnDate = time.Now()
	lib.Transactions[transactionID] = transaction

	return nil
}

func (lib *Library) SearchBooksStore(keyword string) []models.Book {
	keyword = strings.ToLower(keyword)
	results := make([]models.Book, 0)
	for _, book := range lib.Books {
		if strings.Contains(strings.ToLower(book.Title), keyword) || strings.Contains(strings.ToLower(book.Author), keyword) {
			results = append(results, book)
		}
	}
	return results
}
