package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
Context in Go
	- Controlling timeout
	- cancelling goroutines
	- passing metadata across Go application

	Conntext with timeout
	- Context can be used to control the timeout of a goroutine
*/

func openConnection(done chan bool) {
	fmt.Println("Attempting to open connection...")

	if rand.Intn(10) > 5 {
		fmt.Println("Connection hanged")
		time.Sleep(100000 * time.Hour)
	} else {
		time.Sleep(3 * time.Second)
		fmt.Println("Connection opened")
	}

	done <- true
}

func openConnectionWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	defer cancel()

	// print when the context will timeout
	// deadline, _ := ctx.Deadline()
	// fmt.Println("Context will timeout at", deadline)

	// run openConnection in a goroutine
	done := make(chan bool)
	go openConnection(done)

	select {
	case <-done:
		fmt.Println("Connection opened")
	case <-ctx.Done():
		fmt.Println("Connection timed out")
	}
}

/*
Context with value
  - Context can be used to pass values across goroutines
*/
func contextWithValue() {
	type key int
	const UserKey key = 0

	ctx := context.WithValue(context.Background(), UserKey, "1234")
	userId := ctx.Value(UserKey)

	if userId, ok := userId.(string); ok {
		fmt.Println("Correct string: ", userId)
	} else {
		fmt.Println("Incorrect string")
	}
}
