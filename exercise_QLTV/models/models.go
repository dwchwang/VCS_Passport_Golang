package models

import "time"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Borrower struct {
	ID    string
	Name  string
	Email string
}

type Transaction struct {
	ID         string
	BookID     string
	BorrowerID string
	BorrowDate time.Time
	ReturnDate time.Time
}