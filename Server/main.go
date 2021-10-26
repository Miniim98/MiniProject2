package main

import (
	"log"
	"net"

	pb "github.itu.dk/jard/Miniproject2.git/Chat"
	"google.golang.org/grpc"
)

type server struct {
	string net.UnknownNetworkError
	string port
	pb.UnimplementedChittychatServer
}

func main() {
	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen on port 8008: %v", err)
	}

	grpcServer := grpc.NewServer()

	ft.RegisterTimeServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8008: %v", err)

	}
}
