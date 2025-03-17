package main

import (
	"fmt"
	"time"
)

func Channel() {
	/*
		Channels:
			- Channels are used to communicate between goroutines
			- Channels are a typed conduit through which you can send and receive values with the channel operator, <-
			- Channels are similar to pipes or event queues
			- Channels are used to synchronize execution across goroutines
			- typed, thread-safe, unbuffered by default
			- By default, sends and receives block until both the sender and receiver are ready

		Syntax:
			- chan T is a channel of type T
			- chan<- float64 is a send-only channel of type float64
			- <-chan int is a receive-only channel of type int

			For example:
				c := make(chan int) // c is a channel of type int
				c <- 1              // send 1 to c
				<-c                 // receive from c
				close(c)            // close the channel

			Note:
				- Channels are created with the make() function
				- Channels are typed by the values they convey
				- Channels can be buffered, which means that they can hold multiple values without being blocked
				- Buffered channels are created with a capacity argument
				- Buffered channels are only blocked when the buffer is full
				- Buffered channels are not blocked when they are empty

		1. Unbuffered Channels:
			- Unbuffered channels are created with the make() function
			- Unbuffered channels are created with a capacity argument of 0
			- Unbuffered channels are only blocked when both the sender and receiver are ready
			- Unbuffered channels are not blocked when they are empty
			- Unbuffered channels are not blocked when they are full:
				- If the sender is ready, the receiver is blocked

		2. Buffered Channels:
			- Buffered channels are created with a capacity argument
			- Buffered channels are created with a capacity argument of 1
			- Buffered channels are only blocked when the buffer is full
	*/

	message := make(chan string)    // create a channel of type string
	go printChannelMessage(message) // start a goroutine
	fmt.Println("Hello from Channel main function")
	fmt.Println(<-message) // receive message from channel

	/*
		Three Goroutines:
			- 1 Goroutine to send jobs to workers
			- 2 Goroutines to receive jobs from workers
			- 1 Goroutine to print the results
	*/
	const numJobs = 3                  // number of jobs to send to workers
	jobs := make(chan int, numJobs)    // create a channel of type int
	results := make(chan int, numJobs) // create a channel of type int

	// start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) // start a goroutine
	}

	// send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j // send job to channel
	}

	close(jobs) // close the channel

	// print results
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Channel Result:", <-results) //
	}
}

func printChannelMessage(message chan string) {
	time.Sleep(2 * time.Second)
	message <- "Hello from Goroutine Channel" // send message to channel
}

// worker function that receives jobs from the jobs channel and sends results to the results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Channel worker", id, "processing job", j)
		results <- j * 2 // send result to channel
		time.Sleep(time.Second)
	}
}
