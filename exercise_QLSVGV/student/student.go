package student

import "fmt"

type Student struct {
	ID     int
	Name   string
	Class  string
	Scores []float64
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("ID: %d | Name: %s | Class: %s | Average Score: %.2f", s.ID, s.Name, s.Class, s.GetAverageScore())
}

func (s Student) GetAverageScore() float64 {
	if len(s.Scores) == 0 {
		return 0
	}
	total := 0.0
	for _, score := range s.Scores {
		total += score
	}
	return total / float64(len(s.Scores))
}

func (s Student) GetID() int {
	return s.ID
}
