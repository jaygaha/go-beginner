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

## References

- [Microservice Patterns and Resources by Chris Richardson](https://microservices.io/index.html)