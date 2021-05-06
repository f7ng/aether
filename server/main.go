package main

import (
	"log"
	"net"

	pb "github.com/f7ng/aether/protobuf"
	userServer "github.com/f7ng/aether/server/user"
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
	pb.RegisterUserServer(s, &userServer.Server{})
	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve on listener: listener=%v", listener)
	}
}