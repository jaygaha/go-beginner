package main

import (
	"fmt"
	"time"
)

func UnbufferedChannel() {
	/*
		Unbuffered Channels:
			- Unbuffered channels are used to synchronize execution across goroutines
			- Unbuffered channels are used to send and receive values
			- Unbuffered channels are used to send and receive values with the channel operator, <-
			- Unbuffered channels are used to synchronize execution across goroutines
	*/
	orders := make(chan string)

	// start a goroutine to receive orders
	// customer places an order
	go func() {
		// simulating 5 orders
		for i := range 5 {
			order := fmt.Sprintf("Coffee order #%d", i+1)
			orders <- order // block until the barista (receiver) is ready to accpet order
			fmt.Println("Order placed:", order)
		}
		close(orders)
	}()

	// start a goroutine to process orders
	// barista processes the order
	for order := range orders {
		fmt.Printf("Order placed: %s\n", order)
		time.Sleep(2 * time.Second) // simulate processing time; time taken to prepare the order
		fmt.Printf("Order served: %s\n", order)
	}
}
