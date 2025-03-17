package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RoomReservationMutex struct {
	RoomID int
	Status string
	mu     sync.Mutex
}

// global variable
var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func GoroutineMutex() {
	/*
		Mutex:
			- Mutex is used to protect shared resources
			- Mutex is used to lock and unlock resources ensure that only one goroutine can access it at a time
			- Mutex is used to prevent race conditions
			- Mutex is like a traffic light which allows only one car to pass at a time
	*/
	var wg sync.WaitGroup

	wg.Add(3)

	reservations := generateRoomReservationMutex(4)

	for range 3 {
		go func() {
			defer wg.Done()

			for _, reservation := range reservations {
				updateRoomReservationStatusMutex(reservation)
			}
		}()
	}

	go reportRoomReservationStatusMutex(reservations)

	// wait for all goroutines to finish
	// this is blocking call until all goroutines are done
	wg.Wait()

	fmt.Println("All mutex room reservations completed. Exiting...")
	fmt.Println()
}

func generateRoomReservationMutex(count int) []*RoomReservationMutex {
	// initialize reservations with status pending
	reservations := make([]*RoomReservationMutex, count)

	for i := range count {
		reservations[i] = &RoomReservationMutex{
			RoomID: i + 1,
			Status: "Pending",
		}
	}

	return reservations
}

func updateRoomReservationStatusMutex(reservation *RoomReservationMutex) {
	reservation.mu.Lock() // lock the resource
	time.Sleep(
		time.Duration(rand.Intn(1000)) * time.Millisecond,
	)

	status := []string{"Pending", "Confirmed", "Cancelled"}[rand.Intn(3)]
	reservation.Status = status
	fmt.Println("Updating room reservation status for room ID: ", reservation.RoomID, " to: ", status)

	reservation.mu.Unlock() // unlock the resource

	updateMutex.Lock()
	defer updateMutex.Unlock()

	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1
}

func reportRoomReservationStatusMutex(reservations []*RoomReservationMutex) {
	fmt.Println("--- Mutex Room Reservation Report ---")
	for _, reservation := range reservations {
		fmt.Println("Room ID: ", reservation.RoomID, " Status: ", reservation.Status)
	}
	fmt.Println("--- Mutex End of Report ---")
}
