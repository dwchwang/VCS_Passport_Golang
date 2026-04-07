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

func ListBooks() error {
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
