package main

import (
	"errors"
	"fmt"
)

// Expected errors (Sentinal errors): It is a common practice to define a variable for each error that you expect in your program.
var errDivisionByZero = errors.New("cannot divide by zero")

// Custom errors: It is a common practice to define a variable for each error that you expect in your program.
type DivisionError struct {
	Numerator   int
	Denominator int
	Message     string
}

/*
This function is used to return a custom error message.
*/
func (e *DivisionError) Error() string {
	return fmt.Sprintf("%d / %d = %s", e.Numerator, e.Denominator, e.Message)
}

func main() {
	/*
		Error handling
		• Error handling helps deal with errors in Go programs
		• Go has a built-in error type and tools
		• ability to return multiple values from a function

		Interface
		type error interface {
			Error() string
		}
	*/

	// Constructing Errors
	fmt.Println("::Constructing errors::")
	a1, b1 := 9, 0
	result, err := divisionConstructingError(a1, b1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Expected errors
	fmt.Println("::Expected errors::")
	a2, b2 := 11, 0
	result, err = divisionExpectedErrors(a2, b2)
	if err != nil {
		switch {
		// Is method checks if the error is the same as the one we are checking for in the switch statement
		case errors.Is(err, errDivisionByZero):
			fmt.Println(err)
		default:
			fmt.Printf("unknown error: %s\n", err)
		}
	} else {
		fmt.Println()
		fmt.Printf("%d / %d = %d\n", a2, b2, result)
	}

	// Custom errors
	fmt.Println("::Custom errors::")
	a3, b3 := 12, 0
	result, err = divisionCustomErrors(a3, b3)
	if err != nil {
		// type assertion
		var divErr *DivisionError
		switch {
		// As method checks if the error is the same as the one we are checking for in the switch statement
		case errors.As(err, &divErr):
			fmt.Printf("%d / %d is not mathematically valid: %s\n",
				divErr.Numerator, divErr.Denominator, divErr.Error())
		default:
			fmt.Printf("unknown error: %s\n", err)
		}
	} else {
		fmt.Println()
		fmt.Printf("%d / %d = %d\n", a3, b3, result)
	}
}

func divisionConstructingError(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("cannot divide '%d' by zero", numerator)
	}

	return numerator / denominator, nil
}

func divisionExpectedErrors(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errDivisionByZero
	}

	return numerator / denominator, nil
}

func divisionCustomErrors(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, &DivisionError{
			Numerator:   numerator,
			Denominator: denominator,
			Message:     "cannot divide by zero",
		}
	}

	return numerator / denominator, nil
}
