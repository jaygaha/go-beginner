package main

import (
	"fmt"
	"time"
)

func printMessage() {
	fmt.Println("Hello from Goroutine")
}

func Goroutine() {
	/*
		Goroutines:
			- Goroutine is a lightweight thread managed by the Go runtime instead of OS
			- Do not communicate by sharing memory, share memory by communicating
			- Goroutine is a function that is executed concurrently with other functions
			- Goroutines run in the same address space, so access to shared memory must be synchronized

		Syntax:
			go function_name()

			Note: To create a goroutine we use the keyword go followed by a function invocation:

		Goroutine vs Parallelism:
			- Goroutines are lightweight threads managed by the Go runtime
			- Parallelism is a form of concurrency where multiple threads execute simultaneously
	*/

	go printMessage()
	fmt.Println("Hello from Main")

	// To wait for the goroutine to complete, we can use the time.Sleep() function
	// This is not recommended as we don't know how long the goroutine will take to complete
	time.Sleep(1 * time.Second)

	/*
		Output:
			Hello from Main
			Hello from Goroutine

		Note:
			- The order of execution is not guaranteed
			- The program exits before the goroutine completes
			- To wait for the goroutine to complete, we can use the time.Sleep() function
	*/
}
