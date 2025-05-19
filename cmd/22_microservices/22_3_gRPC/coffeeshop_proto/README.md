# gRPC Coffee Shop Service

This directory contains the Protocol Buffers definition for a Coffee Shop service using gRPC.

## Directory Structure

- `coffee_shop.proto`: The Protocol Buffers definition file
- `coffeeshop_proto/`: Directory for generated Go code

## Generating Go Code

To generate the Go code from the Protocol Buffers definition, run the following command from the `22_3_gRPC` directory:

```bash
# Create the output directory if it doesn't exist
mkdir -p coffeeshop_proto

# Generate Go code from the Protocol Buffers definition
protoc --go_out=coffeeshop_proto --go_opt=paths=source_relative \
       --go-grpc_out=coffeeshop_proto --go-grpc_opt=paths=source_relative \
       coffee_shop.proto
```

This will generate the Go code in the `coffeeshop_proto` directory instead of the root directory.

## Running the Service

Implement the server and client code using the generated Go code.