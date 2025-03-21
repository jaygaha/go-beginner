package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		num1     float64
		num2     float64
		operator string
		expected float64
		err      bool
	}{
		{5, 3, "+", 8, false},
		{5, 3, "-", 2, false},
		{5, 3, "*", 15, false},
		{6, 3, "/", 2, false},
		{5, 0, "/", 0, true}, // Division by zero
		{5, 3, "%", 0, true}, // Invalid operator
	}

	for _, test := range tests {
		result, err := Calculate(test.num1, test.num2, test.operator)

		if (err != nil) != test.err {
			t.Errorf("Calculate(%f, %f, %q) error = %v, wantErr %v", test.num1, test.num2, test.operator, err, test.err)
			continue
		}

		if !test.err && result != test.expected {
			t.Errorf("Calculate(%f, %f, %q) = %f, want %f", test.num1, test.num2, test.operator, result, test.expected)
		}
	}
}
