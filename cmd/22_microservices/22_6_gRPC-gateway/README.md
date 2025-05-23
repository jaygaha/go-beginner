# gRPC-Gateway

The `gRPC-Gateway` is a powerful tool that acts as a bridge between gRPC and REST APIs. It allows you to expose your gRPC services as RESTful JSON APIs, making them accessible to clients that don't support gRPC directly.

The `gRPC-Gateway` works using the following mechanism:

1. `protoWrite` the definition of gRPC in the file
2. Implement the gRPC server (`grpc-go` used)
3. `protocUse` the Protocol Buffers compiler to automatically generate gRPC-Gateway code
4. Using the generated code, start an HTTP server that provides the REST API.

## Why Use gRPC-Gateway?

- **Best of Both Worlds** : Use gRPC for efficient service-to-service communication while still providing REST APIs for web clients
- **Single Source of Truth** : Define your API once in Protocol Buffers and generate both gRPC and REST endpoints
- **Automatic Translation** : HTTP requests are automatically converted to gRPC calls and responses back to JSON
- **Simplified Development** : No need to maintain separate REST and gRPC implementations

## When it help?

1. Use the gRPC server when you want to access it from the web front-end or mobile app.
    . gRPC uses binary communication, so it is not directly accessible from regular browsers or mobile devices. However, the gRPC Gateway makes it available as a REST API.
2. Use this when communicating with external services that do not support gRPC.
    . Systems that can only use the REST API can also be linked via the gRPC Gateway.
3. If you're interested in building a gRPC server and providing a traditional REST API.
   . You can use the new and old APIs together to ensure a smooth transition for your clients.

## Advantages

- It is compatible with high-performance gRPC (binary) communication and the REST API.
- You can provide additional REST APIs while using the gRPC server code as is.
- Definition of API `.proto`, By unifying, you can maintain the consistency of gRPC and REST APIs.
- Technology stacks at the front and back end are flexible.

## Projec Structure

```bash
/22_6_gRPC-gateway/
├── README.md           # This documentation
├── proto/              # Protocol Buffer definitions
│   ├── greet.proto     # Service definition with HTTP annotations
│   ├── greet.pb.go     # Generated Go code for messages
│   ├── greet_grpc.pb.go # Generated gRPC service code
│   └── greet.pb.gw.go  # Generated Gateway code
├── server/             # gRPC server implementation
│   └── main.go         # Server code that implements the gRPC service
├── gateway/            # HTTP gateway implementation
│   └── main.go         # Gateway code that connects to the gRPC server
├── tools.go            # Tools dependencies for code generation
└── googleapis/         # Google API annotations for HTTP options
```

### How It Works

1. **Define Your API** : Write a Protocol Buffer file with service definitions and HTTP annotations
2. **Generate Code** : Use the Protocol Buffer compiler with gRPC and Gateway plugins to generate code
3. I**mplement the gRPC Server** : Create a server that implements the gRPC service interface
4. **Set Up the Gateway** : Create an HTTP server that connects to your gRPC server
5. **Run Both Servers** : Start both the gRPC server and the Gateway server

### How to Run
1. Install the required dependencies:
```bash
go mod tidy
```
2. Run the gRPC server:
```bash
go run server/main.go
```
3. Run the gRPC-Gateway server:
```bash
go run gateway/main.go
```
4. Send a request to the gRPC-Gateway server:
```bash
curl -X POST http://localhost:8800/v1/greet \
     -H "Content-Type: application/json" \
     -d '{"name": "Jay", "locale": "np"}'

# {"message":"नमस्ते Jay"}%  
```

### Troubleshooting

#### Missing Protoc Plugins

If you get errors about missing plugins when generating code:

```txt
protoc-gen-go: program not found or is not executable
```

Trae

Used 1 context

I need to examine the server implementation to fix the method signature mismatch for TranslateHello.

cmd/22_microservices/22_6_gRPC-gateway/server/main.go :( 1 - 100 )

Now I need to check the proto file to understand the expected method signature.

cmd/22_microservices/22_6_gRPC-gateway/proto/greet.proto :( 1 - 50 )

Let me check the generated gRPC code to understand the expected interface implementation.

cmd/22_microservices/22_6_gRPC-gateway/proto/greet_grpc.pb.go :( 1 - 100 )

# Fixing the Method Signature Error
The error is occurring because the TranslateHello method in your server implementation doesn't match the signature required by the GreetServiceServer interface.

Looking at the generated gRPC code, the interface expects:

```
TranslateHello(context.Context, *HelloRequest) 
(*HelloReply, error)
```
But your implementation has:

```
TranslateHello(*gp.HelloRequest) (*gp.
HelloReply, error)
```
You need to add the context.Context parameter to your method. Here's the fix:

