package library

import (
	"fmt"

	"dwchwang.com/exercise_qltv/models"
)

type Library struct {
	Books map[string]models.Book
	Borrowers map[string]models.Borrower
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[string]models.Book),
		Borrowers: make(map[string]models.Borrower),
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
		ID:   id,
		Name: name,
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