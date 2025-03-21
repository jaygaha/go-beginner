package main

import (
	"fmt"
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

func main() {
	for {
		fmt.Println("User input calculator")
		fmt.Println("Enter the first number (or type 'exit' to quit):")

		var input1 string
		fmt.Scanln(&input1) // Read user input

		if input1 == "exit" {
			fmt.Println("Exiting the calculator. Goodbye!")
			return
		}

		num1, err := strconv.ParseFloat(input1, 64)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		fmt.Println("Enter the operator (+, -, *, /):")
		var operator string
		fmt.Scanln(&operator)

		fmt.Println("Enter the second number:")
		var input2 string
		fmt.Scanln(&input2)

		num2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		result, err := Calculate(num1, num2, operator)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %.2f\n", result)
	}
}
