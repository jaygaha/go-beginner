# Watermill: A Beginner's Guide to Event-Driven Microservices

## What is Watermill?

`Watermill` is a Go library that makes building event-driven applications simple and intuitive. Think of it as a messenger system for your microservices - it helps different parts of your application communicate with each other without being directly connected.

**Why use Watermill?**

- **Simple to understand**: If you know how HTTP handlers work in Go, you'll feel right at home with Watermill.
- **Flexible messaging**: Works with many message brokers (like RabbitMQ, Kafka) through a consistent API.
- **Decoupled services**: Your microservices can work independently, communicating only through events.
- **Built-in helpers**: Comes with middleware for common needs like retries, logging, and metrics.

## Understanding Pub/Sub (Publisher/Subscriber)

At its core, Watermill uses a pattern called "Pub/Sub" (short for Publisher/Subscriber), which is like a digital postal service for your application:

1. **Publishers** are services that create and send messages (like sending a letter)
2. **Subscribers** are services that receive and process those messages (like receiving mail)
3. **Topics** are like addresses that determine where messages should go

### How It Works:

```
┌────────────┐                 ┌─────────┐                 ┌────────────┐
│            │                 │         │                 │            │
│ Publisher  ├──── Message ───►│  Topic  │◄──── Listen ───┤ Subscriber │
│            │                 │         │                 │            │
└────────────┘                 └─────────┘                 └────────────┘
                                    │                            ▲
                                    │                            │
                                    └────── Message ─────────────┘
```

The basic flow is:

1. Create a `Publisher` and `Subscriber`
2. `Publisher` sends messages to a specific topic
3. `Subscriber` listens to that topic and receives the messages

**Note:** For this tutorial, we'll use in-memory pub/sub which runs everything in a single process. This is perfect for learning but not suitable for production. In real-world applications, you'd use message brokers like RabbitMQ or Kafka that can handle distributed systems.

## Tutorial Example: Book Store System

In this tutorial, we'll build a simple book store system with two microservices:

### 1. Book Service
- Manages the book inventory
- Provides a REST API endpoint: `POST /books` to add new books
- When a book is added, it publishes a `BookCreated` event

### 2. Order Service
- Manages customer orders
- Listens for `BookCreated` events to know which books are available
- Provides a REST API endpoint: `POST /orders` to create new orders

### How They Work Together

```
┌─────────────────┐                                  ┌─────────────────┐
│                 │                                  │                 │
│   Book Service  │                                  │  Order Service  │
│   (Port 8801)   │                                  │   (Port 8802)   │
│                 │                                  │                 │
└────────┬────────┘                                  └────────┬────────┘
         │                                                    │
         │                                                    │
         │ 1. Add Book                                        │ 4. Create Order
         │    POST /books                                     │    POST /orders
         │                                                    │
         ▼                                                    ▼
┌─────────────────┐                                  ┌─────────────────┐
│                 │                                  │                 │
│   Book Store    │                                  │   Order Store   │
│                 │                                  │                 │
└────────┬────────┘                                  └────────┬────────┘
         │                                                    ▲
         │                                                    │
         │ 2. Publish Event                                   │
         │    "books.created"                                 │
         ▼                                                    │
┌─────────────────┐                                           │
│                 │       3. Subscribe & Process              │
│    Watermill    ├───────────────────────────────────────────┘
│    Pub/Sub      │          "books.created"
│                 │
└─────────────────┘
```

This approach has several advantages over direct HTTP calls between services:

1. **Loose coupling**: Services don't need to know about each other directly
2. **Better resilience**: If one service is down, messages can be processed later
3. **Scalability**: Each service can scale independently

## Key Concepts in the Code

- **Message**: A data structure with a unique ID and payload (the actual content)
- **Publisher**: Creates and sends messages to a topic
- **Subscriber**: Receives messages from a topic
- **Handler**: A function that processes received messages
- **Middleware**: Functions that can modify or enhance message processing

By using Watermill, you can focus on your business logic rather than the complexities of service communication.

## References

- [Watermill Website](https://watermill.io/)
- [Watermill Documentation](https://watermill.io/docs/getting-started)
