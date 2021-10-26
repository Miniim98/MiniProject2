package main

import (
	"log"
	"net"

	pb "../Chat"
	"google.golang.org/grpc"
)

type server struct {
	network   string
	port      string
	timestamp int32
	pb.UnimplementedChittychatServer
}

func main() {
	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen on port 8008: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterChittychatServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8008: %v", err)

	}
}
