package teacher

import (
	"fmt"

	"dwchwang.com/exercise_qlsvgv/utils"
)

var teacherLists []Teacher

func addTeacher() {
	fmt.Println("===== Them giang vien =====")
	id := utils.GetPositiveInt("Nhap ID giang vien: ")
	name := utils.GetNotEmptyValue("Nhap ten giang vien: ")
	subject := utils.GetNotEmptyValue("Nhap mon hoc: ")
	salary := utils.GetPositiveFloat("Nhap luong co ban: ")
	bonus := utils.GetPositiveFloat("Nhap thuong: ")

	teacher := Teacher{
		ID:      id,
		Name:    name,
		Subject: subject,
		Salary:  salary,
		Bonus:   bonus,
	}

	teacherLists = append(teacherLists, teacher)
	fmt.Println("Them giang vien thanh cong.")

}

func deleteTeacher() {
	fmt.Println("Xoa giang vien.")
}

func updateTeacher() {
	fmt.Println("Sua giang vien.")
}

func listTeachers() {
	fmt.Println("=====Danh sach giang vien=====")
	if len(teacherLists) == 0 {
		fmt.Println("Khong co giang vien nao.")
		return
	}
	for _, teacher := range teacherLists {
		fmt.Println(teacher.GetInfo())
	}
}

func searchTeachers() {
	fmt.Println("Tim kiem giang vien.")
}

func ManageTeachers() {
	for {
		utils.ClearScreen()
		fmt.Println("=====Quan ly giang vien=====")
		fmt.Println("1. Them giang vien")
		fmt.Println("2. Xoa giang vien")
		fmt.Println("3. Sua giang vien")
		fmt.Println("4. Danh sach giang vien")
		fmt.Println("5. Tim kiem giang vien")
		fmt.Println("6. Quay lai menu chinh")
		teacherChoice := utils.GetPositiveInt("Nhap lua chon cua ban: ")
		switch teacherChoice {
		case 1:
			addTeacher()
		case 2:
			deleteTeacher()
		case 3:
			updateTeacher()
		case 4:
			listTeachers()
		case 5:
			searchTeachers()
		case 6:
			return
		default:
			fmt.Println("Lua chon khong hop le. Vui long chon lai.")
		}
		utils.ReadInput("Nhan Enter de tiep tuc...")
	}
}
