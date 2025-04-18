package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Controlling GOMAXPROCS
-> we can control of OS threads using GOMAXPROCS
-> GOMAXPROCS is an integer that represents the maximum number of OS threads that can be used by the Go runtime
-> the Go runtime will use at most GOMAXPROCS OS threads to execute Go code
-> the default value of GOMAXPROCS is the number of CPU cores on the machine

Useful for:
- CPU-bound workloads
- I/O-bound workloads
- High-performance applications
- Parallel algorithms
- Multithreaded libraries
- Multithreaded applications
*/

func busyWork(id int) {
	for range 10 {
		fmt.Printf("Goroutine %d: %d\n", id, id)

		// simulate CPU-bound work
		for range 1000000 {
		}
	}
}

func main() {
	// Set GOMAXPROCS to 1
	runtime.GOMAXPROCS(1)
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	go busyWork(1)
	go busyWork(2)
	go busyWork(3)

	// Wait for all goroutines to finish
	time.Sleep(2 * time.Second)
	fmt.Println("Done")

	/*
		With GOMAXPROCS set to 1, the Go runtime will only use one OS thread to execute the Go code
		-> goroutines may apper to run sequentially, but they are actually running concurrently on different OS threads

		Do more:
		- Try setting GOMAXPROCS to a higher value and see how the goroutines behave
	*/
}
