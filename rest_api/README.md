# REST API Library Management

## Giới thiệu

`rest_api` là một project Go xây dựng REST API quản lý thư viện mini. Project tập trung vào các nghiệp vụ cơ bản như quản lý sách, quản lý người mượn, tạo phiếu mượn sách và trả sách. Đây là project phù hợp để luyện tập cách tổ chức backend theo nhiều tầng, làm việc và kết nối với cơ sở dữ liệu PostgreSQL.

## Chức năng chính

- Quản lý danh sách sách `books`
- Quản lý danh sách người mượn `borrowers`
- Mượn sách và trả sách thông qua `borrow_records`
- Lấy danh sách sách chưa được mượn

## API hiện có

Tất cả endpoint đều dùng prefix:

```text
/api/v1
```

### Book APIs

- `POST /api/v1/books`
- `GET /api/v1/books`
- `GET /api/v1/books/unborrowed`
- `GET /api/v1/books/:id`
- `PUT /api/v1/books/:id`
- `DELETE /api/v1/books/:id`

### Borrower APIs

- `POST /api/v1/borrowers`
- `GET /api/v1/borrowers`
- `GET /api/v1/borrowers/:id`
- `PUT /api/v1/borrowers/:id`
- `DELETE /api/v1/borrowers/:id`

### Borrow Record APIs

- `POST /api/v1/borrows`
- `POST /api/v1/borrows/return/:id`

## Cấu trúc thư mục

```text
rest_api/
|-- cmd/api               # Điểm khởi động ứng dụng
|-- config                # Load env, cấu hình DB
|-- db/migrations         # SQL migration thủ công
|-- internal/app          # Container và wiring dependency
|-- internal/routes       # Đăng ký route
|-- internal/handler      # Xử lý HTTP request/response
|-- internal/service      # Business logic
|-- internal/repository   # Làm việc với database qua GORM
|-- internal/models       # Khai báo model dữ liệu
|-- go.mod
```


## Kiến thức áp dụng

Project này áp dụng nhiều kiến thức backend cơ bản trong Go:

- RESTful API với framework `Gin`
- Tổ chức kiến trúc nhiều tầng: `client -> routes -> handler -> service -> repository -> DB`
- Dependency Injection thủ công qua `internal/app/container.go`
- GORM cho CRUD, query và transaction
- PostgreSQL làm hệ quản trị cơ sở dữ liệu
- Custom SQL migration thay cho `AutoMigrate`
- Quản lý cấu hình bằng `.env`
- Kết nối database và cấu hình connection pool
- Request binding và validation cơ bản với Gin

### Index trong cơ sở dữ liệu

Project có áp dụng `index` để tối ưu truy vấn:

- Index cho `books.title`
- Index cho `books.author`
- Composite index cho `borrow_records(borrower_id, book_id)`

Mục đích của các index này là:

- Tăng tốc độ tìm kiếm sách theo tên
- Tăng tốc độ tìm kiếm sách theo tác giả
- Hỗ trợ truy vấn kết hợp người mượn và sách hiệu quả hơn

### Subquery trong truy vấn dữ liệu

Project có sử dụng `subquery` trong hàm `FindUnborrowedBooks()` ở tầng repository để lấy danh sách sách chưa được mượn.

Cách làm hiện tại:

- Tạo truy vấn con lấy danh sách `book_id` từ bảng `borrow_records`
- Dùng truy vấn chính trên bảng `books`
- Áp dụng điều kiện `WHERE id NOT IN (subquery)` để lọc ra các sách chưa xuất hiện trong danh sách mượn

Kiến thức được áp dụng ở đây:

- Viết truy vấn lồng nhau bằng GORM
- Kết hợp `Model`, `Select` và `Where` để biểu diễn subquery
- Chuyển một yêu cầu nghiệp vụ thành truy vấn SQL có cấu trúc rõ ràng

### Partition trong PostgreSQL

Project có áp dụng `partition` trên bảng `borrow_records`:

- Bảng `borrow_records` được chia partition theo `borrow_date`
- Migration hiện có partition theo năm
- Đây là kỹ thuật giúp tối ưu quản lý dữ liệu lớn và hỗ trợ truy vấn theo thời gian tốt hơn


### Transaction trong PostgreSQL

Project có sử dụng `transaction` trong module `borrow_records` để xử lý nghiệp vụ mượn sách và trả sách một cách nhất quán.

Trong nghiệp vụ mượn sách:

- Kiểm tra sách có tồn tại hay không
- Kiểm tra sách đã có trạng thái `BORROWED` chưa
- Cập nhật trạng thái sách sang `BORROWED`
- Tạo phiếu mượn mới trong bảng `borrow_records`

Trong nghiệp vụ trả sách:

- Kiểm tra phiếu mượn có tồn tại hay không
- Kiểm tra phiếu mượn đã được trả chưa
- Cập nhật trạng thái phiếu sang `RETURNED`
- Gán thời gian trả sách vào `return_date`
- Cập nhật lại trạng thái sách về `AVAILABLE`

Mục đích của transaction:

- Đảm bảo nhiều thao tác cập nhật dữ liệu được thực hiện trọn vẹn trong cùng một luồng xử lý
- Tránh trường hợp cập nhật trạng thái sách thành công nhưng tạo hoặc cập nhật phiếu mượn bị lỗi
- Nếu có lỗi ở bất kỳ bước nào thì toàn bộ thay đổi sẽ được rollback

## Kỹ thuật CSDL nổi bật

- `Index`: tối ưu tra cứu sách và truy vấn theo cặp người mượn - sách
- `Partition`: chia dữ liệu lịch sử mượn sách theo mốc thời gian
- `Foreign Key`: ràng buộc `borrow_records` với `books` và `borrowers`
- `Transaction`: đảm bảo cập nhật đồng bộ trạng thái sách và phiếu mượn/trả

## Công cụ và thư viện áp dụng

- `Go`
- `Gin`
- `GORM`
- `PostgreSQL`
- `godotenv`
- SQL migration files trong `db/migrations`
- `psql` hoặc công cụ migration thủ công để chạy các file SQL


