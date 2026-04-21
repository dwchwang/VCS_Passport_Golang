-- Chỉ cần DROP bảng cha, Postgres sẽ tự động xóa các bảng phân mảnh (2024, 2025)
DROP TABLE IF EXISTS borrow_records CASCADE;