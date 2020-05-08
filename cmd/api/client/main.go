package main

import (
	"context"
	"github.com/drankou/deep-odds/pkg/api/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const devAddress = "localhost:8200"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(devAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	getInplayEvents(conn)
	getPredictionForEvent(conn)
}

func getInplayEvents(conn *grpc.ClientConn) {
	client := types.NewDeepOddsClient(conn)

	req := &types.InPlayFootballMatchesRequest{}
	eventsResponse, err := client.GetInPlayFootballMatches(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range eventsResponse.GetMatches() {
		log.Printf("%+v", event)
	}
}

func getPredictionForEvent(conn *grpc.ClientConn) {
	client := types.NewDeepOddsClient(conn)

	req := &types.FootballMatchPredictionRequest{
		EventId: "2351755",
	}
	predictionResponse, err := client.GetFootballMatchPrediction(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", predictionResponse.GetPrediction())
}
