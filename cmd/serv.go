package main

import (
	"github.com/sashabaranov/pike/backend"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	backendServer := backend.NewServer()
	defer backendServer.Cleanup()

	grpcServer := grpc.NewServer()
	backend.RegisterBackendServer(
		grpcServer,
		backendServer,
	)

	log.Print("Server started")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
