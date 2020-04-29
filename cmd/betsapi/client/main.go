package main

import (
	"google.golang.org/grpc"
	"log"
)

const (
	devAddress = "localhost:50001"
	prodAddress = "todo"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(devAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
}