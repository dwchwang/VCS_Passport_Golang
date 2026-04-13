package calculator

import (
	"errors"
)

func Divide(a, b float64) (float64, error) {
	if isZero(b) {
		return 0.0, errors.New("Division by zeros")
	}

	return a / b, nil
}

func isZero(f float64) bool {
	return f == 0.0
}
