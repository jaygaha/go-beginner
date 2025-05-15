module github.com/jaygaha/go-beginner/cmd/22_microservices/22_2_watermill

go 1.24.0

require (
	github.com/ThreeDotsLabs/watermill v1.4.6
	github.com/jaygaha/go-beginner/cmd/22_microservices/22_2_watermill/book-service v0.0.0-00010101000000-000000000000
	github.com/jaygaha/go-beginner/cmd/22_microservices/22_2_watermill/order-service v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)

replace (
	github.com/jaygaha/go-beginner/cmd/22_microservices/22_2_watermill/book-service => ./book-service
	github.com/jaygaha/go-beginner/cmd/22_microservices/22_2_watermill/order-service => ./order-service
)
