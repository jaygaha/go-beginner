package main

import "fmt"

func main() {
	year := 2025

	// Print
	fmt.Print("Hello, ジェイ") // This prints without a new line

	// Println
	fmt.Println("Hello, Golang World!") // This prints with a new line
	fmt.Println("You are developing Golang App since", year)

	// Printf
	fmt.Printf("Hello, %v \n", "Go-pher!") // This prints with a new line and a format
	fmt.Printf("You are developing Golang App since %v, %q  \n", year, "YAY!")
	fmt.Printf("year is of type of %T \n", year)
	fmt.Printf("The PI is %f \n", 3.1415)
}
