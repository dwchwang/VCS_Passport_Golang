# 📘 Bài Tập: Quản Lý Thư Viện (QLTV)

## 📝 Mô Tả

Hệ thống quản lý thư viện đầy đủ chức năng qua giao diện dòng lệnh (CLI), bao gồm:

- **Quản lý sách:** Thêm sách mới, hiển thị danh sách sách, tìm kiếm sách theo từ khóa (tiêu đề hoặc tác giả)
- **Quản lý người mượn:** Thêm người mượn mới, hiển thị danh sách người mượn
- **Mượn / Trả sách:** Tạo giao dịch mượn sách, trả sách, xem lịch sử mượn theo người mượn
- **Trạng thái sách:** Tự động cập nhật trạng thái "Còn" / "Đã mượn" khi mượn/trả

Mỗi đối tượng (sách, người mượn, giao dịch) được sinh ID tự động bằng UUID.

---

## 📁 Cấu Trúc Thư Mục

```
exercise_QLTV/
├── main.go                        # Entry point, menu chính 9 chức năng
├── models/
│   └── models.go                  # Struct Book, Borrower, Transaction
├── library/
│   ├── library_store.go           # Tầng lưu trữ dữ liệu (data store)
│   └── library_service.go         # Tầng nghiệp vụ (business logic + UI)
└── utils/
    └── utils.go                   # Hàm tiện ích: input, validate, generate ID
```

---

## 🎓 Kiến Thức Golang Áp Dụng

### 1. Struct (Cấu trúc dữ liệu)
- `Book` — ID, Title, Author, IsBorrowed (trạng thái mượn)
- `Borrower` — ID, Name, Email
- `Transaction` — ID, BookID, BorrowerID, BorrowDate, ReturnDate
- Sử dụng struct để biểu diễn các thực thể trong hệ thống thư viện

### 2. Map (Cấu trúc dữ liệu key-value)
- `map[string]models.Book` — lưu sách theo ID
- `map[string]models.Borrower` — lưu người mượn theo ID
- `map[string]models.Transaction` — lưu giao dịch theo ID
- Truy xuất nhanh O(1) theo key, kiểm tra tồn tại: `value, exists := map[key]`
- Khởi tạo bằng `make(map[string]Type)`

### 3. Pointer Receiver
- Struct `Library` sử dụng pointer receiver `(lib *Library)` cho tất cả method
- Cho phép method thay đổi trực tiếp trạng thái bên trong struct (thêm/sửa/xóa dữ liệu trong map)
- Nếu dùng value receiver sẽ chỉ thao tác trên bản sao → dữ liệu gốc không bị thay đổi

### 4. Constructor Pattern
- Hàm `NewLibrary() *Library` — tạo và trả về instance `Library` đã khởi tạo sẵn các map
- Tránh lỗi `nil map` khi ghi dữ liệu vào map chưa được `make()`
- Đây là pattern phổ biến trong Go thay cho constructor trong OOP truyền thống

### 5. Tách tầng Service / Store (Separation of Concerns)
- **`library_store.go`** — chỉ xử lý lưu trữ: thêm, đọc, cập nhật dữ liệu trong map. Trả về `error` nếu có lỗi logic
- **`library_service.go`** — xử lý giao tiếp với người dùng: nhận input, gọi store, hiển thị kết quả
- Tách biệt giúp: dễ bảo trì, dễ test từng tầng, dễ thay đổi nguồn dữ liệu (ví dụ chuyển từ map sang database)

### 6. Error Handling (Xử lý lỗi)
- Các hàm store trả về `error`: `func (lib *Library) AddBookStore(...) error`
- Tạo lỗi tùy chỉnh: `fmt.Errorf("Sach voi ID %s da ton tai", id)`
- Kiểm tra lỗi ở service: `if err := lib.AddBookStore(...); err != nil { return err }`
- Nhiều tầng xử lý lỗi: store → service → main

### 7. Package `time`
- `time.Now()` — ghi nhận thời điểm mượn sách
- `time.Time.IsZero()` — kiểm tra sách đã trả chưa (ReturnDate chưa được set = zero value)
- `time.Time.Format("2006-01-02")` — định dạng ngày hiển thị (Go dùng layout chuẩn: tháng 01, ngày 02, năm 2006)

### 8. Package `strings` (Xử lý chuỗi)
- `strings.ToLower()` — chuyển thành chữ thường để tìm kiếm không phân biệt hoa/thường
- `strings.Contains()` — kiểm tra chuỗi có chứa từ khóa không
- Kết hợp 2 hàm trên cho chức năng tìm kiếm sách linh hoạt

### 9. Thư viện bên thứ 3 — `github.com/google/uuid`
- `uuid.New().String()` — sinh chuỗi ID duy nhất (UUID v4)
- Đảm bảo không bao giờ trùng ID giữa các đối tượng
- Import và quản lý bằng Go Modules (`go.mod`, `go.sum`)

### 10. Generics (Go 1.18+)
- Tái sử dụng hàm `IsIdUnique[T HasID]()` từ package `utils`
- Interface `HasID` làm type constraint cho generics

### 11. Zero Value trong Go
- `time.Time{}` (zero value) dùng để biểu diễn "chưa có ngày trả"
- Kiểm tra bằng `.IsZero()` — hiểu rõ zero value của từng kiểu dữ liệu là quan trọng trong Go
