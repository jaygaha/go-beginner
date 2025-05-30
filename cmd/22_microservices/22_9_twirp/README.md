# Twirp: A Simple RPC Framework in Go

## Introduction

Twirp is a simple RPC framework built on Protocol Buffers. It's designed to be a simpler alternative to gRPC, with a focus on ease of use. This document provides a guide to implementing a basic Twirp service based on the example in this repository.

## Key Features of Twirp

- **Protocol Buffer-based**: Uses protobuf for schema definition and serialization
- **HTTP-based**: Works over standard HTTP 1.1/2.0
- **Dual protocol support**: Supports both Protocol Buffers and JSON encoding
- **Code generation**: Generates server and client code from proto definitions
- **Error handling**: Provides standardized error types and handling

## Implementation Steps

### 1. Define Your Service in Protocol Buffers

Create a `.proto` file that defines your service interface and message types:

```protobuf
syntax = "proto3";

package mission;
option go_package = "github.com/jaygaha/go-beginner/cmd/22_microservices/22_9_twirp/rpc/mission";

service MissionService {
  rpc PlanMission(PlanMissionRequest) returns (PlanMissionResponse);
}

message PlanMissionRequest {
  string planet_name = 1; // required
  string spacecraft = 2;  // optional
}

message PlanMissionResponse {
  string mission_id = 1;
  string planet_name = 2;
  string spacecraft = 3;
  string launch_date = 4; // ISO 8601 format
  int64 travel_time_days = 5;
}
```

### 2. Generate Code from Proto Definition

Use the protoc compiler with the Twirp plugin to generate Go code:

```bash
protoc --proto_path=. \
       --twirp_out=. \
       --go_out=. \
       rpc/mission/service.proto
```

This generates:
- `service.pb.go`: Contains the Go structs for your messages
- `service.twirp.go`: Contains the Twirp server and client interfaces

### 3. Implement the Service Interface

Create a server that implements the generated interface:

```go
package missionserver

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/jaygaha/go-beginner/cmd/22_microservices/22_9_twirp/rpc/mission"
	"github.com/twitchtv/twirp"
)

// MissionServiceServer implements the MissionService interface.
type MissionServiceServer struct {
	planets  map[string]Planet
	missions map[string]Mission
	mu       sync.Mutex
}

// NewMissionServiceServer initializes the server with sample data.
func NewMissionServiceServer() *MissionServiceServer {
	return &MissionServiceServer{
		planets: map[string]Planet{
			"Mars":    {Name: "Mars", DistanceAU: 0.52},
			"Jupiter": {Name: "Jupiter", DistanceAU: 4.2},
			"Saturn":  {Name: "Saturn", DistanceAU: 8.5},
		},
		missions: make(map[string]Mission),
	}
}

// PlanMission handles the PlanMission RPC.
func (s *MissionServiceServer) PlanMission(ctx context.Context, req *pb.PlanMissionRequest) (*pb.PlanMissionResponse, error) {
	// Validate required fields
	if req.PlanetName == "" {
		return nil, twirp.RequiredArgumentError("planet_name")
	}

	// Implementation logic...
	return &pb.PlanMissionResponse{
		// Fill response fields
	}, nil
}
```

### 4. Set Up the HTTP Server

Create a main function that initializes your service and starts an HTTP server:

```go
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
```

## Error Handling in Twirp

Twirp provides standardized error types that you can use in your service implementation:

```go
// Example error handling
if req.PlanetName == "" {
    return nil, twirp.RequiredArgumentError("planet_name")
}

if !exists {
    return nil, twirp.InvalidArgumentError("planet_name", fmt.Sprintf("unknown planet: %s", req.PlanetName))
}
```

Common error types:
- `RequiredArgumentError`: When a required field is missing
- `InvalidArgumentError`: When an argument has an invalid value
- `InternalError`: For unexpected server errors
- `NotFoundError`: When a requested resource doesn't exist

## Client Usage

Twirp automatically generates client code that you can use to call your service:

```go
// Create a client
client := pb.NewMissionServiceProtobufClient("http://localhost:8800", &http.Client{})

// Make an RPC call
resp, err := client.PlanMission(context.Background(), &pb.PlanMissionRequest{
    PlanetName: "Mars",
    Spacecraft: "Perseverance",
})
```

## Benefits of Using Twirp

1. **Simplicity**: Easier to understand and implement than gRPC
2. **HTTP/1.1 Compatible**: Works with standard HTTP infrastructure
3. **JSON Support**: Allows for easy debugging and browser testing
4. **Type Safety**: Leverages Protocol Buffers for type checking
5. **Code Generation**: Reduces boilerplate and ensures consistency

## Dependencies

This implementation uses:
- `github.com/twitchtv/twirp v8.1.3+incompatible`
- `google.golang.org/protobuf v1.36.6`

## Conclusion

Twirp provides a straightforward way to implement RPC services in Go. By combining the schema definition capabilities of Protocol Buffers with the simplicity of HTTP, it offers a practical alternative to more complex RPC frameworks while still providing type safety and code generation benefits.

## References

- [Twirp](https://twitchtv.github.io/twirp/docs/intro.html)
        