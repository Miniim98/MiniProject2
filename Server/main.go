package main

import (
	"context"
	"log"
	"net"
	"os"
	"sync"

	pb "github.com/Miniim98/MiniProject2/Chat"
	"google.golang.org/grpc"
)

type server struct {
	ch chan string
	pb.UnimplementedChittychatServer
}

type timestamp struct {
	time int32
	mu   sync.Mutex
}

var Time timestamp

type Client struct {
	username   string
	stream_in  pb.Chittychat_PublishServer
	stream_out pb.Chittychat_BroadcastServer
}

func (c *timestamp) UpTimestamp() {
	c.mu.Lock()
	defer c.mu.Unlock()
	Time.time++
}

var clients []Client

func (s *server) Connect(ctx context.Context, in *pb.ConnectionRequest) (*pb.ConnectionResponse, error) {
	if Time.time < in.Timestamp {
		Time.time = in.Timestamp
	}
	Time.UpTimestamp()
	if !(len(clients) == 0) {
		for _, c := range clients {
			if c.username == in.UserName {
				return &pb.ConnectionResponse{Succeded: false, Timestamp: Time.time}, nil
			}
		}
	}
	log.Printf("Recived ConnectRequest from %v at time %d", in.UserName, Time.time)
	// saving the client to the client slice
	clients = append(clients, Client{username: in.UserName})
	return &pb.ConnectionResponse{Succeded: true, Timestamp: Time.time}, nil
}

func (s *server) Publish(stream pb.Chittychat_PublishServer) error {
	//Streaming clientside to recived publish
	for {
		msg, err := stream.Recv()
		if Time.time < msg.Timestamp {
			Time.time = msg.Timestamp
		}
		Time.UpTimestamp()
		log.Printf("Recived Message from %v at time %d", msg.UserName, Time.time)
		//Saving the stream_in
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

	if Time.time < in.Timestamp {
		Time.time = in.Timestamp
	}
	Time.UpTimestamp()
	log.Printf("Recived BroadcastRequest from %v at time %d", in.UserName, Time.time)
	//Saving the stream_out to the right client
	for i := 0; i < len(clients); i++ {
		if clients[i].username == in.UserName {
			clients[i].stream_out = stream
		}
	}

	//When client is connected this is broadcasted to all connected clients
	timeNow := Time.time
	for _, c := range clients {
		Time.UpTimestamp()
		err := c.stream_out.Send(&pb.BroadcastResponse{Message: ">> " + in.UserName + " has joined the chat", Timestamp: timeNow})
		if err != nil {
			log.Printf("Error when sending message to %v Error : %v", c.username, err)
		}
		log.Printf("Send welcome message to %v at time %d", c.username, Time.time)
	}

	//Streaming serverside to broadcast
	for {
		msg := <-s.ch
		timeNow = Time.time
		for _, c := range clients {
			Time.UpTimestamp()
			err := c.stream_out.Send(&pb.BroadcastResponse{Message: ">> " + msg, Timestamp: timeNow})
			if err != nil {
				log.Printf("Error when sending message to %v Error : %v", c.username, err)
			}
			log.Printf("Send BroadCast to %v at time %d", c.username, Time.time)
		}
	}
}

func main() {
	Time.time = 0
	SetUpLog()
	//We listen on port 8008:
	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen on port 8008: %v", err)
	}

	var opts []grpc.ServerOption

	//use grpc to create a new server:
	grpcServer := grpc.NewServer(opts...)

	//register the chittychat server to be the one we have just created
	pb.RegisterChittychatServer(grpcServer, &server{ch: make(chan string)})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8008: %v", err)
	}
}

func SetUpLog() {
	var filename = "log"
	LOG_FILE := filename
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
}
