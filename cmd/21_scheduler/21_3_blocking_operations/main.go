package main

import (
	"fmt"
	"time"
)

/*
Blocking Operations
-> a blocking operation is an operation that blocks the execution of the current goroutine until it completes
*/

func compute(id int) {
	for range 3 {
		fmt.Printf("Goroutine %d: %d\n", id, id)

		// simulate blocking i/o

		time.Sleep(1 * time.Second)
	}
}

func main() {
	go compute(1)
	go compute(2)

	time.Sleep(2 * time.Second)
	fmt.Println("Main function done")

	/*
		Undertand the output:
		-> Output will show interleaved execution of the two goroutines

		When a goroutine sleeps (simulating I/O), the scheduler parks it and runs another goroutine,
		 demonstrating efficient handling of blocking operations.
	*/
}
