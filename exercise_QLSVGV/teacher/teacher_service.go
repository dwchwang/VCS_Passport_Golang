package teacher

import (
	"fmt"

	"dwchwang.com/exercise_qlsvgv/utils"
)

var teacherLists []Teacher

func addTeacher() {
	fmt.Println("===== Them giang vien =====")
	var id int
	for {
		id = utils.GetPositiveInt("Nhap ID giang vien: ")
		if utils.IsIdUnique(id, teacherLists) {
			break
		}
		fmt.Println("ID da ton tai. Vui long nhap ID khac.")
	}

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
	fmt.Println("===== Xoa giang vien =====")
	id := utils.GetPositiveInt("Nhap ID giang vien can xoa: ")
	for key, t := range teacherLists {
		if t.ID == id {
			teacherLists = append(teacherLists[:key], teacherLists[key+1:]...)
			fmt.Println("Xoa giang vien thanh cong.")
			return
		}
	}
	fmt.Println("Khong tim thay giang vien nao voi ID da nhap.")
}

func updateTeacher() {
	fmt.Println("===== Sua giang vien =====")
	id := utils.GetPositiveInt("Nhap ID giang vien can sua: ")

	for key, t := range teacherLists {
		if t.ID == id {
			fmt.Printf("Day la id: %d \n", key + 1)
			fmt.Println("Nhap thong moi (Nhan Enter de giu nguyen gia tri hien tai)")
			name := utils.GetOptionalValue(fmt.Sprintf("Nhap ten (%s):", t.Name), t.Name)
			subject := utils.GetOptionalValue(fmt.Sprintf("Nhap mon hoc (%s):", t.Subject), t.Subject)
			salary := utils.GetOptionalPositiveFloat(fmt.Sprintf("Nhap so luong co ban (%.2f)", t.Salary), t.Salary)
			bonus := utils.GetOptionalPositiveFloat(fmt.Sprintf("Nhap so thuong (%.2f)", t.Bonus), t.Bonus)

			teacherLists[key] = Teacher{
				ID:      id,
				Name:    name,
				Subject: subject,
				Salary:  salary,
				Bonus:   bonus,
			}
			fmt.Println("Sua giang vien thanh cong.")
			return
		}
	}
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
	fmt.Println("===== Tim kiem giang vien =====")
	id := utils.GetPositiveInt("Nhap ID giang vien can tim: ")
	for _, t := range teacherLists {
		if t.ID == id {
			fmt.Println("Giang vien tim thay:")
			fmt.Println(t.GetInfo())
			return
		}
	}
	fmt.Println("Khong tim thay giang vien nao voi ID da nhap.")
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
