package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_7_Rpcx/model"
	"github.com/smallnest/rpcx/server"
)

// UserService defines methods that can be remotely called by clients
type UserService struct{}

// In-memory storage for users
var (
	users = make(map[int]*model.User)
	mu    sync.Mutex
)

// AddUser adds a new user to the store
func (s *UserService) AddUser(ctx context.Context, req *model.User, res *model.User) error {
	mu.Lock()
	defer mu.Unlock()

	users[req.ID] = req
	*res = *req // copy request user to response

	return nil
}

// GetUser retrieves a user from the store
func (s *UserService) GetUser(ctx context.Context, req *model.User, res *model.User) error {
	mu.Lock()
	defer mu.Unlock()

	if user, ok := users[req.ID]; ok {
		*res = *user

		return nil
	}

	return fmt.Errorf("user not found")
}

func main() {
	// Create a new RPCX server
	server := server.NewServer()

	// Register the UserService to the server
	if err := server.RegisterName("UserService", new(UserService), ""); err != nil {
		panic(err)
	}

	fmt.Println("RPCX server is running on port 8972...")
	// Start the server
	if err := server.Serve("tcp", ":8972"); err != nil {
		panic(err)
	}
}
