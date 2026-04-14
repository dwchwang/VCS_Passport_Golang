# 📘 Bài Tập: Quản Lý Hệ Thống — System Monitor (QLHT)

## 📝 Mô Tả

Ứng dụng giám sát hệ thống theo thời gian thực, mô phỏng công cụ monitoring chuyên nghiệp. Các chức năng chính:

- **Giám sát CPU:** Theo dõi phần trăm sử dụng CPU, cảnh báo khi vượt 60%
- **Giám sát RAM:** Theo dõi phần trăm bộ nhớ đã dùng, cảnh báo khi vượt 60%
- **Giám sát Disk:** Theo dõi phần trăm dung lượng ổ đĩa đã dùng
- **Giám sát Network:** Hiển thị lượng dữ liệu gửi/nhận (KB)
- **Top Processes:** Liệt kê top 5 process tiêu tốn CPU và RAM nhiều nhất
- **Xuất CSV:** Ghi dữ liệu process ra file `process_stats.csv`
- **Log cảnh báo:** Ghi cảnh báo vào file `alert.log` khi metric vượt ngưỡng

Chương trình chạy các monitor song song bằng goroutine, thu thập dữ liệu qua channel, và tự động dừng sau 10 giây.

---

## 📁 Cấu Trúc Thư Mục

```
exercise_QLHT/
├── main.go                        # Entry point, khởi tạo context, goroutine, channel
├── models/
│   └── models.go                  # Interface Monitor, struct SystemStat, ProcStat, Mutex
├── monitor/
│   ├── cpu.go                     # CpuMonitor — implement Monitor interface
│   ├── mem.go                     # MemMonitor — implement Monitor interface
│   ├── disk.go                    # DiskMonitor — implement Monitor interface
│   └── net.go                     # NetMonitor — implement Monitor interface
├── processor/
│   └── processor.go               # RunMonitor, GetTopProcessor, ExportToCSV, LogAlert
├── process_stats.csv              # File CSV xuất kết quả
└── alert.log                      # File log cảnh báo
```

---

## 🎓 Kiến Thức Golang Áp Dụng

### 1. Interface (Giao diện)
- Định nghĩa interface `Monitor` với 2 method:
  - `Name() string` — trả về tên monitor
  - `Check(ctx context.Context) (string, bool)` — kiểm tra metric, trả về giá trị và cờ cảnh báo
- 4 struct `CpuMonitor`, `MemMonitor`, `DiskMonitor`, `NetMonitor` đều implement interface này
- Tuân theo nguyên lý Open/Closed: thêm monitor mới chỉ cần tạo struct mới implement `Monitor`, không sửa code cũ

### 2. Polymorphism (Đa hình)
- Slice `[]models.Monitor` chứa nhiều loại monitor khác nhau
- Gọi `m.Check(ctx)` thống nhất trên mỗi phần tử — Go tự động dispatch đến method đúng
- Hàm `RunMonitor()` nhận tham số kiểu `models.Monitor` — hoạt động với bất kỳ monitor nào

### 3. Goroutine (Lập trình đồng thời)
- Mỗi monitor chạy trên 1 goroutine riêng: `go processor.RunMonitor(ctx, &wg, statCh, m)`
- Goroutine riêng để thu thập stat từ channel
- Goroutine riêng để in kết quả định kỳ
- Goroutine riêng để đóng channel sau khi WaitGroup hoàn thành
- Trong `GetTopProcessor()`: mỗi process cũng được xử lý trong goroutine riêng → xử lý hàng trăm process song song

### 4. Channel (Kênh giao tiếp)
- `statCh := make(chan models.SystemStat)` — unbuffered channel giao tiếp giữa monitor goroutine và main goroutine
- `procChan := make(chan models.ProcStat, len(processes))` — buffered channel tránh blocking khi xử lý nhiều process
- Send-only channel: `statCh chan<- models.SystemStat` — giới hạn goroutine chỉ được gửi, không được nhận
- `for stat := range statCh` — nhận dữ liệu liên tục từ channel cho đến khi channel bị đóng

