package main

import (
	"context"
	"log"
	"net"

	cp "github.com/jaygaha/go-beginner/cmd/22_microservices/22_3_gRPC/coffeeshop_proto"
	"google.golang.org/grpc"
)

type server struct {
	cp.UnimplementedCoffeeShopServer // UnimplementedCoffeeShopServer is a struct that implements the CoffeeShopServer interface
}

// GetMenu is a method that returns a stream of Menu
func (s *server) GetMenu(req *cp.MenuRequest, stream cp.CoffeeShop_GetMenuServer) error {
	// Menu is a message that contains a list of items
	// Item is a message that contains a name and price
	// MenuRequest is a message that contains a list of categories
	items := []*cp.Item{
		&cp.Item{Id: "1", Name: "Espresso"},
		&cp.Item{Id: "2", Name: "Latte"},
		&cp.Item{Id: "3", Name: "Cappuccino"},
		&cp.Item{Id: "4", Name: "Americano"},
		&cp.Item{Id: "5", Name: "Mocha"},
		&cp.Item{Id: "6", Name: "Macchiato"},
		&cp.Item{Id: "7", Name: "Cold Brew"},
		&cp.Item{Id: "8", Name: "Hot Chocolate"},
		&cp.Item{Id: "9", Name: "Tea"},
		&cp.Item{Id: "10", Name: "Green Tea"},
		&cp.Item{Id: "11", Name: "Black Tea"},
	}

	// Send the items to the client
	for i, _ := range items {
		// Send the item to the client
		stream.Send(&cp.Menu{
			Items: []*cp.Item{items[i]},
		})
	}

	return nil
}

// PlaceOrder is a method that returns a Receipt
func (s *server) PlaceOrder(ctx context.Context, order *cp.Order) (*cp.Receipt, error) {
	// placeholder for order processing logic
	return &cp.Receipt{
		Id: "Receipt100",
	}, nil
}

// GetOrderStatus is a method that returns an OrderStatus
func (s *server) GetOrderStatus(ctx context.Context, receipt *cp.Receipt) (*cp.OrderStatus, error) {
	// placeholder for order status logic
	return &cp.OrderStatus{
		OrderId: receipt.Id,
		Status:  "In Progress",
	}, nil
}

func main() {
	// setup the gRPC server on port 8801
	// start the gRPC server
	listn, err := net.Listen("tcp", ":8801")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cp.RegisterCoffeeShopServer(grpcServer, &server{}) // RegisterCoffeeShopServer is a method that registers the CoffeeShopServer
	log.Println("gRPC server started on port 8801")
	if err := grpcServer.Serve(listn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
