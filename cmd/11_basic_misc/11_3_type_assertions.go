package main

import "fmt"

type Hoge interface {
	Fuga() // Method
}

type HogeStruct struct{}

func (h *HogeStruct) Fuga() {
	fmt.Println("Fuga")
}

func (h *HogeStruct) Piyo() {
	fmt.Println("Piyo")
}

func TypeAssertion() {
	/*
		Type Assertion
			- type assertion is the process of converting the type of a variable
			- type assertion is used to convert the type of a variable to a different type
			- allows to access the concreate value stored within an interface by asserting that
				the interface value holds a specific type

			Syntax:
			v := i.(T) // i is the variable to be converted and T is the type to be converted to
			eg:
			i := 3.1415
			f := i.(float64) // converting to float
	*/
	fmt.Println("Type Assertion")
	var a any = "Go-phers"
	// interface{} or any is a type that can hold any value

	b := a.(string)
	fmt.Println(b) // Go-phers

	// type assertion with if else
	if c, ok := a.(int); ok {
		fmt.Println(c)
	} else {
		fmt.Println("a is not an integer")
	}

	// Methods
	var h Hoge = &HogeStruct{}

	// type assertion with if else
	if h, ok := h.(interface{ Piyo() }); ok {
		h.Piyo()
	} else {
		fmt.Println("h is not a HogeStruct")
	}

	println("Type Assertion End")
}
