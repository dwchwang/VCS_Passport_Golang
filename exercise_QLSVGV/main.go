package main

import (
	"fmt"

	"dwchwang.com/exercise_qlsvgv/student"
	"dwchwang.com/exercise_qlsvgv/teacher"
	"dwchwang.com/exercise_qlsvgv/utils"
)

func main() {
	for {
		utils.ClearScreen()
		fmt.Println("Chuong trinh quan ly")
		fmt.Println("1. Quan ly sinh vien")
		fmt.Println("2. Quan ly giang vien")
		fmt.Println("3. Thoat chuong trinh")

		choice := utils.GetPositiveInt("Nhap lua chon cua ban: ")
		switch choice {
		case 1:
			student.ManageStudents()
		case 2:
			teacher.ManageTeachers()
		case 3:
			return
		default:
			fmt.Println("Lua chon khong hop le. Vui long chon lai.")
		}
		utils.ReadInput("Nhan Enter de tiep tuc...")
	}
}
