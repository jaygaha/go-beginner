package main

import "fmt"

func TypeInference() {
	/*
		Type Inference
			- Go allows type inference which is the process of automatically determining the type of a variable
			- Go uses the value of the variable to determine the type
			- Go allows type inference for variables, constants, and function parameters
			- Go doesn't allow type inference for function return values
	*/
	println("Type Inference")
	var i1 = 69
	f1 := 3.1415
	fmt.Printf("Variable `i1` is of type %T\n", i1) // Variable `i1` is of type int
	fmt.Printf("Variable `f1` is of type %T\n", f1) // Variable `f1` is of type float64

	var s1 = "Hello Go-phers"
	fmt.Printf("Variable `s1` is of type %T\n", s1) // Variable `s1` is of type string

	// type inference allows to declare and initialize multiple variables of different data types
	var firstName, lastName, age, salary = "John", "Doe", 30, 300000.0
	fmt.Printf("`firstName`: %T, `lastName`: %T, `age`: %T, `salary`: %T\n",
		firstName, lastName, age, salary) // `firstName`: string, `lastName`: string, `age`: int, `salary`: float64

	println("Type Inference End")
}
