package main

import (
	"fmt"
	"time"
)

func BufferedChannel() {
	/*
		Buffered Channels:
			- Channels can be buffered
			- Buffered channels accept a limited number of values without a corresponding receiver for those values
			- Buffered channels are useful when the sender and receiver have asynchronous properties
			- Buffered channels are created by passing the desired buffer length as the second argument to make
			- Buffered channels block only when the buffer is full

		Example:
			- orders := make(chan string, 3) // buffered channel with a capacity of 3

		Syntax:
			- chan T is a channel of type T
			- chan<- float64 is a send-only channel of type float64
			- <-chan int is a receive-only channel of type int
			- chan<- int is a send-only channel of type int
			- <-chan int is a receive-only channel of type int

		Why Buffered Channels?
			- Buffered channels can be used to decouple the sender and receiver
			- Buffered channels can be used to limit the number of goroutines that are executing at the same time
			- Buffered channels can be used to limit the number of requests that are queued up
	*/
	orders := make(chan string, 3) // buffered channel with a capacity of 3

	// start a goroutine to receive orders
	// customer places an order
	go func() {
		// simulating 5 orders
		for i := range 5 {
			order := fmt.Sprintf("Bufferred Coffee order #%d", i+1)
			orders <- order // block until the barista (receiver) is ready to accpet order
			fmt.Println("Bufferred order placed:", order)
		}
		close(orders)
	}()

	// start a goroutine to process orders
	// barista processes the order
	for order := range orders {
		fmt.Printf("Bufferred order placed: %s\n", order)
		time.Sleep(2 * time.Second) // simulate processing time; time taken to prepare the order
		fmt.Printf("Bufferred order served: %s\n", order)
	}
}
