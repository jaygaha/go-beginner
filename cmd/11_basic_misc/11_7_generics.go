package main

import (
	"fmt"
)

func Generics() {
	/*
		Generics
			- Go 1.18 introduces a new feature called generics, also known as type parameters.
			- Generics allow you to write code that can work with multiple types.

			Syntax:
				func Add[T int | float64](a, b T) T {
					return a + b
				}
	*/

	println("Generics")
	ints := []int{1, 2, 3, 4, 5}
	ints32 := []int32{6, 7, 8, 9, 10}
	floats := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	floats32 := []float32{1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Printf("Sum of ints: %v\n", SumOfNumbers(ints))         // 15
	fmt.Printf("Sum of int32: %v\n", SumOfNumbers(ints32))      // 30
	fmt.Printf("Sum of floats: %v\n", SumOfNumbers(floats))     // 15.0
	fmt.Printf("Sum of floats32: %v\n", SumOfNumbers(floats32)) // 15.0
}

type Numb interface {
	int | int32 | int64 | float32 | float64
}

func SumOfNumbers[T Numb](numbers []T) T {
	var total T
	for i := range numbers {
		total += numbers[i]
	}

	return total
}
