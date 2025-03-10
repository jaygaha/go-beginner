package main

/*
Defer, Panic, and Recover
Defer: A defer statement defers the execution of a function until the surrounding function returns.
Panic: A panic statement stops the ordinary flow of control and begins panicking.
Recover: A recover function allows a program to manage behavior after a runtime panic.

try {
    // code that may panic
} catch (err) {
    // code to handle the error
} finally {
    // code that always runs, regardless of whether an error occurred
}
*/

import "fmt"

func main() {
	// Defer
	defer fmt.Println("end") // it will execute at last
	fmt.Println("start")

	// Panic
	// basic: it will stop the execution of the program
	// panic("something went wrong")

	// Recover
	// it will recover from the panic and continue the execution of the program
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			// Re-panic to propagate the panic up
			panic(r)
		}
	}()

	// panic("something went wrong")

	// Handling using function
	raisePanic()
	// fmt.Println("end")

	// Repanicking: When you catch a panic with recover perform some operations and decide to panic again.
	// Useful for logging and cleaning up resources
	causePanic()
}

func raisePanic() {
	panic("something went wrong")
}

func causePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in causePanic:", r)
			// Re-panic to propagate the panic up
			panic(r)
		}
	}()

	// This will cause a panic
	var a []int
	fmt.Println(a[1]) // Attempting to access an out-of-bounds index
}