```
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    gp "github.com/jaygaha/go-beginner/cmd/
    22_microservices/22_6_gRPC-gateway/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

type greetServer struct {
    gp.UnimplementedGreetServiceServer
}

// TranslateHello translates the given name to 
the given locale which implement greet_proto.
GreetServiceServer.
func (s *greetServer) TranslateHello(ctx 
context.Context, req *gp.HelloRequest) (*gp.
HelloReply, error) {
    name := req.GetName()
    locale := req.GetLocale()

    var greeting string

    switch locale {
    case "ja":
        greeting = fmt.Sprintf("こんにちは %s", 
        name)
    case "np":
        greeting = fmt.Sprintf("नमस्ते %s", 
        name)
    case "sp":
        greeting = fmt.Sprintf("Hola %s", name)
    default:
        greeting = fmt.Sprintf("Hello %s", 
        name)
    }

    return &gp.HelloReply{
        Message: greeting,
    }, nil
}

func main() {
    // gRPC server
    listener, err := net.Listen("tcp", 
    ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)

    }

    grpcServer := grpc.NewServer
    ()                            // gRPC 
    server
    reflection.Register
    (grpcServer)                           // 
    reflection is a service that provides a 
    way to list all the services that are 
    available on the server
    gp.RegisterGreetServiceServer(grpcServer, &
    greetServer{}) // register the service

    if err := grpcServer.Serve(listener); err 
    != nil {
        log.Fatalf("failed to serve: %v", err)
    }
    log.Println("gRPC server is running on 
    port 50051")
}
```
The key changes:

1. Added context to the import list
2. Modified the TranslateHello method signature to include the context.Context parameter:
   ```
   func (s *greetServer) TranslateHello(ctx 
   context.Context, req *gp.HelloRequest) (*gp.
   HelloReply, error)
   ```
This change makes your implementation match the interface required by gRPC, which will resolve the compilation error.

user9685142696123

Copy

Delete

22_6_gRPC-gateway Update a readme for the detail about gRPC-gateway implementation. The content should be beginner friendly.

Trae

18 context(s) used

Workspace

greet_grpc.pb.go 1-125

main.go 1-27

greet.pb.gw.go 1-156

main.go 14-26

greet_grpc.pb.go 91-108

... More ( 12 +)

I need to check the existing README.md file to understand what's already there before creating an updated version.

cmd/22_microservices/22_6_gRPC-gateway/README.md :( 1 - 50 )

I need to examine the project structure to understand the implementation details for the README.

cmd/22_microservices/22_6_gRPC-gateway/proto/greet.proto :( 1 - 30 )

Let me check the server implementation to understand how the gRPC service is implemented.

cmd/22_microservices/22_6_gRPC-gateway/server/main.go :( 1 - 60 )

Now I need to check the gateway implementation to understand how it connects to the gRPC server.

cmd/22_microservices/22_6_gRPC-gateway/gateway/main.go :( 1 - 30 )

Let me check the tools.go file to understand the dependencies and tools used in the project.

cmd/22_microservices/22_6_gRPC-gateway/tools.go :( 1 - 20 )

# gRPC-Gateway Tutorial
## What is gRPC-Gateway?
The gRPC-Gateway is a powerful tool that acts as a bridge between gRPC and REST APIs. It allows you to expose your gRPC services as RESTful JSON APIs, making them accessible to clients that don't support gRPC directly.

## Why Use gRPC-Gateway?
- Best of Both Worlds : Use gRPC for efficient service-to-service communication while still providing REST APIs for web clients
- Single Source of Truth : Define your API once in Protocol Buffers and generate both gRPC and REST endpoints
- Automatic Translation : HTTP requests are automatically converted to gRPC calls and responses back to JSON
- Simplified Development : No need to maintain separate REST and gRPC implementations
## Project Structure
```
/22_6_gRPC-gateway/
├── README.md           # This documentation
├── proto/              # Protocol Buffer 
definitions
│   ├── greet.proto     # Service definition 
with HTTP annotations
│   ├── greet.pb.go     # Generated Go code 
for messages
│   ├── greet_grpc.pb.go # Generated gRPC 
service code
│   └── greet.pb.gw.go  # Generated Gateway 
code
├── server/             # gRPC server 
implementation
│   └── main.go         # Server code that 
implements the gRPC service
├── gateway/            # HTTP gateway 
implementation
│   └── main.go         # Gateway code that 
connects to the gRPC server
├── tools.go            # Tools dependencies 
for code generation
└── googleapis/         # Google API 
annotations for HTTP options
```
## How It Works
1. Define Your API : Write a Protocol Buffer file with service definitions and HTTP annotations
2. Generate Code : Use the Protocol Buffer compiler with gRPC and Gateway plugins to generate code
3. Implement the gRPC Server : Create a server that implements the gRPC service interface
4. Set Up the Gateway : Create an HTTP server that connects to your gRPC server
5. Run Both Servers : Start both the gRPC server and the Gateway server
## Step-by-Step Implementation
### 1. Define Your API (proto/greet.proto)
The Protocol Buffer file defines your service, including HTTP annotations for the Gateway:

