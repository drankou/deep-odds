package main

import (
	"context"
	"github.com/drankou/deep-odds/pkg/betsapi/types"
	"github.com/drankou/deep-odds/pkg/betsapi/types/constants"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	devAddress  = "localhost:8080"
	prodAddress = "todo"
)

func main() {
	log.Info("Connecting to betsapi server")
	// Set up a connection to the server.
	conn, err := grpc.Dial(devAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	getInplayEvents(conn)
}

func getInplayEvents(conn *grpc.ClientConn) {
	c := types.NewBetsapiClient(conn)

	req := &types.InPlayEventsRequest{
		SportId: constants.SoccerId,
	}

	response, err := c.GetInPlayEvents(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range response.GetEvents() {
		log.Print(event.GetId())
	}
}
