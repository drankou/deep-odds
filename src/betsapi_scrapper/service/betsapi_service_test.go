package service

import (
	"betsapi_scrapper/types"
	"betsapi_scrapper/types/constants"
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"testing"
)

const defaultBetsapiServicePort = ":50001"

func TestBetsapiService_Init(t *testing.T) {
	s := BetsapiService{}
	err := s.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBetsapiService_GetEventView(t *testing.T) {
	go RunBetsapiService()

	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
	go RunBetsapiService()

	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
	go RunBetsapiService()

	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
	go RunBetsapiService()

	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
	go RunBetsapiService()

	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.LeaguesRequest{
		SportId: constants.SoccerId,
	}
	stream, err := client.GetLeagues(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	var leagues []*types.League
	for {
		league, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		leagues = append(leagues, league)
	}

	if len(leagues) == 0 {
		t.Fatal("there are no leagues in response")
	}
	t.Logf("Number of leagues in respone: %d", len(leagues))
}

func TestBetsapiService_GetTeams(t *testing.T) {
	go RunBetsapiService()

	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.TeamsRequest{
		SportId: constants.SoccerId,
	}
	stream, err := client.GetTeams(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	var teams []*types.Team
	for {
		team, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		teams = append(teams, team)
	}

	if len(teams) == 0 {
		t.Fatal("there are no teams in response")
	}
	t.Logf("Number of teams in respone: %d", len(teams))
}

func TestBetsapiService_GetEndedEvents(t *testing.T) {
	go RunBetsapiService()

	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.EndedEventsRequest{
		SportId: constants.SoccerId,
	}
	stream, err := client.GetEndedEvents(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	var endedEvents []*types.Event
	for {
		event, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		endedEvents = append(endedEvents, event)
	}

	if len(endedEvents) == 0 {
		t.Fatal("there are no ended events in response")
	}
	t.Logf("Number of ended events in respone: %d", len(endedEvents))
}

func TestBetsapiService_GetInPlayEvents(t *testing.T) {
	go RunBetsapiService()

	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.InPlayEventsRequest{
		SportId: constants.SoccerId,
	}
	stream, err := client.GetInPlayEvents(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	var inplayEvents []*types.Event
	for {
		event, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		inplayEvents = append(inplayEvents, event)
	}

	if len(inplayEvents) == 0 {
		t.Fatal("there are no in-play events in response")
	}
	t.Logf("Number of in-play events in respone: %d", len(inplayEvents))
}

func TestBetsapiService_GetUpcomingEvents(t *testing.T) {
	go RunBetsapiService()

	conn, err := grpc.Dial("localhost"+defaultBetsapiServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := types.NewBetsapiClient(conn)

	req := &types.UpcomingEventsRequest{
		SportId: constants.SoccerId,
	}
	stream, err := client.GetUpcomingEvents(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	var upcomingEvents []*types.Event
	for {
		event, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		upcomingEvents = append(upcomingEvents, event)
	}

	if len(upcomingEvents) == 0 {
		t.Fatal("there are no upcoming events in response")
	}
	t.Logf("Number of upcoming events in respone: %d", len(upcomingEvents))
}

func RunBetsapiService() {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	grpcServer := grpc.NewServer()

	betsapiService := &BetsapiService{}
	err := betsapiService.Init()
	if err != nil {
		log.Fatal(err)
	}

	types.RegisterBetsapiServer(grpcServer, betsapiService)
	log.Info("BetsapiService is ready.")

	lis, err := net.Listen("tcp", defaultBetsapiServicePort)
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
