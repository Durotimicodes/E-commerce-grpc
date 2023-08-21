package main

import (
	"log"

	ecommerce "github.com/durotimicodes/e-commerce-grpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	productService := ecommerce.NewProductServiceClient(conn)
	orderService := ecommerce.NewOrderServiceClient(conn)

	// Use productService and orderService to interact with the services
	// Implement client-side logic here
}