### 5. sync.WaitGroup
- `wg.Add(1)` trước khi khởi chạy mỗi goroutine
- `defer wg.Done()` trong goroutine để báo hoàn thành
- `wg.Wait()` chờ tất cả goroutine kết thúc trước khi close channel và thoát

### 6. sync.Mutex (Khóa loại trừ)
- `StatMutex sync.Mutex` — bảo vệ truy cập đồng thời vào `map[string]SystemStat`
- `StatMutex.Lock()` / `StatMutex.Unlock()` trước/sau khi đọc/ghi map
- `defer models.StatMutex.Unlock()` — đảm bảo unlock ngay cả khi có panic
- Tránh data race khi nhiều goroutine đọc/ghi cùng 1 map

### 7. Context (Quản lý vòng đời goroutine)
- `context.WithCancel(context.Background())` — tạo context có thể hủy
- `cancel()` — gửi tín hiệu dừng đến tất cả goroutine đang lắng nghe
- Các goroutine kiểm tra `ctx.Done()` để biết khi nào cần dừng
- Pattern chuẩn trong Go để quản lý vòng đời và timeout goroutine

### 8. Select Statement
```go
select {
case <-ctx.Done():    // nhận tín hiệu hủy → dừng goroutine
    return
case <-ticker.C:      // nhận tín hiệu tick → thực hiện kiểm tra
    value, alert := m.Check(ctx)
}
```
- Lắng nghe nhiều channel cùng lúc
- Channel nào có dữ liệu trước thì xử lý trước
- Kết hợp `ctx.Done()` để graceful shutdown

### 9. time.Ticker (Bộ đếm thời gian lặp lại)
- `time.NewTicker(1 * time.Second)` — kiểm tra metric mỗi 1 giây
- `time.NewTicker(4 * time.Second)` — in kết quả tổng hợp mỗi 4 giây
- `defer ticker.Stop()` — dọn dẹp ticker khi không cần nữa
- `time.Sleep(10 * time.Second)` — chạy chương trình trong 10 giây

### 10. Thư viện bên thứ 3 — `github.com/shirou/gopsutil/v4`
- `cpu.PercentWithContext()` — đo phần trăm CPU
- `mem.VirtualMemoryWithContext()` — đo RAM
- `disk.UsageWithContext()` — đo dung lượng ổ đĩa
- `net.IOCountersWithContext()` — đo lưu lượng mạng
- `process.ProcessesWithContext()` — liệt kê tất cả process
- Tất cả hàm đều nhận `context.Context` → hỗ trợ cancel

### 11. sort.Slice (Sắp xếp)
- `sort.Slice(cpuList, func(i, j int) bool { return cpuList[i].CPUPercent > cpuList[j].CPUPercent })`
- Sắp xếp giảm dần theo CPU% và Memory%
- Dùng closure làm hàm so sánh tùy chỉnh

### 12. File I/O (Đọc ghi file)
- `os.OpenFile("file.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)` — mở file để ghi thêm
- `f.Stat()` — kiểm tra thông tin file (size = 0 → ghi header CSV)
- `f.WriteString()` — ghi chuỗi vào file
- `defer f.Close()` — đảm bảo đóng file khi hàm kết thúc
- Flag: `O_APPEND` (ghi tiếp), `O_CREATE` (tạo nếu chưa có), `O_WRONLY` (chỉ ghi)

### 13. Closure trong Goroutine
```go
go func(p *process.Process) {
    defer wg.Done()
    // sử dụng p an toàn
}(p)
```
- Truyền biến `p` qua tham số closure thay vì capture trực tiếp
- Tránh race condition khi biến loop thay đổi giá trị trước khi goroutine chạy

### 14. Cross-platform
- `runtime.GOOS == "windows"` → dùng path `C:\` cho disk monitor
- Default → dùng path `/` cho Linux/macOS
