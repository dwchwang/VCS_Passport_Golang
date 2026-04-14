# 📘 Bài Tập: REST API Server (Server TCP to HTTP)

## 📝 Mô Tả

REST API Server quản lý User, xây dựng bằng `net/http` với routing nâng cao của Go 1.22+. Server cung cấp 3 endpoint:

| Method | Endpoint | Mô tả |
|--------|----------|-------|
| `GET` | `/users` | Lấy danh sách tất cả users |
| `POST` | `/users` | Tạo user mới (gửi JSON body) |
| `GET` | `/users/{id}` | Tìm user theo ID |

Bài tập này nâng cấp từ `package_HTTP`: sử dụng `http.NewServeMux()` thay DefaultServeMux, routing theo method, path parameter, xử lý JSON request body, và tổ chức code theo pattern API handler.

---

## 📁 Cấu Trúc Thư Mục

```
server_TCP_to_HTTP/
├── main.go                        # Khởi tạo http.Server, ServeMux, đăng ký route
├── api.go                         # Struct api, handler methods, logic validation
├── user.go                        # Struct User với JSON tags
└── go.mod                         # Go module, dependency uuid
```

---

## 🎓 Kiến Thức Golang Áp Dụng

### 1. http.NewServeMux() — Custom Router
```go
mux := http.NewServeMux()
```
- Tạo ServeMux riêng thay vì dùng `DefaultServeMux` (global)
- Ưu điểm: tránh xung đột route khi có nhiều module, kiểm soát tốt hơn, dễ test

### 2. http.Server Struct — Cấu hình server
```go
srv := &http.Server{
    Addr:    api.addr,
    Handler: mux,
}
srv.ListenAndServe()
```
- Tạo instance `http.Server` riêng thay vì gọi `http.ListenAndServe()` trực tiếp
- Có thể cấu hình thêm: `ReadTimeout`, `WriteTimeout`, `MaxHeaderBytes`, TLS...
- Pattern chuẩn trong production Go server

### 3. Method-based Routing (Go 1.22+)
```go
mux.HandleFunc("GET /users", api.getUsersHandler)
mux.HandleFunc("POST /users", api.createUsershandler)
mux.HandleFunc("GET /users/{id}", api.getUserById)
```
- Khai báo cả HTTP method và path trong 1 pattern: `"GET /users"`
- Trước Go 1.22 phải tự kiểm tra `req.Method` trong handler
- Giúp routing sạch hơn, không cần thư viện router bên ngoài (gorilla/mux, chi...)

### 4. Path Parameters
```go
id := r.PathValue("id")  // lấy giá trị từ URL pattern {id}
```
- `{id}` trong route pattern `GET /users/{id}` — dynamic path segment
- `r.PathValue("id")` — trích xuất giá trị — tính năng mới Go 1.22+
- Trước đó phải parse URL thủ công hoặc dùng thư viện bên ngoài

### 5. Struct với JSON Tags
```go
type User struct {
    ID        string `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}
```
- JSON tag `json:"first_name"` — ánh xạ field Go `FirstName` thành key JSON `first_name`
- Tuân theo convention snake_case cho JSON API
- Tag được `encoding/json` package đọc khi encode/decode

### 6. JSON Decode (Đọc Request Body)
```go
var payload User
err := json.NewDecoder(r.Body).Decode(&payload)
```
- `json.NewDecoder(r.Body)` — tạo decoder từ request body (io.Reader)
- `.Decode(&payload)` — parse JSON thành struct Go
- Dùng pointer `&payload` để decoder ghi giá trị vào biến

### 7. JSON Encode (Ghi Response)
```go
json.NewEncoder(w).Encode(users)
```
- Encode slice/struct thành JSON và ghi thẳng vào ResponseWriter

### 8. Struct Method — API Handler Pattern
```go
type api struct {
    addr string
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) { ... }
```
- Gắn handler vào struct `api` qua pointer receiver
- `api` struct có thể chứa config, database connection, logger... → share state giữa các handler
- Pattern phổ biến trong Go web application

### 9. Validation & Error Handling
```go
func insertUser(u User) error {
    if u.FirstName == "" {
        return errors.New("First name is required")
    }
    for _, user := range users {
        if user.FirstName == u.FirstName && user.LastName == u.LastName {
            return errors.New("User already exists")
        }
    }
    users = append(users, u)
    return nil
}
```
- Kiểm tra dữ liệu bắt buộc (required fields)
- Kiểm tra trùng lặp trước khi thêm
- Trả về `error` cụ thể cho từng trường hợp

### 10. HTTP Status Codes
- `http.StatusOK` (200) — thành công
- `http.StatusCreated` (201) — tạo mới thành công
- `http.StatusBadRequest` (400) — request không hợp lệ
- `http.StatusNotFound` (404) — không tìm thấy
- `http.StatusInternalServerError` (500) — lỗi server
- Sử dụng hằng số thay vì số → tự document hóa code

### 11. Thư viện `github.com/google/uuid`
- `uuid.New().String()` — sinh UUID v4 làm ID cho user mới
- Đảm bảo ID unique không cần database auto-increment

### 12. Package `errors`
- `errors.New("message")` — tạo error từ chuỗi
- Pattern chuẩn trong Go để trả lỗi từ function
