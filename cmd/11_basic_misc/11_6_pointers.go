package main

import "fmt"

func Pointers() {
	/*
			Pointers
		 		- A pointer is a variable that stores the memory address of another variable.
				- allows to pass reference to a value instead of copying the value

				Syntax:
					var ptr *int
					ptr = &x
					ptr = new(int)

				The type *T is a pointer to a T value. Its zero value is nil.
				The & operator generates a pointer to its operand.
				The * operator denotes the pointer's underlying value.
				Unlike C, Go has no pointer arithmetic.
	*/
	println("Pointers")
	x, y := 1, 2

	p := &x         // p points to x
	fmt.Println(*p) // read x through the pointer p (1)
	*p = 21         // set x through the pointer p
	fmt.Println(x)  // see the new value of x (21)

	p = &y         // p points to y
	*p = *p / 37   // divide y through the pointer p
	fmt.Println(y) // see the new value of y (0)

	// Zero Value
	println("\nZero Value")
	i := 1
	fmt.Println("Initial value: ", i)

	zeroVal(i)
	fmt.Println("After zeroVal: ", i)

	zeroPtr(&i)
	fmt.Println("After zeroPtr: ", i)

	fmt.Println("Pointer: ", &i)
}

// arguments are passed by value, so the original value is not changed
func zeroVal(iVal int) {
	iVal = 0
}

// arguments are passed by reference, so the original value is changed
func zeroPtr(iPtr *int) {
	*iPtr = 0
}
