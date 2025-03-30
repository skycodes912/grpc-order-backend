package main

import (
	"log"
	"net"

	"github.com/skycodes912/grpc-order-backend/internal/service"
	"github.com/skycodes912/grpc-order-backend/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderService := service.NewOrderServiceServer()
	proto.RegisterOrderServiceServer(grpcServer, orderService)

	log.Println("Starting gRPC server on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
