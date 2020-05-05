package api

import (
	"context"
	"github.com/drankou/deep-odds/pkg/api/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"os"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	os.Setenv("TF_SERVER", "localhost:8500")
	os.Setenv("BETSAPI_SERVER", "localhost:50001")

	lis = bufconn.Listen(bufSize)
	deepOdds := &DeepOddsServer{}
	err := deepOdds.Init()
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	types.RegisterDeepOddsServer(s, deepOdds)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestDeepOddsServer_GetInPlayFootballEvents(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := types.NewDeepOddsClient(conn)

	req := &types.InPlayFootballMatchesRequest{}

	matchesResponse, err := client.GetInPlayFootballMatches(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Number of football events in-play: ", len(matchesResponse.GetMatches()))
}

func TestDeepOddsServer_GetFootballEventPrediction(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := types.NewDeepOddsClient(conn)

	req := &types.FootballMatchPredictionRequest{
		EventId: "111",
	}

	predictionResponse, err := client.GetFootballMatchPrediction(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Prediction: %+v", predictionResponse.GetPrediction())
}
