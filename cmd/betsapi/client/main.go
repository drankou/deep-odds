package main

import (
	"betsapi_scrapper/types"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

const defaultBetsapiServicePort = ":50001"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.InPlayEventsRequest{
		SportId: "1",
	}
	r, err := client.GetInPlayEvents(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := r.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%+v", msg)
	}
}
