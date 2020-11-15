package main

import (
	"context"
	"log"
	"os"
	"time"

	"proto"
	"google.golang.org/grpc"
)
const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewClient(conn)

	// Contact the server and print out its response.
	username := defaultName
	name := defaultName

	//COMMAND LINE ARGUMENTS
	if len(os.Args) > 1 {
		name = os.Args[1]
		username = os.Args[2]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.data(ctx, &proto.userRequest{Username: username}, &proto.userRequest{Name:name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())





}