```
syntax = "proto3";

package greet;

option go_package = "github.com/jaygaha/
go-beginner/cmd/22_microservices/
22_6_gRPC-gateway/proto/greet_proto";

import "google/api/annotations.proto";

service GreetService{
    // Translates "hello" into the requested 
    locale with the given name.
    rpc TranslateHello(HelloRequest) returns 
    (HelloReply) {
        option (google.api.http) = {
            post: "/v1/greet"
            body: "*"
        };
    }
}

// The request message containing the locale 
code and a name.
message HelloRequest {
    string locale = 1; // e.g. "en", "jp", 
    "fr", etc.
    string name = 2;   // Name to include in 
    the greeting.
}

// The response message containing the 
greeting translation.
message HelloReply {
    string message = 1;
}
```
The key part is the option (google.api.http) annotation, which maps the gRPC method to an HTTP endpoint.

### 2. Generate Code
Install the required tools:

```
go install google.golang.org/protobuf/cmd/
protoc-gen-go@latest
go install google.golang.org/grpc/cmd/
protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/
grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```
Generate the code:

```
protoc -I . \
  -I ./googleapis \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=. \
  --grpc-gateway_opt=paths=source_relative \
  ./proto/greet.proto
```
This generates three files:

- greet.pb.go : Message definitions
- greet_grpc.pb.go : gRPC service definitions
- greet.pb.gw.go : Gateway code for HTTP/JSON translation
### 3. Implement the gRPC Server (server/main.go)
```
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    gp "github.com/jaygaha/go-beginner/cmd/
    22_microservices/22_6_gRPC-gateway/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

type greetServer struct {
    gp.UnimplementedGreetServiceServer
}

// TranslateHello translates the given name to 
the given locale
func (s *greetServer) TranslateHello(ctx 
context.Context, req *gp.HelloRequest) (*gp.
HelloReply, error) {
    name := req.GetName()
    locale := req.GetLocale()

    var greeting string

    switch locale {
    case "ja":
        greeting = fmt.Sprintf("こんにちは %s", 
        name)
    case "np":
        greeting = fmt.Sprintf("नमस्ते %s", 
        name)
    case "sp":
        greeting = fmt.Sprintf("Hola %s", name)
    default:
        greeting = fmt.Sprintf("Hello %s", 
        name)
    }

    return &gp.HelloReply{
        Message: greeting,
    }, nil
}

func main() {
    // gRPC server
    listener, err := net.Listen("tcp", 
    ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    reflection.Register(grpcServer) // Enables 
    server reflection
    gp.RegisterGreetServiceServer(grpcServer, &
    greetServer{})

    log.Println("gRPC server is running on 
    port 50051")
    if err := grpcServer.Serve(listener); err 
    != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```
### 4. Implement the Gateway (gateway/main.go)
```
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/grpc-ecosystem/grpc-gateway/v2/
    runtime"
    gp "github.com/jaygaha/go-beginner/cmd/
    22_microservices/22_6_gRPC-gateway/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/
    insecure"
)

func main() {
    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.
    WithTransportCredentials(insecure.
    NewCredentials())}

    // Connect gRPC-Gateway to gRPC server
    if err := gp.
    RegisterGreetServiceHandlerFromEndpoint
    (context.Background(), mux, 
    "localhost:50051", opts); err != nil {
        log.Fatalf("Failed to register 
        gRPC-Gateway: %v", err)
    }

    log.Println("gRPC-Gateway is running on 
    port 8800")
    log.Fatal(http.ListenAndServe(":8800", 
    mux))
}
```
## Running the Example
1. Start the gRPC server:
```
cd server
go run main.go
```
2. In a separate terminal, start the Gateway server:
```
cd gateway
go run main.go
```
## Testing the API
You can now access your gRPC service via HTTP:

```
curl -X POST -d '{"name":"World","locale":"ja"}
' \
  -H "Content-Type: application/json" \
  http://localhost:8800/v1/greet
```
Response:

```
{"message":"こんにちは World"}
```
Try different locales:

- "locale":"np" for Nepali
- "locale":"sp" for Spanish
- Any other value for English
## Troubleshooting
### Missing Protoc Plugins
If you get errors about missing plugins when generating code:

```
protoc-gen-go: program not found or is not 
executable
```
Make sure you've installed the required tools and they're in your PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## References

- [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/)
- [Protocol Buffers Documentation](https://protobuf.dev/)
- [gRPC Documentation](https://grpc.io/docs/)
