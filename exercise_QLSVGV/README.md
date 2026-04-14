# 📘 Bài Tập: Quản Lý Sinh Viên & Giảng Viên (QLSVGV)

## 📝 Mô Tả

Chương trình quản lý sinh viên và giảng viên qua giao diện dòng lệnh (CLI).  
Hỗ trợ các thao tác CRUD đầy đủ cho cả 2 đối tượng:

- **Sinh viên:** Thêm, xóa, sửa, hiển thị danh sách, tìm kiếm theo ID. Tính điểm trung bình tự động.
- **Giảng viên:** Thêm, xóa, sửa, hiển thị danh sách, tìm kiếm theo ID. Tính tổng lương (lương cơ bản + thưởng).

Chương trình có menu chính để chọn quản lý sinh viên hoặc giảng viên, mỗi phần có menu con riêng.

---

## 📁 Cấu Trúc Thư Mục

```
exercise_QLSVGV/
├── main.go                    # Entry point, hiển thị menu chính
├── student/
│   ├── student.go             # Struct Student & các method
│   └── student_service.go     # Logic CRUD sinh viên
├── teacher/
│   ├── teacher.go             # Struct Teacher & các method
│   └── teacher_service.go     # Logic CRUD giảng viên
└── utils/
    └── utils.go               # Hàm tiện ích dùng chung (input, validate, clear screen)
```

---

## 🎓 Kiến Thức Golang Áp Dụng

### 1. Struct (Cấu trúc dữ liệu)
- Định nghĩa `Student` với các trường: `ID`, `Name`, `Class`, `Scores []float64`
- Định nghĩa `Teacher` với các trường: `ID`, `Name`, `Subject`, `Salary`, `Bonus`
- Sử dụng struct để mô hình hóa đối tượng trong thực tế

### 2. Method trên Struct
- `(s Student) GetInfo()` — trả về thông tin sinh viên dạng chuỗi
- `(s Student) GetAverageScore()` — tính điểm trung bình từ slice điểm
- `(t Teacher) GetTotalSalary()` — tính tổng lương = lương cơ bản + thưởng
- `GetID()` — method chung dùng cho interface

### 3. Package (Tổ chức code)
- Tách code thành 3 package riêng biệt: `student`, `teacher`, `utils`
- Mỗi package đảm nhiệm 1 trách nhiệm cụ thể (Separation of Concerns)
- Import package nội bộ: `"dwchwang.com/exercise_qlsvgv/student"`

### 4. Exported / Unexported (Encapsulation)
- Hàm `ManageStudents()`, `ManageTeachers()` viết hoa → exported, gọi được từ `main`
- Hàm `addStudent()`, `deleteStudent()`, `updateStudent()` viết thường → unexported, chỉ dùng nội bộ trong package
- Áp dụng nguyên lý đóng gói: chỉ expose những gì cần thiết

### 5. Slice (Mảng động)
- `[]Student` và `[]Teacher` làm danh sách lưu trữ
- `append(slice, item)` — thêm phần tử
- `append(slice[:i], slice[i+1:]...)` — xóa phần tử tại vị trí `i`
- `for _, item := range slice` — duyệt slice
- `len(slice)` — kiểm tra danh sách rỗng

### 6. Interface + Generics (Go 1.18+)
- Định nghĩa interface `HasID` với method `GetID() int`
- Cả `Student` và `Teacher` đều implement `HasID`
- Hàm generic `IsIdUnique[T HasID](id int, list []T) bool` — kiểm tra ID trùng lặp, dùng chung cho cả 2 đối tượng
- Tránh lặp code nhờ type constraint

### 7. Switch-Case
- Menu chọn chức năng sử dụng `switch choice { case 1: ... case 2: ... default: ... }`
- Thay thế nhiều `if-else` giúp code sạch và dễ đọc hơn

### 8. Vòng lặp
- `for {}` — vòng lặp vô hạn cho menu chương trình (thoát bằng `return`)
- `for range` — duyệt danh sách sinh viên/giảng viên
- `for i := 0; i < n; i++` — nhập điểm theo số lượng

### 9. Xử lý Input từ người dùng
- `bufio.NewReader(os.Stdin)` — đọc chuỗi từ bàn phím
- `strings.TrimSpace()` — loại bỏ khoảng trắng và ký tự xuống dòng
- `strconv.Atoi()` — chuyển chuỗi sang số nguyên
- `strconv.ParseFloat()` — chuyển chuỗi sang số thực
- Validate vòng lặp: yêu cầu nhập lại nếu dữ liệu không hợp lệ

### 10. Cross-platform
- `runtime.GOOS` — kiểm tra hệ điều hành đang chạy
- `exec.Command("cmd", "/c", "cls")` cho Windows, `exec.Command("clear")` cho Unix
- Đảm bảo chương trình chạy được trên nhiều OS

### 11. Package `fmt`
- `fmt.Println()` — in ra console
- `fmt.Printf()` — in có format
- `fmt.Sprintf()` — tạo chuỗi có format
- `fmt.Print()` — in không xuống dòng (dùng cho prompt)
