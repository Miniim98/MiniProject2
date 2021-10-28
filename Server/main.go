package main

import (
	"log"
	"net"

	pb "github.itu.dk/jard/Miniproject2.git/Chat"
	"google.golang.org/grpc"
)

type server struct {
	network   string
	port      string
	timestamp int32
	pb.UnimplementedChittychatServer
}

func main() {
	//We listen on port 8008:
	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen on port 8008: %v", err)
	}

	//use grpc to create a new server:
	grpcServer := grpc.NewServer()

	//register the chittychat server to be the one we have just created
	pb.RegisterChittychatServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8008: %v", err)
	}

}
