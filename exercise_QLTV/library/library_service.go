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
		if book.IsBorrowed {
			status = "Da muon"
		}
		fmt.Printf("ID: %s | Tieu de: %s | Tac gia: %s | Trang thai: %s\n", book.ID, book.Title, book.Author, status)
	}
	return nil
}

func AddBorrower(lib *Library) error {
	id := utils.GenerateID()
	name := utils.GetNotEmptyValue("Nhap ten nguoi muon:")
	email := utils.GetNotEmptyValue("Nhap email:")

	if err := lib.AddBorrowerStore(id, name, email); err != nil {
		return err
	}

	fmt.Println("Nguoi muon da duoc them thanh cong!")
	return nil
}

func ListBorrowers(lib *Library) error {
	borrowers := lib.ListBorrowersStore()
	if len(borrowers) == 0 {
		fmt.Println("Khong co nguoi muon nao trong thu vien.")
		return nil
	}
	fmt.Println("Danh sach nguoi muon trong thu vien:")
	for _, borrower := range borrowers {
		fmt.Printf("ID: %s | Ten: %s | Email: %s\n", borrower.ID, borrower.Name, borrower.Email)
	}
	return nil
}

func BorrowBook(lib *Library) error {
	id := utils.GenerateID()
	bookID := utils.GetNotEmptyValue("Nhap ID sach can muon:")
	borrowerID := utils.GetNotEmptyValue("Nhap ID nguoi muon:")
	if err := lib.BorrowBookStore(id, bookID, borrowerID); err != nil {
		return err
	}

	fmt.Println("Sach da duoc muon thanh cong! ID giao dich:", id)
	return nil
}

func ListBorrowHistory(lib *Library) error {
	borrowerId := utils.GetNotEmptyValue("Nhap ID nguoi muon de xem lich su muon:")
	transactions := lib.ListBorrowHistoryByBorrowerStore(borrowerId)
	if len(transactions) == 0 {
		fmt.Println("Khong co lich su muon nao.")
		return nil
	}
	fmt.Println("Lich su muon:")
	for _, transaction := range transactions {
		returnDate := "Chua tra"
		if !transaction.ReturnDate.IsZero() {
			returnDate = transaction.ReturnDate.Format("2006-01-02")
		}
		fmt.Printf("ID Transaction: %s | Ten Sach muon: %s | Ngay muon: %v | Ngay tra: %v\n", transaction.ID, lib.GetBookTitleStore(transaction.BookID), transaction.BorrowDate.Format("2006-01-02"), returnDate)
	}
	return nil
}



func ReturnBook(lib *Library) error {
	transactionID := utils.GetNotEmptyValue("Nhap ID giao dich muon:")

	if err := lib.ReturnBookStore(transactionID); err != nil {
		return err
	}
	fmt.Println("Sach da duoc tra thanh cong!")
	return nil
}

func SearchBooks(lib *Library) error {
	keyword := utils.GetNotEmptyValue("Nhap tu khoa de tim kiem sach (tieu de hoac tac gia):")
	results := lib.SearchBooksStore(keyword)
	if len(results) == 0 {
		fmt.Println("Khong tim thay sach nao voi tu khoa:", keyword)
		return nil
	}
	fmt.Println("Ket qua tim kiem:")
	for _, book := range results {
		status := "Con"	
		if book.IsBorrowed {
			status = "Da muon"
		}
		fmt.Printf("ID: %s | Tieu de: %s | Tac gia: %s | Trang thai: %s\n", book.ID, book.Title, book.Author, status)
	}
	return nil
}