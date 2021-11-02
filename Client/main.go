package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	pb "github.com/Miniim98/MiniProject2/Chat"
	"google.golang.org/grpc"
)

var name string
var tcpServer = flag.String("server", ":8008", "Tcp server")

type timestamp struct {
	time int32
	mu   sync.Mutex
}

var Time timestamp

func (c *timestamp) UpTimestamp() {
	c.mu.Lock()
	defer c.mu.Unlock()
	Time.time++
}

func main() {
	Time.time = 0
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
	Time.UpTimestamp()
	log.Printf("calling Connect to server at time %d", Time.time)
	response, err := c.Connect(context.Background(), &pb.ConnectionRequest{UserName: name, Timestamp: Time.time})
	if err != nil {
		fmt.Printf("Failure trying to call Connect %v", err)
		fmt.Println(response)
	}

	//Go rutines for the chat application
	go SendBroadcastRequest(c)
	go SendPublishRequest(c)

	for {

	}

}

func SendPublishRequest(c pb.ChittychatClient) {
	for {
		//Reading input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Err() != nil {
			log.Println("Error during reading input")
		}
		msg := scanner.Text()
		Time.UpTimestamp()
		log.Printf("Calling Publish at time %d", Time.time)
		//Calling Publish / opening stream
		stream, err := c.Publish(context.Background())
		if err != nil {
			log.Printf("Error calling Publish : %v", err)
		}
		//Sending message through stream
		Time.UpTimestamp()
		log.Printf("Sending message to server through stream at time %d", Time.time)
		if err := stream.Send(&pb.PublishRequest{Message: msg, UserName: name, Timestamp: Time.time}); err != nil {
			log.Printf("Error sending stream to Publish : %v", err)
		}

	}
}

func SendBroadcastRequest(c pb.ChittychatClient) {
	//Opening stream
	Time.UpTimestamp()
	log.Printf("Calling broadcast at time %d", Time.time)
	stream, err := c.Broadcast(context.Background(), &pb.BroadcastRequest{UserName: name, Timestamp: Time.time})
	if err != nil {
		fmt.Printf("Error calling Broadcast : %v", err)
	}
	//Listening and reciving messeges
	for {
		response, err := stream.Recv()
		if Time.time < response.Timestamp {
			Time.time = response.Timestamp
		}
		Time.UpTimestamp()
		log.Printf("Reciving %v at time %d", response.Message, Time.time)
		if err != nil {
			fmt.Printf("Error when reciving messages: %v", err)
		}
		fmt.Println(response.Message + " at lamporttime " + fmt.Sprint(Time.time))
	}
}
