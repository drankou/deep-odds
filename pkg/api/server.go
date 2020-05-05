package api

import (
	"context"
	"github.com/drankou/deep-odds/pkg/api/types"
	betsapiTypes "github.com/drankou/deep-odds/pkg/betsapi/types"
	betsapiConstants "github.com/drankou/deep-odds/pkg/betsapi/types/constants"
	tfclient "github.com/drankou/deep-odds/pkg/tf-serving/client"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
)

type DeepOddsServer struct {
	BetsapiClient    betsapiTypes.BetsapiClient
	TensorflowClient *tfclient.PredictionClient
}

func (d *DeepOddsServer) Init() error {

	client, err := tfclient.NewPredictionClient(os.Getenv("TF_SERVER"))
	if err != nil {
		return err
	}
	d.TensorflowClient = client

	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Getenv("BETSAPI_SERVER"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	d.BetsapiClient = betsapiTypes.NewBetsapiClient(conn)

	return nil
}

func (d DeepOddsServer) GetInPlayFootballEvents(ctx context.Context, req *types.InPlayEventsRequest) (*types.EventsResponse, error) {
	response := &types.EventsResponse{}

	betsapiReq := &betsapiTypes.InPlayEventsRequest{
		SportId:  betsapiConstants.SoccerId,
		LeagueId: req.GetLeagueId(),
	}
	eventsResponse, err := d.BetsapiClient.GetInPlayEvents(ctx, betsapiReq)
	if err != nil {
		return nil, err
	}

	for _, event := range eventsResponse.GetEvents() {
		response.Events = append(response.Events, &types.Event{
			Id:         event.GetId(),
			TimeStatus: event.GetTimeStatus(),
			Score:      event.GetScore(),
			HomeTeam:   event.GetHomeTeam().GetName(),
			AwayTeam:   event.GetAwayTeam().GetName(),
			LeagueName: event.GetLeague().GetName(),
			Timer: &types.Timer{
				Minutes:   event.GetTimer().GetMinutes(),
				Seconds:   event.GetTimer().GetSeconds(),
				AddedTime: event.GetTimer().GetAddedTime(),
			},
		})
	}

	return response, nil
}

func (d DeepOddsServer) GetFootballEventPrediction(ctx context.Context, req *types.EventPredictionRequest) (*types.PredictionResponse, error) {
	response := &types.PredictionResponse{}

	//TODO
	_, err := d.BetsapiClient.GetEventView(ctx, &betsapiTypes.EventViewRequest{EventId: req.GetEventId()})
	if err != nil {
		return nil, err
	}

	return response, nil
}

//TODO
func constructInputFromEvent(event *betsapiTypes.Event){}