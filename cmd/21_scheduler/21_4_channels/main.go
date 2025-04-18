package main

import (
	"fmt"
	"time"
)

/*
Channels
-> a channel is a communication mechanism that allows goroutines to send and receive values with the channel operator, <-
-> the channel operator is used to send and receive values from channels

The scheduler manages the goroutines, and the channel ensures main waits for both workers to complete.
The buffered channel allows workers to send messages without blocking immediately.
*/

func worker(id int, ch chan string) {
	for i := range 3 {
		fmt.Printf("Worker %d: Working %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
	ch <- fmt.Sprintf("Worker %d done", id) // <- this means that the worker is done and sends a message to the channel
}

func main() {
	ch := make(chan string, 2) // Buffered channel
	go worker(1, ch)
	go worker(2, ch)

	// Wait for both workers to finish
	for range 2 {
		fmt.Println(<-ch)
	}
	fmt.Println("Main function done")
}
