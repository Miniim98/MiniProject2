package main

import (
	"fmt"
	"log"

	//"net"

	pb "github.itu.dk/jard/Miniproject2.git/Chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewChittychatClient(conn)
	fmt.Println("Connected")

}
