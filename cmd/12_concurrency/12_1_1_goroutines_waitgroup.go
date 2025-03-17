package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RoomReservation struct {
	RoomID int
	Status string
}

func GoroutineWaitGroup() {
	/*
		WaitGroup:
			- WaitGroup is used to wait for multiple goroutines to finish
			- WaitGroup is used to wait for all goroutines to finish
			- WaitGroup is used to wait for a specific number of goroutines to finish
			- WaitGroup is used to wait for a specific time period
	*/
	// this is used to wait for all goroutines to finish
	var wg sync.WaitGroup
	// wg := sync.WaitGroup{}

	// add 3 goroutine to wait for that means we have 3 goroutines to wait for
	wg.Add(3)

	reservations := generateRoomReservation(10)

	go func() {
		defer wg.Done()
		// with go keyword, it spawns a new goroutine or run in background
		go processRoomReservation(reservations)
	}()

	go func() {
		defer wg.Done()

		go updateRoomReservationStatuses(reservations)
	}()

	go func() {
		defer wg.Done()

		go reportRoomReservationStatus(reservations)
	}()

	// wait for all goroutines to finish
	// this is blocking call until all goroutines are done
	wg.Wait()

	fmt.Println("All room reservations completed. Exiting...")
	fmt.Println()
}

func generateRoomReservation(count int) []*RoomReservation {
	// initialize reservations with status pending
	reservations := make([]*RoomReservation, count)

	for i := 0; i < count; i++ {
		reservations[i] = &RoomReservation{
			RoomID: i + 1,
			Status: "Pending",
		}
	}

	return reservations
}

func processRoomReservation(reservations []*RoomReservation) {
	for _, reservation := range reservations {
		time.Sleep(
			time.Duration(rand.Intn(1000)) * time.Millisecond,
		)

		fmt.Println("Processing room reservation for room ID: ", reservation.RoomID)
	}
}

func updateRoomReservationStatuses(reservations []*RoomReservation) {
	for _, reservation := range reservations {
		time.Sleep(
			time.Duration(rand.Intn(1000)) * time.Millisecond,
		)

		status := []string{"Pending", "Confirmed", "Cancelled"}[rand.Intn(3)]
		reservation.Status = status
		fmt.Println("Updating room reservation status for room ID: ", reservation.RoomID, " to: ", status)
	}
}

func reportRoomReservationStatus(reservations []*RoomReservation) {
	for i := 0; i < len(reservations); i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("--- Room Reservation Report ---")
		for _, reservation := range reservations {
			fmt.Println("Room ID: ", reservation.RoomID, " Status: ", reservation.Status)
		}
		fmt.Println("--- End of Report ---")
	}
}
