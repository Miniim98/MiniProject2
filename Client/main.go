package main

import (
	"context"
	"fmt"
	"log"

	//"net"

	pb "github.com/Miniim98/MiniProject2/Chat"
	"google.golang.org/grpc"
)

var name string

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8008", grpc.WithInsecure())
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
		fmt.Printf("Failure trying to call Connect %s" , err)
		fmt.Println(response)
	}

}
