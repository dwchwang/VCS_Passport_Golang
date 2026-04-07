package library

import (
	"fmt"
	"dwchwang.com/exercise_qltv/utils"
)



func AddBook(lib *Library) error {
	id := utils.GenerateID()
	title := utils.GetNotEmptyValue("Nhap tieu de:")
	author := utils.GetNotEmptyValue("Nhap ten tac gia:")

	if err := lib.AddBookStore(id, title, author); err != nil {
		return err
	}

	fmt.Println("Sach da duoc them thanh cong!")
	return nil
}

func ListBooks(lib *Library) error {
	books := lib.ListBooksStore()
	if len(books) == 0 {
		fmt.Println("Khong co sach nao trong thu vien.")
		return nil
	}
	fmt.Println("Danh sach sach trong thu vien:")
	for _, book := range books {
		status := "Con"
		if(book.IsBorrowed) {
			status = "Da muon"
		}
		fmt.Printf("ID: %s | Tieu de: %s | Tac gia: %s | Trang thai: %s\n", book.ID, book.Title, book.Author, status)
	}
	return nil
}

func AddBorrower() error {
	// Implementation for adding a borrower
	return nil
}

func ListBorrowers() error {
	// Implementation for listing borrowers
	return nil
}

func BorrowBook() error {
	// Implementation for borrowing a book
	return nil
}

func ListBorrowHistory() error {
	// Implementation for listing borrow history
	return nil
}

func ReturnBook() error {
	// Implementation for returning a book
	return nil
}

func SearchBooks() error {
	// Implementation for searching books
	return nil
}
