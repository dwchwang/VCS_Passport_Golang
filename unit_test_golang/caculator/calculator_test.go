package calculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestDivide(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.expected

			gotValue, gotError := Divide(10.0, tc.divisor)

			assert.Equal(t, tc.expectedError, gotError)
			assert.Equal(t, expected, gotValue)
		})
	}
}

var isZeroTestCases = []struct {
	name     string
	expected bool
	arg      float64
}{
	{"argument is zero", true, 0.0},
	{"argument is not zero", false, 1.0},
}

func TestIsZero(t *testing.T) {
	for _, tc := range isZeroTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isZero(tc.arg)
			assert.Equal(t, tc.expected, got)
		})
	}
}
