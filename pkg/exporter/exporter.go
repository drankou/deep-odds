package exporter

import (
	"context"
	"errors"
	"fmt"
	"github.com/drankou/deep-odds/pkg/betsapi/types"
	"github.com/drankou/deep-odds/pkg/betsapi/types/constants"
	"github.com/drankou/deep-odds/pkg/storage"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"strconv"
	"strings"
	"time"
)

// Exporter for storing data from betsapi API to mongo database
type Exporter struct {
	mongo   *storage.MongoWrapper
	betsapi types.BetsapiClient
}

func (e *Exporter) Init() error {
	//connect to mongo db
	e.mongo = storage.GetMongoWrapper()

	// Set up a connection to the betsapi server.
	log.Info("Connecting to betsapi server")
	conn, err := grpc.Dial(os.Getenv("BETSAPI_SERVER"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}

	e.betsapi = types.NewBetsapiClient(conn)

	return nil
}

func (e *Exporter) ExportFootballEventsFromDate(from int64, to int64) {
	//format YYYYMMDD, eg: 20180814 (min 20160901)
	for timestamp := from; timestamp < to; timestamp += 86400 {
		date := time.Unix(timestamp, 0).Truncate(24 * time.Hour).Format("2006-01-02")
		log.Infof("Exporting events for day: %v", date)

		for page := 1; page > 0; {
			req := &types.EndedEventsRequest{
				SportId: constants.SoccerId,
				Day:     strings.ReplaceAll(date, "-", ""),
				Page:    strconv.Itoa(page),
			}

			resp, err := e.betsapi.GetEndedEvents(context.Background(), req)
			if err != nil {
				log.Errorf("Exporter: ended events: %s", err)
			}

			err = e.SaveFootballEvents(resp.GetEvents())
			if err != nil {
				log.Error(err)
				return
			}

			page = int(resp.GetNextPage())
		}
	}
}

func (e *Exporter) SaveFootballEvents(events []*types.EventView) error {
	for _, event := range events {
		footballEvent, err := e.GetFootballEventById(event.GetId())
		if err != nil {
			log.Error(err)
			continue
		}

		if footballEvent != nil {
			//save event to mongo
			_, err := e.mongo.Insert("football_event_validation", footballEvent)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *Exporter) GetFootballEventById(eventId string) (*types.FootballEvent, error) {
	//calling event view because ended events could return event with missing data
	eventView, err := e.betsapi.GetEventView(context.Background(), &types.EventViewRequest{EventId: eventId})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Exporter: GetFootballEventById: %s", err))
	}

	//fill additional info about event
	footballEvent := e.EventViewToFootballEvent(eventView)

	//clean unnecessary data
	footballEvent.Clean()

	return footballEvent, nil
}

//Request for stats, odds and history of given event
func (e *Exporter) EventViewToFootballEvent(event *types.EventView) *types.FootballEvent {
	//get event's stats trend
	statsTrendRequest := &types.EventStatsTrendRequest{
		EventId: event.GetId(),
	}
	statsTrend, err := e.betsapi.GetEventStatsTrend(context.Background(), statsTrendRequest)
	if err != nil {
		log.Errorf("Exporter: event stats: %s: eventId: %s", err, event.GetId())
	}

	//get event's live odds
	oddsRequest := &types.EventOddsRequest{
		EventId: event.GetId(),
	}
	odds, err := e.betsapi.GetEventOdds(context.Background(), oddsRequest)
	if err != nil {
		log.Errorf("Exporter: event odds: %s: eventId:%s", err, event.GetId())
	}

	//get event's history
	historyReq := &types.EventHistoryRequest{
		EventId: event.GetId(),
		Qty:     "20",
	}
	history, err := e.betsapi.GetEventHistory(context.Background(), historyReq)
	if err != nil {
		log.Errorf("Exporter: event history: %s: eventId: %s", err, event.GetId())
	}

	//construct football event
	footballEvent := &types.FootballEvent{
		Event:      event,
		History:    history,
		Odds:       odds,
		StatsTrend: statsTrend,
	}

	return footballEvent
}
