package main

import "fmt"

func main() {
	fmt.Println("******Concurrency******")

	fmt.Println("-----Goroutines-----")
	Goroutine()

	fmt.Println()
	fmt.Println("-----Goroutines with WaitGroup-----")
	GoroutineWaitGroup()

	fmt.Println()
	fmt.Println("-----Goroutines with Mutex-----")
	GoroutineMutex()

	fmt.Println()
	fmt.Println("-----Channels-----")
	Channel()

	fmt.Println()
	fmt.Println("-----Unbuffered Channels-----")
	UnbufferedChannel()

	fmt.Println()
	fmt.Println("-----Buffered Channels-----")
	BufferedChannel()

	fmt.Println()
	fmt.Println("-----Select-----")
	SelectChannel()

	fmt.Println()
	fmt.Println("-----Context Package-----")
	openConnectionWithTimeout()
	contextWithValue()
}
