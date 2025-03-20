package calculator_test

import (
	"testing"
	calculator "tests"
)

// struct to hold the test cases
type testCase struct {
	input    int
	expected bool
	actual   bool
}

/*
- t *testing.T is a special type that is used to run tests
  - to manage test state and reporting failures
  - methods:
  - Errorf() to report a test failure
  - Fail() to report a test failure and stop the test
  - FailNow() to report a test failure and stop the test immediately
  - Fatal() to report a test failure and stop the test immediately
  - Fatalf() to report a test failure and stop the test immediately
  - Log() to report a test message
  - Logf() to report a test message
  - Skip() to skip the test
  - SkipNow() to skip the test immediately
  - Skipf() to skip the test
  - Helper() to mark the calling function as a helper
  - Name() to return the name of the calling function
*/
func TestCalculateIsArmstrong(t *testing.T) {
	testCase := testCase{
		input:    153,
		expected: true,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.input)

	if testCase.actual != testCase.expected {
		t.Errorf("Expected %v, got %v", testCase.expected, testCase.actual)
		t.Fail()
	}
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	testCase := testCase{
		input:    154,
		expected: false,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.input)

	if testCase.actual != testCase.expected {
		t.Errorf("Expected %v, got %v", testCase.expected, testCase.actual)
		t.Fail()
	}
}

/*
Tests within tests (subtests)
- t.Run() is a method that allows you to run a test within a test
- t.Run() takes two arguments:
  - name of the test
  - function to run

- t.Run() is useful when you want to run a test with different inputs
*/
func TestCalculateIsArmstrongInSubTests(t *testing.T) {
	t.Run("should return true for 153", func(t *testing.T) {
		testCase := testCase{
			input:    153,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.input)
		if testCase.actual != testCase.expected {
			t.Errorf("Expected %v, got %v", testCase.expected, testCase.actual)
			t.Fail()
		}
	})

	t.Run("should return true for 371", func(t *testing.T) {
		testCase := testCase{
			input:    371,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.input)
		if testCase.actual != testCase.expected {
			t.Errorf("Expected %v, got %v", testCase.expected, testCase.actual)
			t.Fail()
		}
	})
}

func TestNegativeCalculateIsArmstrongSubTests(t *testing.T) {
	t.Run("should return false for 300", func(t *testing.T) {
		testCase := testCase{
			input:    300,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.input)
		if testCase.actual != testCase.expected {
			t.Errorf("Expected %v, got %v", testCase.expected, testCase.actual)
			t.Fail()
		}
	})
}

// Benchmarking
func BenchmarkCalculateIsArmstrong(b *testing.B) {
	for b.Loop() {
		calculator.CalculateIsArmstrong(153)
	}
}
