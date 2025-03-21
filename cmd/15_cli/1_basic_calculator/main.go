package main

/*
CLIs (Command Line Interfaces)
- Command line applications
- Interact with the application through the command line
- text based

Arguments
- Arguments are passed to the application when it is run
- os.Args is a slice of strings
os.Args[0] is the name of the application
os.Args[1] is the first argument
os.Args[2] is the second argument
*/

import (
	"fmt"
	"os"
	"strconv"
)

// Calculate performs the arithmetic operation based on the operator
func Calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}

		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

// Simple calculator CLI
func main() {
	// check if the number of arguments is correct
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <num1> <operator> <num2>")
		fmt.Println("Allowed operators: +, -, *, /")
		return
	}

	// parse the first number
	// strconv.ParseFloat converts a string to a float64
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid first number: ", os.Args[1])
		return
	}

	// check the operator
	operator := os.Args[2]

	// parse the second number
	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid second number: ", os.Args[3])
		return
	}

	// perform the calculation
	result, err := Calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// print the result
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
