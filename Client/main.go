package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	//"net"

	pb "github.com/Miniim98/MiniProject2/Chat"
	"google.golang.org/grpc"
)

var name string
var tcpServer = flag.String("server", ":8008", "Tcp server")

func main() {
	//Dialing the server
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(*tcpServer, opts...)
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewChittychatClient(conn)
	//Calling SendConnectRequest that then calls the others
	SendConnectRequest(c)
}

func SendConnectRequest(c pb.ChittychatClient) {
	//Scanning input
	fmt.Println("Choose a username")
	fmt.Scanln(&name)
	//Calling Connect
	response, err := c.Connect(context.Background(), &pb.ConnectionRequest{UserName: name})
	if err != nil {
		fmt.Printf("Failure trying to call Connect %v", err)
		fmt.Println(response)
	}
	//Go rutines for the chat application
	go SendBroadcastRequest(c)
	go SendPublishRequest(c)
	HandleInterrupt(c)

	for { 

	}

}

func SendPublishRequest(c pb.ChittychatClient) {

	for {
		//Reading input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("Error during reading input")
		}
		msg := scanner.Text()
		//Calling Publish / opening stream
		stream, err := c.Publish(context.Background())
		if err != nil {
			fmt.Printf("Error calling Publish : %v", err)
		}
		//Sending message through stream
		if err := stream.Send(&pb.PublishRequest{Message: msg, UserName: name, Timestamp: &pb.LamportTimeStamp{Events: 1}}); err != nil {
			fmt.Printf("Error sending stream to Publish : %v", err)
		}

	}
}

func SendBroadcastRequest(c pb.ChittychatClient) {
	//Opening stream
	stream, err := c.Broadcast(context.Background(), &pb.BroadcastRequest{UserName: name, Timestamp: &pb.LamportTimeStamp{Events: 1}})
	if err != nil {
		fmt.Printf("Error calling Broadcast : %v", err)
	}
	//Listening and reciving messeges
	for {
		response, err := stream.Recv()
		if err != nil {
			fmt.Printf("Error when reciving messages: %v", err)
		}
		fmt.Println(response.Message + " at lamporttime " + response.Timestamp.String())
	}
}

func SendLeaveRequest(client pb.ChittychatClient) {
	_, err := client.Leave(context.Background(), &pb.LeaveRequest{UserName: name, Timestamp: &pb.LamportTimeStamp{Events: 1}})
	if (err != nil) {
		fmt.Printf("Trouble sending leave request: %v", err)
	}
}

func HandleInterrupt(client pb.ChittychatClient) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		<-c
		SendLeaveRequest(client)
		os.Exit(0)
	}()

}
