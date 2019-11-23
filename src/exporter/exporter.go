package exporter

import (
	"betsapiScrapper/src/betsapi"
	"betsapiScrapper/src/storage"
	"betsapiScrapper/src/types"
	log "github.com/sirupsen/logrus"
)

//represents endpoint for all requests on betsapi data
type Exporter struct {
	//TODO cache
	Mongo   *storage.MongoWrapper
	Betsapi *betsapi.BetsapiWrapper
}

func (e *Exporter) Init() error {
	e.Mongo = storage.GetMongoWrapper()
	e.Betsapi = betsapi.GetBetsapiWrapper()

	return nil
}

func (e *Exporter) GetFootballEventsByLeague(leagueId string) []*types.FootballEvent {
	var fooballEvents []*types.FootballEvent

	events, err := e.Betsapi.GetEndedEvents(types.SoccerId, leagueId, "", "", "", "")
	if err != nil {
		log.Errorf("Exporter: ended events: %s", err)
	}

	log.Info("Number of ended events for league: ", len(events))

	//fill additional info about events
	for _, event := range events {
		footballEvent := e.EventToFootballEvent(&event)
		fooballEvents = append(fooballEvents, footballEvent)
	}

	return fooballEvents
}

func (e *Exporter) GetFootballEventById(eventId string) *types.FootballEvent {
	event, err := e.Betsapi.GetEventView(eventId)
	if err != nil {
		log.Errorf("Exporter: event view: %s", err)
	}

	//fill additional info about event
	footballEvent := e.EventToFootballEvent(event)

	return footballEvent
}

func (e *Exporter) EventToFootballEvent(event *types.Event) *types.FootballEvent {
	statsTrend, err := e.Betsapi.GetEventStatsTrend(event.Id)
	if err != nil {
		log.Errorf("Exporter: event stats: %s: eventId: %s", err, event.Id)
	}

	//get event's live odds
	odds, err := e.Betsapi.GetEventOdds(event.Id)
	if err != nil {
		log.Errorf("Exporter: event odds: %s: eventId:%s", err, event.Id)
	}

	//get event's history
	history, err := e.Betsapi.GetEventHistory(event.Id, "20")
	if err != nil {
		log.Errorf("Exporter: event history: %s: eventId: %s", err, event.Id)
	}

	footballEvent := &types.FootballEvent{
		Event:      event,
		History:    history,
		Odds:       odds,
		StatsTrend: statsTrend,
	}

	return footballEvent
}

func (e *Exporter) SaveFootballEvent(footballEvent types.FootballEvent) {
	_, err := e.Mongo.Insert("football_event", footballEvent)
	if err != nil {
		log.Error(err)
	}

	log.Info("mongo: entry inserted: ", footballEvent.Event.Id)
}
