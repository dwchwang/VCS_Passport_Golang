package main

import (
	"fmt"

	"dwchwang.com/exercise_qltv/library"
	"dwchwang.com/exercise_qltv/utils"
)

func main() {
	lib := library.NewLibrary()
	for {
		utils.ClearScreen()
		fmt.Println("CHUONG TRINH QUAN LY THU VIEN")
		fmt.Println("1. Them sach")
		fmt.Println("2. Hien thi danh sach sach")
		fmt.Println("3. Them nguoi muon")
		fmt.Println("4. Hien thi danh sach nguoi muon")
		fmt.Println("5. Muon sach")
		fmt.Println("6. Xem lich su muon")
		fmt.Println("7. Tra sach")
		fmt.Println("8. Tim kiem sach")
		fmt.Println("9. Thoat chuong trinh")

		choice := utils.GetPositiveInt("Nhap lua chon cua ban: ")
		switch choice {
		case 1:
			fmt.Println("===== Them sach =====")
			if err := library.AddBook(lib); err != nil {
				fmt.Println("Loi khi them sach:", err)
			}
		case 2:
			fmt.Println("===== Danh sach sach =====")
			if err := library.ListBooks(); err != nil {
				fmt.Println("Loi khi hien thi danh sach sach:", err)
			}
		case 3:
			fmt.Println("===== Them nguoi muon =====")
			if err := library.AddBorrower(); err != nil {
				fmt.Println("Loi khi them nguoi muon:", err)
			}
		case 4:
			fmt.Println("===== Danh sach nguoi muon =====")
			if err := library.ListBorrowers(); err != nil {
				fmt.Println("Loi khi hien thi danh sach nguoi muon:", err)
			}
		case 5:
			fmt.Println("===== Muon sach =====")
			if err := library.BorrowBook(); err != nil {
				fmt.Println("Loi khi muon sach:", err)
			}
		case 6:
			fmt.Println("===== Lich su muon =====")
			if err := library.ListBorrowHistory(); err != nil {
				fmt.Println("Loi khi xem lich su muon:", err)
			}
		case 7:
			fmt.Println("===== Tra sach =====")
			if err := library.ReturnBook(); err != nil {
				fmt.Println("Loi khi tra sach:", err)
			}
		case 8:
			fmt.Println("===== Tim kiem sach =====")
			if err := library.SearchBooks(); err != nil {
				fmt.Println("Loi khi tim kiem sach:", err)
			}
		case 9:
			return
		default:
			fmt.Println("Lua chon khong hop le. Vui long chon lai.")
		}
		utils.ReadInput("Nhan Enter de tiep tuc...")
	}

}
