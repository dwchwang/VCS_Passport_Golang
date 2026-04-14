# 📘 Bài Tập: Unit Testing trong Golang

## 📝 Mô Tả

Thực hành viết Unit Test trong Go với 2 module:

### Module 1: Calculator
- Hàm `Divide(a, b float64)` — thực hiện phép chia, xử lý trường hợp chia cho 0
- Hàm `isZero(f float64)` — kiểm tra số có bằng 0 không (hàm helper, unexported)
- Test sử dụng **Table-Driven Tests** với nhiều test case

### Module 2: Order
- Struct `Order` chứa danh sách `Item`, mỗi item có số lượng và đơn giá
- Method `ComputeTotal()` — tính tổng tiền đơn hàng
- Sử dụng thư viện `go-money` để xử lý tiền tệ chính xác (tránh lỗi floating-point)
- Test kiểm tra cả trường hợp bình thường và lỗi currency

---

## 📁 Cấu Trúc Thư Mục

```
unit_test_golang/
├── caculator/
│   ├── calculator.go              # Hàm Divide, isZero
│   └── calculator_test.go         # Table-Driven Tests cho Divide & isZero
├── order/
│   ├── order.go                   # Struct Order, Item & method ComputeTotal
│   └── order_test.go              # Test cho ComputeTotal (nominal & error case)
└── go.mod                         # Go module, dependencies: testify, go-money
```

---

## 🎓 Kiến Thức Golang Áp Dụng

### 1. Package `testing` — Framework test chuẩn của Go
- Import: `import "testing"`
- Hàm test phải có tên bắt đầu bằng `Test`: `func TestDivide(t *testing.T)`
- File test phải có hậu tố `_test.go`: `calculator_test.go`
- Chạy test: `go test ./...` hoặc `go test -v ./caculator/`

### 2. Table-Driven Tests — Pattern test chuẩn trong Go
```go
var testCases = []struct {
    name          string
    expected      float64
    expectedError error
    divisor       float64
}{
    {"division", 2.0, nil, 5.0},
    {"division by negative value", -2.0, nil, -5.0},
    {"division by zero", 0.0, errors.New("Division by zeros"), 0.0},
}
```
- Định nghĩa slice of anonymous struct chứa test cases
- Mỗi test case có: tên, input, expected output, expected error
- Duyệt bằng `for range` → dễ dàng thêm test case mới mà không thêm hàm test
- Pattern được khuyến khích chính thức bởi Go team

### 3. Sub-tests — `t.Run()`
```go
for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
        gotValue, gotError := Divide(10.0, tc.divisor)
        assert.Equal(t, tc.expectedError, gotError)
        assert.Equal(t, tc.expected, gotValue)
    })
}
```
- Mỗi test case chạy như 1 sub-test độc lập
- Output hiển thị rõ test case nào pass/fail: `TestDivide/division`, `TestDivide/division_by_zero`
- Có thể chạy riêng 1 sub-test: `go test -run TestDivide/division`

### 4. Thư viện `testify` — Assertion library
```go
import "github.com/stretchr/testify/assert"

assert.Equal(t, expected, actual)     // kiểm tra bằng nhau
assert.NoError(t, err)                // kiểm tra không có error
```
- Cung cấp assertion dễ đọc hơn so với `if got != expected { t.Errorf(...) }`
- Message lỗi tự động chi tiết: hiển thị expected vs actual
- Thư viện phổ biến nhất trong Go testing ecosystem

### 5. Test Unexported Function
```go
// Trong calculator.go: hàm unexported
func isZero(f float64) bool { return f == 0.0 }

// Trong calculator_test.go: test hàm unexported
func TestIsZero(t *testing.T) { ... }
```
- Go cho phép test hàm private (unexported) khi file test nằm **cùng package**
- `calculator_test.go` khai báo `package calculator` → truy cập được `isZero()`
- Nếu dùng `package calculator_test` (external test) thì không truy cập được

### 6. Multiple Return Values + Error
```go
func Divide(a, b float64) (float64, error) {
    if isZero(b) {
        return 0.0, errors.New("Division by zeros")
    }
    return a / b, nil
}
```
- Hàm trả về 2 giá trị: kết quả và error
- Test kiểm tra cả 2: `gotValue, gotError := Divide(10.0, tc.divisor)`
- `nil` error = thành công, non-nil error = có lỗi

### 7. Thư viện `go-money` — Xử lý tiền tệ chính xác
```go
import "github.com/Rhymond/go-money"

money.New(100, "EUR")                    // 1.00 EUR (đơn vị cents)
item.UnitPrice.Multiply(int64(quantity)) // nhân đơn giá với số lượng
amount.Add(...)                          // cộng 2 số tiền
total.Amount()                           // lấy giá trị (int64)
total.Currency().Code                    // lấy mã tiền tệ
```
- Tránh lỗi floating-point khi tính toán tiền (ví dụ: 0.1 + 0.2 ≠ 0.3)
- Lưu trữ bằng integer (cents) thay vì float
- Kiểm tra currency mismatch khi cộng số tiền khác loại

### 8. Struct Composition
```go
type Order struct {
    ID                string
    CurrencyAlphaCode string
    Items             []Item
}

type Item struct {
    ID        string
    Quantity  uint
    UnitPrice *money.Money
}
```
- `Order` chứa `[]Item` — quan hệ has-many
- `Item` chứa `*money.Money` (pointer to external type)
- Biểu diễn cấu trúc đơn hàng thực tế

### 9. Error Wrapping (Go 1.13+)
```go
return nil, fmt.Errorf("impossible to add item elements: %w", err)
```
- `%w` (wrap) bọc error gốc vào error mới
- Giữ nguyên thông tin error gốc, có thể unwrap bằng `errors.Unwrap()` hoặc `errors.Is()`
- Khác `%v` chỉ in chuỗi error, `%w` duy trì error chain

### 10. Method trên Struct
```go
func (o Order) ComputeTotal() (*money.Money, error) { ... }
```
- Value receiver `(o Order)` — không cần thay đổi Order
- Trả về `(*money.Money, error)` — pointer kết quả + error
- Duyệt `o.Items` để tính tổng
