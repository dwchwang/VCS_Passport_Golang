package student

import (
	"fmt"

	"dwchwang.com/exercise_qlsvgv/utils"
)

var studentLists []Student

func addStudent() {
	var scores []float64
	fmt.Println("===== Them sinh vien =====")
	var id int
	for {
		id = utils.GetPositiveInt("Nhap ID sinh vien: ")
		if utils.IsIdUnique(id, studentLists) {
			break
		}
		fmt.Println("ID da ton tai. Vui long nhap ID khac.")
	}
	name := utils.GetNotEmptyValue("Nhap ten sinh vien: ")
	class := utils.GetNotEmptyValue("Nhap lop sinh vien: ")
	totalPoints := utils.GetPositiveInt("Nhap so luong diem sinh vien: ")
	for i := 0; i < totalPoints; i++ {
		score := utils.GetPositiveFloat(fmt.Sprintf("Nhap diem thu %d: ", i+1))
		scores = append(scores, score)
	}

	student := Student{
		ID:     id,
		Name:   name,
		Class:  class,
		Scores: scores,
	}

	studentLists = append(studentLists, student)
	fmt.Println("Them sinh vien thanh cong.")

}

func deleteStudent() {
	fmt.Println("Xoa sinh vien.")
}

func updateStudent() {
	fmt.Println("===== Sua sinh vien =====")
	id := utils.GetPositiveInt("Nhap ID sinh vien can sua: ")

	for key, s := range studentLists {
		if s.ID == id {
			fmt.Printf("Day la id: %d \n", key + 1)
			fmt.Println("Nhap thong moi (Nhan Enter de giu nguyen gia tri hien tai)")
			name := utils.GetOptionalValue(fmt.Sprintf("Nhap ten (%s):", s.Name), s.Name)
			class := utils.GetOptionalValue(fmt.Sprintf("Nhap lop (%s):", s.Class), s.Class)

			newScores := make([]float64, len(s.Scores))
			for i, score := range s.Scores {
				newScores[i] = utils.GetOptionalPositiveFloat(fmt.Sprintf("Nhap diem %d (%.2f): ", i+1, score), score)
			}
			studentLists[key] = Student{
				ID:     id,
				Name:   name,
				Class:  class,
				Scores: newScores,
			}
			fmt.Println("Sua sinh vien thanh cong.")
			return
		}
	}
}

func listStudents() {
	fmt.Println("===== Danh sach sinh vien =====")
	if len(studentLists) == 0 {
		fmt.Println("Khong co sinh vien nao.")
		return
	}
	for _, student := range studentLists {
		fmt.Println(student.GetInfo())
	}
}

func searchStudents() {
	fmt.Println("Tim kiem sinh vien.")
}

func ManageStudents() {
	for {
		utils.ClearScreen()
		fmt.Println("=====Quan ly sinh vien=====")
		fmt.Println("1. Them sinh vien")
		fmt.Println("2. Xoa sinh vien")
		fmt.Println("3. Sua sinh vien")
		fmt.Println("4. Danh sach sinh vien")
		fmt.Println("5. Tim kiem sinh vien")
		fmt.Println("6. Quay lai menu chinh")
		studentChoice := utils.GetPositiveInt("Nhap lua chon cua ban: ")
		switch studentChoice {
		case 1:
			addStudent()
		case 2:
			deleteStudent()
		case 3:
			updateStudent()
		case 4:
			listStudents()
		case 5:
			searchStudents()
		case 6:
			return
		default:
			fmt.Println("Lua chon khong hop le. Vui long chon lai.")
		}
		utils.ReadInput("Nhan Enter de tiep tuc...")
	}
}
