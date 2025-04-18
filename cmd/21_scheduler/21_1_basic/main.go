/*
	Scheduler in Go
	-> it helps to understand more deep about how Golang works in background
	-> manages the goroutines, enabling them to run concurrently


*/

package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := range 5 {
		fmt.Printf("%s (%d)\n", message, i)
		time.Sleep(1 * time.Second)

		// simulate some work
	}
}

func main() {
	go printMessage("Goroutine 1")
	go printMessage("Goroutine 2")

	time.Sleep(1 * time.Second)

	// Allow goroutines to finish
	fmt.Println("Main function done")

	/*
		The scheduler runs both goroutines concurrently, switching between them as they sleep.
		The time.Sleep in main ensures the program doesn’t exit before the goroutines finish.
	*/
}
