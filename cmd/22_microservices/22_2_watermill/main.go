package main

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	book_service "github.com/jaygaha/go-beginner/cmd/22_microservices/22_2_watermill/book-service"
	order_service "github.com/jaygaha/go-beginner/cmd/22_microservices/22_2_watermill/order-service"
)

func main() {
	// Initialize shared pubSub for both services
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	defer pubSub.Close()

	// Set the pubSub in both services
	book_service.SetPubSub(pubSub)
	order_service.SetPubSub(pubSub)

	// start the book service
	go func() {
		log.Println("Starting book service...")
		book_service.Main()
	}()

	// start the order service
	log.Println("Starting order service...")
	order_service.Main()
}
