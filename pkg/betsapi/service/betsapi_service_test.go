package service

import (
	"context"
	"github.com/drankou/deep-odds/pkg/betsapi/types"
	"github.com/drankou/deep-odds/pkg/betsapi/types/constants"
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
	os.Setenv("BETSAPI_TOKEN", "25493-N8mbuk79ltAeGs")
	lis = bufconn.Listen(bufSize)

	betsapiService := &BetsapiService{}
	err := betsapiService.Init()
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	types.RegisterBetsapiServer(s, betsapiService)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestBetsapiService_Init(t *testing.T) {
	s := BetsapiService{}
	err := s.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBetsapiService_GetEventView(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := types.NewBetsapiClient(conn)

	req := &types.EventViewRequest{
		EventId: "92149",
	}

	event, err := client.GetEventView(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if event == nil || event.GetId() == "" {
		t.Fatal("empty event view response")
	}
	t.Logf("%+v", event)
}

func TestBetsapiService_GetEventOdds(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.EventOddsRequest{
		EventId: "1989042",
	}

	eventOdds, err := client.GetEventOdds(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if eventOdds == nil || len(eventOdds.GetFullTime()) == 0 {
		t.Fatal("empty event odds response")
	}
	t.Logf("%+v", eventOdds)
}

func TestBetsapiService_GetEventHistory(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.EventHistoryRequest{
		EventId: "1989042",
	}

	eventHistory, err := client.GetEventHistory(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if eventHistory == nil || len(eventHistory.GetHome()) == 0 || len(eventHistory.GetAway()) == 0 {
		t.Fatal("empty event history response")
	}
	t.Logf("%+v", eventHistory)
}

func TestBetsapiService_GetEventStatsTrend(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.EventStatsTrendRequest{
		EventId: "1989042",
	}

	eventStatsTrend, err := client.GetEventStatsTrend(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if eventStatsTrend == nil || len(eventStatsTrend.GetAttacks().GetHome()) == 0 {
		t.Fatal("empty event stats trend response")
	}
	t.Logf("%+v", eventStatsTrend)
}

func TestBetsapiService_GetLeagues(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.LeaguesRequest{
		SportId: constants.SoccerId,
	}
	resp, err := client.GetLeagues(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.GetLeagues()) == 0 {
		t.Fatal("there are no leagues in response")
	}
	t.Logf("Number of leagues in respone: %d", len(resp.GetLeagues()))
}

func TestBetsapiService_GetTeams(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.TeamsRequest{
		SportId: constants.SoccerId,
	}
	resp, err := client.GetTeams(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.GetTeams()) == 0 {
		t.Fatal("there are no teams in response")
	}
	t.Logf("Number of teams in respone: %d", len(resp.GetTeams()))
}

func TestBetsapiService_GetEndedEvents(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.EndedEventsRequest{
		SportId: constants.SoccerId,
	}

	resp, err := client.GetEndedEvents(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.GetEvents()) == 0 {
		t.Fatal("there are no ended events in response")
	}
	t.Logf("Number of ended events in respone: %d", len(resp.GetEvents()))
}

func TestBetsapiService_GetInPlayEvents(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.InPlayEventsRequest{
		SportId: constants.SoccerId,
	}

	resp, err := client.GetInPlayEvents(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.GetEvents()) == 0 {
		t.Fatal("there are no in-play events in response")
	}
	t.Logf("Number of in-play events in respone: %d", len(resp.GetEvents()))
}

func TestBetsapiService_GetUpcomingEvents(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.UpcomingEventsRequest{
		SportId: constants.SoccerId,
	}

	resp, err := client.GetUpcomingEvents(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.GetEvents()) == 0 {
		t.Fatal("there are no upcoming events in response")
	}
	t.Logf("Number of upcoming events in respone: %d", len(resp.GetEvents()))
}
