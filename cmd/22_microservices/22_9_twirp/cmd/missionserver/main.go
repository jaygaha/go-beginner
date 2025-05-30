package main

import (
	"log"
	"net/http"

	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_9_twirp/internal/missionserver"
	pb "github.com/jaygaha/go-beginner/cmd/22_microservices/22_9_twirp/rpc/mission"
)

func main() {
	// Initialize the server
	server := missionserver.NewMissionServiceServer()
	twirpHandler := pb.NewMissionServiceServer(server)

	// Set up HTTP server
	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)

	log.Println("Space Mission Service running on http://localhost:8800" + twirpHandler.PathPrefix())
	if err := http.ListenAndServe(":8800", mux); err != nil {
		log.Fatal(err)
	}
}
