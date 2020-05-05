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
	// Set up a connection to the tensorflow serving server.
	log.Infof("Connecting to tensorflow serving: %s", os.Getenv("TF_SERVER"))
	client, err := tfclient.NewPredictionClient(os.Getenv("TF_SERVER"))
	if err != nil {
		return err
	}

	d.TensorflowClient = client
	log.Info("Connected.")


	// Set up a connection to the betsapi server.
	log.Infof("Connecting to betsapi server: %s", os.Getenv("BETSAPI_SERVER"))
	conn, err := grpc.Dial(os.Getenv("BETSAPI_SERVER"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	d.BetsapiClient = betsapiTypes.NewBetsapiClient(conn)
	log.Info("Connected.")

	return nil
}

func (d *DeepOddsServer) GetInPlayFootballMatches(ctx context.Context, req *types.InPlayFootballMatchesRequest) (*types.FootballMatchesResponse, error) {
	response := &types.FootballMatchesResponse{}

	betsapiReq := &betsapiTypes.InPlayEventsRequest{
		SportId:  betsapiConstants.SoccerId,
		LeagueId: req.GetLeagueId(),
	}
	eventsResponse, err := d.BetsapiClient.GetInPlayEvents(ctx, betsapiReq)
	if err != nil {
		return nil, err
	}

	for _, event := range eventsResponse.GetEvents() {
		response.Matches = append(response.Matches, &types.FootballMatch{
			Id:         event.GetId(),
			TimeStatus: event.GetTimeStatus(),
			Score:      event.GetScore(),
			HomeTeam:   event.GetHomeTeam().GetName(),
			AwayTeam:   event.GetAwayTeam().GetName(),
			LeagueName: event.GetLeague().GetName(),
			Time: &types.Time{
				Minutes:   event.GetTimer().GetMinutes(),
				Seconds:   event.GetTimer().GetSeconds(),
				AddedTime: event.GetTimer().GetAddedTime(),
			},
		})
	}

	return response, nil
}

func (d *DeepOddsServer) GetFootballMatchPrediction(ctx context.Context, req *types.FootballMatchPredictionRequest) (*types.FootballMatchPredictionResponse, error) {
	response := &types.FootballMatchPredictionResponse{}

	//TODO
	_, err := d.BetsapiClient.GetEventView(ctx, &betsapiTypes.EventViewRequest{EventId: req.GetEventId()})
	if err != nil {
		return nil, err
	}

	return response, nil
}

//TODO
func constructInputFromEvent(event *betsapiTypes.Event) {}
