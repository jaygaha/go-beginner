package main

import "fmt"

func process[T any](value T) {
	/*
		First value is casted to any; generic are determined only at compile time
			.(type) -> get dynamic type
	*/
	switch v := any(value).(type) {
	case int:
		fmt.Println("Integer:", v)
	case float64:
		fmt.Println("Float:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func TypeSwitches() {
	/*
		Type Switches
			- Type switches are used to compare the type of an interface value
			- The syntax is similar to a switch statement, but the cases are the types, not the values
			- The type switch is used to compare the type of an interface value
			- The syntax is similar to a switch statement, but the cases are the types, not the values

			Syntax:
			switch v := i.(type) {
			case T:
				// v is of type T
			case S:
				// v is of type S
			default:
				// v is neither T nor S
			}
	*/
	fmt.Println("Type Switches")
	a := 10
	b := 3.14
	c := "Hello"

	process(a) // Integer: 10
	process(b) // Float: 3.14
	process(c) // String: Hello

	println("Type Switches End")
}
