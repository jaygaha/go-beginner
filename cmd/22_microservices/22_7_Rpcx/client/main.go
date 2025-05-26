package main

import (
	"context"
	"fmt"

	"github.com/jaygaha/go-beginner/cmd/22_microservices/22_7_Rpcx/model"
	"github.com/smallnest/rpcx/client"
)

func main() {
	// Create a peer-to-peer (P2P) discovery pointing to the server's address
	d, err := client.NewPeer2PeerDiscovery("tcp@:8972", "")

	if err != nil {
		panic(err)
	}

	// Create a new RPCX client using the P2P discovery
	c := client.NewXClient("UserService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer c.Close()

	// Add a new user remotely via RPC
	addUserReq := &model.User{
		ID:     1,
		Name:   "John Doe",
		Gender: "Male",
	}
	var addUserResp model.User
	err = c.Call(context.Background(), "AddUser", addUserReq, &addUserResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added user: %+v\n", addUserResp)

	// Get a user remotely via RPC
	getUserReq := &model.User{ID: 1}
	var getUserResp model.User
	err = c.Call(context.Background(), "GetUser", getUserReq, &getUserResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got user: %+v\n", getUserResp)

}
