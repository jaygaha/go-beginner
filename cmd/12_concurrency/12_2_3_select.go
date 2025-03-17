package main

import (
	"fmt"
	"time"
)

func SelectChannel() {
	/*
		Select:
			- Similar to switch statement, but for channels
				- Functioning of select statement is different from switch statement
			- The select statement lets a goroutine wait on multiple communication operations
			- The select statement blocks until one of its cases can run, then it executes that case
			- It chooses one at random if multiple are ready
			- If no case is ready, it blocks until one is

		Syntax:
			- select {
			- case <-ch1:
			- case x := <-ch2:
			- case ch3 <- y:
			- default:
			- }
	*/
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		defer close(ch1)

		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		defer close(ch2)

		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	var ch1Closed, ch2Closed bool

	for !ch1Closed || !ch2Closed {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				ch1Closed = true
				break
			}
			fmt.Println("received from", msg1)

		case msg2, ok := <-ch2:
			if !ok {
				ch2Closed = true
				break
			}
			fmt.Println("received from", msg2)
		}
	}

	fmt.Println("All channels are closed")
}
