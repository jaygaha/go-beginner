package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gp "github.com/jaygaha/go-beginner/cmd/22_microservices/22_6_gRPC-gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mux := runtime.NewServeMux() // create a new ServeMux instance; it's a multiplexer that routes incoming requests to the appropriate handlers
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Connect gRPC-Gateway to gRPC server
	if err := gp.RegisterGreetServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts); err != nil {
		log.Fatalf("Failed to register gRPC-Gateway: %v", err)
	}

	log.Println("gRPC-Gateway is running on port 8800")
	log.Fatal(http.ListenAndServe(":8800", mux))
}
