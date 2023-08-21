package main

import (
	"log"
	"net"
	ecommerce "github.com/durotimicodes/e-commerce-grpc"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	productSvc := &ProductService{} // Implement ProductService
	orderSvc := &OrderService{}     // Implement OrderService
	ecommerce.RegisterProductServiceServer(s, productSvc)
	ecommerce.RegisterOrderServiceServer(s, orderSvc)

	log.Println("Starting gRPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
