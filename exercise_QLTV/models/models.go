package models
type Book struct {
	ID     string
	Title  string
	Author string
	IsBorrowed bool
}

type Borrower struct {
	ID   string
	Name string
	Email string
}