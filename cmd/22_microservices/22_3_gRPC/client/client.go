package main

import (
	"context"
	"io"
	"log"
	"time"

	cp "github.com/jaygaha/go-beginner/cmd/22_microservices/22_3_gRPC/coffeeshop_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:8801",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("failed to connect to gRPC server: ", err)
	}

	defer conn.Close()

	client := cp.NewCoffeeShopClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	menuStream, err := client.GetMenu(ctx, &cp.MenuRequest{})
	if err != nil {
		log.Fatal("failed to get menu: ", err)
	}

	done := make(chan bool)

	var items []*cp.Item

	go func() {
		for {
			resp, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("failed to receive menu: %v", err)
			}

			items = append(items, resp.Items...)
			log.Printf("Received menu: %v", resp.Items)
		}
	}()

	<-done

	receipt, err := client.PlaceOrder(ctx, &cp.Order{Items: items})
	if err != nil {
		log.Fatal("failed to place order: ", err)
	}
	log.Printf("Order placed: %v", receipt)

	orderStatus, err := client.GetOrderStatus(ctx, receipt)
	if err != nil {
		log.Fatal("failed to get order status: ", err)
	}
	log.Printf("Order status: %v", orderStatus)
}
