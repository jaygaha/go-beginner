package main

import (
	"fmt"

	"github.com/olahol/melody"
)

/*
Melody
-> Minimalist websocket library for Go
-> based on Gorilla WebSocket
-> prvivides basic functions suchas connection handling, message broadcasting, and middleware support
*/

func main() {
	// simple usage
	m := melody.New() // creates a new melody instance
	fmt.Println("Melody server started", m)
	// output: Melody server started &{0x140000162d0 0x14000102180 0x10422a070 0x10422a080 0x10422a090 0x10422a0a0 0x10422a0b0 <nil> 0x10422a0c0 0x10422a0d0 0x10422a0e0 0x14000124050}

	fmt.Println("Check melody implementation in subfolders")
}
