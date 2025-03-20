package calculator

import "math"

/*
	Testing in Go:
	- Test functions must be in a file with the name _test.go
	- Test functions must be named with the prefix Test

	Why test?
		- To ensure that the code is working as expected
		- To ensure that the code is working as expected when it is changed
		- To ensure that the code is working as expected when it is refactored

	Unit tests:
		- Tests which test a small part of a codebase, usually mocking out external dependencies

	Integration Tests:
		- Tests which include external resources like APIs or databases

	End to End Tests:
		- Tests which include external resources like APIs or databases

	How to run:
		- go test // run all tests in the current directory
		- go test ./... // run all tests in the package and its subpackages
		- go test -run TestCalculateIsArmstrong // run a specific test
		- go test -v // verbose output
		- go test -cover // coverage report
		- go test -coverprofile=result.out // coverage report in a file

	Benchmarking:
		- go test -bench=. // run all benchmarks in the current directory
		- go test -bench=CalculateIsArmstrong // run a specific benchmark
		- go test -bench=. -benchmem // benchmark with memory usage

*/

// Basic test function

/*
An Armstrong number is a number that is equal to the sum of its own digits each raised to the power of the number of digits.
  - takes 3 digits number `n`
  - returns true if the number is an armstrong number
    Example:
  - 153 is an armstrong number because 153 == 1^3 + 5^3 + 3^3
*/
func CalculateIsArmstrong(n int) bool {
	a := n / 100
	b := n % 100 / 10
	c := n % 10

	// return a*a*a + b*b*b + c*c*c == n
	return n == int(math.Pow(float64(a), 3))+int(math.Pow(float64(b), 3))+int(math.Pow(float64(c), 3))
}
