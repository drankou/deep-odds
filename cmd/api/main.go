package api

import (
	"context"
	"github.com/drankou/deep-odds/pkg/api/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Getenv("BETSAPI_SERVER"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	getInplayEvents(conn)
	//getPredictionForEvent(conn)
}

func getInplayEvents(conn *grpc.ClientConn) {
	client := types.NewDeepOddsClient(conn)

	req := &types.InPlayEventsRequest{}
	eventsResponse, err := client.GetInPlayFootballEvents(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range eventsResponse.GetEvents() {
		log.Printf("%+v", event)
	}
}

func getPredictionForEvent(conn *grpc.ClientConn) {
	client := types.NewDeepOddsClient(conn)

	req := &types.EventPredictionRequest{
		EventId: "111",
	}
	predictionResponse, err := client.GetFootballEventPrediction(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", predictionResponse.GetPrediction())
}
