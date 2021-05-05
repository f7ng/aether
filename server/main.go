package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = "50051"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port: port=%v", port)
	}

	s := grpc.NewServer()
	
}