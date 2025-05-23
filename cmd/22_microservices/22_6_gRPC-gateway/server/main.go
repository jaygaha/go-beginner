package main

import (
	"context"
	"fmt"
	"log"
	"net"

	gp "github.com/jaygaha/go-beginner/cmd/22_microservices/22_6_gRPC-gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type greetServer struct {
	gp.UnimplementedGreetServiceServer
}

// TranslateHello translates the given name to the given locale which implement greet_proto.GreetServiceServer.
func (s *greetServer) TranslateHello(ctx context.Context, req *gp.HelloRequest) (*gp.HelloReply, error) {
	name := req.GetName()
	locale := req.GetLocale()

	var greeting string

	switch locale {
	case "ja":
		greeting = fmt.Sprintf("こんにちは %s", name)
	case "np":
		greeting = fmt.Sprintf("नमस्ते %s", name)
	case "sp":
		greeting = fmt.Sprintf("Hola %s", name)
	default:
		greeting = fmt.Sprintf("Hello %s", name)
	}

	return &gp.HelloReply{
		Message: greeting,
	}, nil
}

func main() {
	// gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}

	grpcServer := grpc.NewServer()                            // gRPC server
	reflection.Register(grpcServer)                           // reflection is a service that provides a way to list all the services that are available on the server
	gp.RegisterGreetServiceServer(grpcServer, &greetServer{}) // register the service

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server is running on port 50051")
}
