CREATE TABLE borrow_records (
    id SERIAL, 
    borrower_id INT NOT NULL,
    book_id INT NOT NULL,
    borrow_date TIMESTAMP WITH TIME ZONE NOT NULL,
    return_date TIMESTAMP WITH TIME ZONE,
    status VARCHAR(20) DEFAULT 'BORROWED',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    -- Ràng buộc Khóa ngoại (Đảm bảo người mượn và sách phải tồn tại)
    CONSTRAINT fk_borrower FOREIGN KEY(borrower_id) REFERENCES borrowers(id),
    CONSTRAINT fk_book FOREIGN KEY(book_id) REFERENCES books(id)
) PARTITION BY RANGE (borrow_date);

CREATE TABLE borrow_records_2025 PARTITION OF borrow_records FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');
CREATE TABLE borrow_records_2026 PARTITION OF borrow_records FOR VALUES FROM ('2026-01-01') TO ('2027-01-01');

CREATE INDEX idx_borrow_records_borrower_book ON borrow_records(borrower_id, book_id);