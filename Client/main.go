package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	//"net"

	pb "github.com/Miniim98/MiniProject2/Chat"
	"google.golang.org/grpc"
)

var name string
var tcpServer = flag.String("server", ":8008", "Tcp server")

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())
	conn, err := grpc.Dial(*tcpServer, opts...)
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewChittychatClient(conn)
	SendConnectRequest(c)
	fmt.Println("Connected")
}

func SendConnectRequest(c pb.ChittychatClient) {
	fmt.Println("Choose a username")
	fmt.Scanln(&name)
	response, err := c.Connect(context.Background(), &pb.ConnectionRequest{UserName: name})
	if err != nil {
		fmt.Printf("Failure trying to call Connect %v", err)
		fmt.Println(response)
	}
	go SendBroadcastRequest(c)
	go SendPublishRequest(c)

	for {

	}

}

func SendPublishRequest(c pb.ChittychatClient) {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("Error during reading input")
		}
		msg := scanner.Text()
		stream, err := c.Publish(context.Background())
		if err != nil {
			fmt.Printf("Error calling Publish : %v", err)
		}
		if err := stream.Send(&pb.PublishRequest{Message: msg, UserName: name, Timestamp: &pb.LamportTimeStamp{Events: 1}}); err != nil {
			fmt.Printf("Error sending stream to Publish : %v", err)
		}

	}
}

func SendBroadcastRequest(c pb.ChittychatClient) {
	stream, err := c.Broadcast(context.Background(), &pb.BroadcastRequest{UserName: name, Timestamp: &pb.LamportTimeStamp{Events: 1}})
	if err != nil {
		fmt.Printf("Error calling Broadcast : %v", err)
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			fmt.Printf("Error when reciving messages: %v", err)
		}
		fmt.Println(response.Message + " at lamporttime " + response.Timestamp.String())
	}
}
