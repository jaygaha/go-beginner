# Microservices

Microservices are a software architecture where an application is split into small, independent services that:
- perform a single responsibility (e.g. authentication, payment, etc.)
- communicate over a network (e.g. HTTP, gRPC, etc.)
- can be developed, deployed, and scaled independently
- contrasts with monolithic architecture, where all functionality is in one application (codebase)

## Benefits

- Easier to maintain and scale specific parts of the application
- Teams can work on different services simultaneously
- Flexible technology stack per service

## Challenges

- Complexity of managing multiple services
- Network communication introduces latency and failure points
- Data consistency across services

## Tools

1. [RESTful APIs](./22_1_rest/README.md)
2. [Watermill](./22_2_watermill/README.md)
3. [gRPC](./22_3_grpc/README.md)
4. [go-kit](./22_4_go_kit/README.md)
5. [Micro](./22_5_go_micro/README.md)
6. [gRPC-Gateway](./22_6_gRPC-gateway/README.md)
7. [RPCX](./22_7_rpcx/README.md)
8. [Protocol Buffers](./22_8_protocol_buffers/README.md)
9. [Twirp](./22_9_twirp/README.md)

## References

- [Microservice Patterns and Resources by Chris Richardson](https://microservices.io/index.html)