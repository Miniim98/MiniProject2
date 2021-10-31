package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Miniim98/MiniProject2/Chat"
	"google.golang.org/grpc"
)

type server struct {
	timestamp int32
	ch        chan string
	pb.UnimplementedChittychatServer
}

type Client struct {
	username   string
	stream_in  pb.Chittychat_PublishServer
	stream_out pb.Chittychat_BroadcastServer
}

var clients []Client

func (s *server) Connect(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	// saving the client to the client slice
	clients = append(clients, Client{username: in.UserName})
	return &pb.ConnectionResponse{Succeded: true, Timestamp: &pb.LamportTimeStamp{Events: 1}}, nil
}

func (s *server) Publish(stream pb.Chittychat_PublishServer) error {
	//Streaming clientside to recived publish
	for {
		msg, err := stream.Recv()

		for i := 0; i < len(clients); i++ {
			if clients[i].username == msg.UserName {
				clients[i].stream_in = stream
			}
		}

		if err != nil {
			return err
		}
		s.ch <- msg.UserName + ": " + msg.Message
	}
}

func (s *server) Broadcast(in *pb.BroadcastRequest, stream pb.Chittychat_BroadcastServer) error {
	for i := 0; i < len(clients); i++ {
		if clients[i].username == in.UserName {
			clients[i].stream_out = stream
		}
	}
	for _, c := range clients {
		c.stream_out.Send(&pb.BroadcastResponse{Message: in.UserName + " has joined the chat", Timestamp: &pb.LamportTimeStamp{Events: 1}})
	}
	//Streaming serverside to broadcast
	for {
		msg := <-s.ch
		for _, c := range clients {
			if c.stream_out == nil {
				for _, c2 := range clients {
					if c != c2 {
						err := c2.stream_out.Send(&pb.BroadcastResponse{Message: c.username + " left the chat \n", Timestamp: &pb.LamportTimeStamp{Events: 1}})
						if err != nil {
							fmt.Printf("Error when sending message to %v Error : %v", c.username, err)
						}
					}
				}
			}
			err := c.stream_out.Send(&pb.BroadcastResponse{Message: msg, Timestamp: &pb.LamportTimeStamp{Events: 1}})
			if err != nil {
				fmt.Printf("Error when sending message to %v Error : %v", c.username, err)
			}
		}

	}

}

func main() {
	//We listen on port 8008:
	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen on port 8008: %v", err)
	}

	var opts []grpc.ServerOption

	//use grpc to create a new server:
	grpcServer := grpc.NewServer(opts...)

	//register the chittychat server to be the one we have just created
	pb.RegisterChittychatServer(grpcServer, &server{timestamp: 0, ch: make(chan string)})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8008: %v", err)
	}

}
