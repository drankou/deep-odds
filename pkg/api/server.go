package api

import (
	"context"
	"fmt"
	"github.com/drankou/deep-odds/pkg/api/types"
	betsapiTypes "github.com/drankou/deep-odds/pkg/betsapi/types"
	betsapiConstants "github.com/drankou/deep-odds/pkg/betsapi/types/constants"
	tfclient "github.com/drankou/deep-odds/pkg/tf-serving/client"
	"github.com/drankou/deep-odds/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"strconv"
	"strings"
	"time"
)

type DeepOddsServer struct {
	BetsapiClient    betsapiTypes.BetsapiClient
	TensorflowClient *tfclient.PredictionClient
	cache            *utils.Cache
}

func (d *DeepOddsServer) Init() error {
	d.cache = utils.GetLocalCache()

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

	//check if given match was analyzed in last 30 seconds and is presented in cache
	if value, cached := d.cache.Load(fmt.Sprintf("%s_prediction", req.GetEventId())); cached {
		response.Prediction = value.(*types.Prediction)
		return response, nil
	} else {
		defer func() {
			d.cache.Store(
				fmt.Sprintf("%s_prediction", req.GetEventId()),
				response.Prediction,
				&utils.StoreOption{Ttl: time.Second * 30})
		}()
	}

	eventView, err := d.BetsapiClient.GetEventView(ctx, &betsapiTypes.EventViewRequest{EventId: req.GetEventId()})
	if err != nil {
		return nil, err
	}

	eventOdds, err := d.BetsapiClient.GetEventOdds(ctx, &betsapiTypes.EventOddsRequest{EventId: req.GetEventId()})
	if err != nil {
		return nil, err
	}

	input := make([]float32, 0, 25)
	stats := constructInputFromEvent(eventView)
	liveOdds := getLiveOdds(eventOdds)
	startOdds := getStartOdds(eventOdds)

	input = append(input, stats...)
	input = append(input, liveOdds...)
	input = append(input, startOdds...)

	log.Debug("Inputs: ", input)

	prediction, err := d.TensorflowClient.Predict("model_67", input)
	if err != nil {
		return nil, err
	}

	response.Prediction = &types.Prediction{
		HomeWin: float64(prediction.HomeWin),
		Draw:    float64(prediction.Draw),
		AwayWin: float64(prediction.AwayWin),
	}

	return response, nil
}

func constructInputFromEvent(event *betsapiTypes.EventView) []float32 {
	result := make([]float32, 0, 19)

	result = append(result, float32(event.GetTimer().GetMinutes()))
	result = append(result, getGoalsFromScore(event.GetScore())...)
	result = append(result, stringSliceToFloat32(event.GetStats().GetAttacks())...)
	result = append(result, stringSliceToFloat32(event.GetStats().GetDangerousAttacks())...)
	result = append(result, stringSliceToFloat32(event.GetStats().GetOffTarget())...)
	result = append(result, stringSliceToFloat32(event.GetStats().GetOnTarget())...)
	result = append(result, stringSliceToFloat32(event.GetStats().GetCorners())...)
	result = append(result, stringSliceToFloat32(event.GetStats().GetYellowCards())...)
	result = append(result, stringSliceToFloat32(event.GetStats().GetRedCards())...)
	result = append(result, stringSliceToFloat32(event.GetStats().GetSubstitutions())...)

	return result
}

func getGoalsFromScore(score string) []float32 {
	out := make([]float32, 0, 2)
	goals := strings.Split(score, "-")

	if len(goals) == 2 {
		out = append(out, stringToFloat32(goals[0]))
		out = append(out, stringToFloat32(goals[1]))
	}

	return out
}

func stringSliceToFloat32(stringSlice []string) []float32 {
	out := make([]float32, 0, len(stringSlice))
	for _, str := range stringSlice {
		valFloat64, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return out
		}

		out = append(out, float32(valFloat64))
	}

	return out
}

func stringToFloat32(valStr string) float32 {
	valFloat64, err := strconv.ParseFloat(valStr, 32)
	if err != nil {
		return 0
	}

	return float32(valFloat64)
}

//TODO handling missing odds with -1
func getLiveOdds(eventOdds *betsapiTypes.EventOdds) []float32 {
	if len(eventOdds.GetFullTime()) > 0 {
		return []float32{
			float32(eventOdds.GetFullTime()[0].GetHomeOdds()),
			float32(eventOdds.GetFullTime()[0].GetDrawOdds()),
			float32(eventOdds.GetFullTime()[0].GetAwayOdds()),
		}
	} else {
		return []float32{-1, -1, -1}
	}
}

func getStartOdds(eventOdds *betsapiTypes.EventOdds) []float32 {
	if len(eventOdds.GetFullTime()) > 0 {
		startOddsIndex := len(eventOdds.GetFullTime()) - 1
		return []float32{
			float32(eventOdds.GetFullTime()[startOddsIndex].GetHomeOdds()),
			float32(eventOdds.GetFullTime()[startOddsIndex].GetDrawOdds()),
			float32(eventOdds.GetFullTime()[startOddsIndex].GetAwayOdds()),
		}
	} else {
		return []float32{-1, -1, -1}
	}
}
