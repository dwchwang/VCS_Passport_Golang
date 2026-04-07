package teacher

import "fmt"

type Teacher struct {
	ID      int
	Name    string
	Subject string
	Salary  float64
	Bonus   float64
}

func (t Teacher) GetInfo() string {
	return fmt.Sprintf("ID: %d | Name: %s | Subject: %s | TotalSalary: %.2f", t.ID, t.Name, t.Subject, t.GetTotalSalary())
}

func (t Teacher) GetTotalSalary() float64 {
	return t.Salary + t.Bonus
}
