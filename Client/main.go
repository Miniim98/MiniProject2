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
	SetUpLog()
	//Calling Connect
	Time.UpTimestamp()
	log.Printf("calling Connect to server at time %d", Time.time)
	response, err := c.Connect(context.Background(), &pb.ConnectionRequest{UserName: name, Timestamp: Time.time})
	if err != nil {
		log.Printf("Failure trying to call Connect %v", err)
	}
	var succes bool
	succes = response.Succeded
	for !succes {
		fmt.Println("Username is taken. Please choose another")
		fmt.Scanln(&name)
		response, err := c.Connect(context.Background(), &pb.ConnectionRequest{UserName: name, Timestamp: Time.time})
		if err != nil {
			log.Printf("Failure trying to call Connect %v", err)
		}
		succes = response.Succeded

	}

	//Go rutines for the chat application
	go SendBroadcastRequest(c)
	go SendPublishRequest(c)
	HandleInterrupt(c)

	for {

	}

}

func SetUpLog() {
	var filename = "log" + name
	LOG_FILE := filename
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
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
		if msg == "" {
			break
		}
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
		log.Printf("Error calling Broadcast : %v", err)
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
			log.Printf("Error when reciving messages: %v", err)
		}
		fmt.Println(response.Message + " at lamporttime " + fmt.Sprint(Time.time))
	}
}

func SendLeaveRequest(client pb.ChittychatClient) {
	_, err := client.Leave(context.Background(), &pb.LeaveRequest{UserName: name, Timestamp: Time.time})
	if err != nil {
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
