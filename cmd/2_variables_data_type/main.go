package main

import "fmt"

const globalConst int = 99

var globalVar int = 99

func main() {
	// 1. Constants
	fmt.Println("Constants")
	fmt.Println("Global Constant: ", globalConst)

	const country string = "Japan"
	fmt.Println("Country: ", country)

	const age, siblings int = 25, 1
	fmt.Println("Age: ", age, "Siblings: ", siblings)

	const PI float32 = 3.14
	fmt.Println("PI: ", PI)

	const isMarried bool = false
	fmt.Println("Is Married: ", isMarried)

	// Numerical Constants
	const (
		num1 int = 10
		num2 int = 20
	)
	fmt.Println("Numerical Constants")
	fmt.Println("Sum: ", num1+num2)

	// 2. Variables
	fmt.Println()
	fmt.Println("Variables")
	fmt.Println("Global Variable: ", globalVar)

	var name string = "John"
	fmt.Println("Name: ", name)

	var (
		firstName string = "John"
		lastName  string = "Doe"
	)
	fmt.Println("First Name: ", firstName, "Last Name: ", lastName)

	var grade int = 75
	fmt.Println("Grade: ", grade)

	var avgGrade float32 = 65.5
	fmt.Println("Average Grade: ", avgGrade)

	var isPassed bool = true
	fmt.Println("Is Passed: ", isPassed)

	fruit := "Apple" // short hand declaration
	fmt.Println("Fruit: ", fruit)

	// 3. Data Types
	fmt.Println()
	fmt.Println("Basic Data Types")

	// Bool
	var isBool bool = true
	fmt.Println("Bool: ", isBool)
	// String
	var str_1 string = "Hello, World!"
	fmt.Println("String: ", str_1)
	// Integer
	var int_1 int = 10
	var int_8 int8 = -10
	var int_16 int16 = -10
	var int_32 int32 = 10
	var int_64 int64 = 10
	fmt.Println("Integer: ", int_1, int_8, int_16, int_32, int_64)
	// Unsigned Integer
	var uint_1 uint = 10
	var uint_8 uint8 = 10
	var uint_16 uint16 = 10
	var uint_32 uint32 = 10
	var uint_64 uint64 = 10
	fmt.Println("Unsigned Integer: ", uint_1, uint_8, uint_16, uint_32, uint_64)
	// Float
	var float_32 float32 = 10.5
	var float_64 float64 = 10.5
	fmt.Println("Float: ", float_32, float_64)
	// Complex
	var complex_64 complex64 = 10 + 5i
	var complex_128 complex128 = 10 + 5i
	fmt.Println("Complex: ", complex_64, complex_128)

	// Type conversion
	fmt.Println()
	fmt.Println("Type Conversion")
	var conversion int = 75
	fmt.Println("Conversion: ", conversion)
	var conversionFloat float32 = float32(conversion)
	fmt.Println("float32 Conversion: ", conversionFloat)
	var conversionUint uint = uint(conversion)
	fmt.Println("uint Conversion: ", conversionUint)

	// Type inference
	fmt.Println()
	inference := 75.99 // change this value to check
	fmt.Printf("inference is of type %T\n", inference)
}
