# 📘 Bài Tập: HTTP Server Cơ Bản (Package HTTP)

## 📝 Mô Tả

HTTP Server đơn giản sử dụng package chuẩn `net/http` của Go. Server lắng nghe tại port 8080 và xử lý request đến endpoint `/demo`:

- Chỉ chấp nhận method **GET**
- Trả về response dạng **JSON** chứa thông tin chào mừng
- Trả về lỗi `405 Method Not Allowed` nếu dùng method khác GET

Đây là bài tập nền tảng để hiểu cách Go xử lý HTTP request/response trước khi chuyển sang xây dựng REST API phức tạp hơn.

---

## 📁 Cấu Trúc Thư Mục

```
package_HTTP/
├── main.go                        # Khởi tạo server, đăng ký handler, xử lý request
└── go.mod                         # Go module
```

---

## 🎓 Kiến Thức Golang Áp Dụng

### 1. Package `net/http` — Xây dựng HTTP Server
- `http.HandleFunc("/demo", demoHandler)` — đăng ký handler function cho route `/demo` trên DefaultServeMux
- `http.ListenAndServe(":8080", nil)` — khởi động server lắng nghe tại port 8080
- Tham số `nil` nghĩa là sử dụng `DefaultServeMux` làm router mặc định

### 2. HTTP Handler Function
- Signature chuẩn: `func demoHandler(res http.ResponseWriter, req *http.Request)`
- `http.ResponseWriter` — interface dùng để ghi HTTP response (header, body, status code)
- `*http.Request` — struct chứa toàn bộ thông tin request (method, URL, header, body)

### 3. Kiểm tra HTTP Method
```go
if req.Method != http.MethodGet {
    http.Error(res, "...", http.StatusMethodNotAllowed)
    return
}
```
- Sử dụng hằng số `http.MethodGet` thay vì chuỗi `"GET"` — tránh lỗi typo
- `return` ngay sau khi gửi error — pattern "early return" giúp code dễ đọc

### 4. JSON Response
- **Cách 1 (Commented):** `json.Marshal()` → `res.Write(data)` — encode thành `[]byte` rồi ghi
- **Cách 2 (Đang dùng):** `json.NewEncoder(res).Encode(response)` — encode và ghi trực tiếp vào `ResponseWriter`
- Cách 2 gọn hơn, không cần biến trung gian, phù hợp cho HTTP response

### 5. HTTP Header
- `res.Header().Set("Content-Type", "application/json")` — báo cho client biết response là JSON
- Phải set header **trước** khi ghi body

### 6. HTTP Error Response
- `http.Error(res, message, statusCode)` — hàm tiện ích gửi response lỗi
- `http.StatusMethodNotAllowed` (405) — method không được phép
- `http.StatusInternalServerError` (500) — lỗi server nội bộ

### 7. Map Literal
```go
response := map[string]string{
    "message": "Chao mung ban",
    "author":  "dwchwang",
}
```
- Khai báo và khởi tạo map cùng lúc
- Dùng làm dữ liệu response nhanh mà không cần định nghĩa struct

### 8. Package `log`
- `log.Println("Khoi dong server...")` — in log kèm timestamp
- `log.Fatal("Loi...")` — in log rồi gọi `os.Exit(1)` — dừng chương trình ngay nếu server không start được
- `log.Printf("%+v", res)` — in chi tiết struct với tên field (format `%+v`)

### 9. Package `encoding/json`
- `json.NewEncoder(writer)` — tạo encoder ghi vào bất kỳ `io.Writer` nào
- `.Encode(v)` — encode giá trị Go thành JSON và ghi ra writer
- Hỗ trợ cả struct, map, slice → tự động chuyển thành JSON
