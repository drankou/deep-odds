package betsapi

import (
	"betsapiScrapper/types"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestBetsapiCrawler_Init(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBetsapiCrawler_GetInPlayEvents(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}

	sportId := types.SoccerId

	inPlayEvents := betsapi.GetInPlayEvents(sportId)
	if len(inPlayEvents) == 0 {
		t.Errorf("There are no in play events for given sport id")
	}

	t.Logf("Number of in-play events: %d", len(inPlayEvents))
}

func TestBetsapiCrawler_GetInPlayEvents_Cached(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}

	log.SetLevel(log.DebugLevel)

	sportId := types.SoccerId
	inPlayEvents := betsapi.GetInPlayEvents(sportId)
	if len(inPlayEvents) == 0 {
		t.Errorf("There are no in play events for given sport id")
	}

	t.Logf("Number of in-play events: %d", len(inPlayEvents))

	inPlayEventsCached := betsapi.GetInPlayEvents(sportId)
	if len(inPlayEventsCached) == 0 {
		t.Errorf("There are should be cached in-play events")
	}

	t.Logf("Number of cached in-play events: %d", len(inPlayEventsCached))
}

func TestBetsapiCrawler_GetStartingEvents(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}

	sportId := types.SoccerId

	startingEvents := betsapi.GetStartingEvents(sportId, 10)
	if len(startingEvents) == 0 {
		t.Errorf("There are no starting events for given sport id")
	}

	t.Logf("Number of starting events: %d", len(startingEvents))
}

func TestBetsapiCrawler_GetStartingEvents_Cached(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	sportId := types.SoccerId

	startingEvents := betsapi.GetStartingEvents(sportId, 10)
	if len(startingEvents) == 0 {
		t.Errorf("There are no starting events for given sport id")
	}

	t.Logf("Number of starting events: %d", len(startingEvents))

	startingEventsCached := betsapi.GetStartingEvents(sportId, 10)
	if len(startingEventsCached) == 0 {
		t.Errorf("There are should be cached starting events")
	}

	t.Logf("Number of cached starting events: %d", len(startingEventsCached))
}

func TestBetsapiCrawler_GetEventView(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	eventId := "92149"
	result := betsapi.GetEventView(eventId)
	footballEvent := result.(types.FootballEvent)
	t.Logf("%+v", footballEvent)
}

func TestBetsapiCrawler_GetLeagues(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	allLeagues := betsapi.GetLeagues(types.SoccerId, "", "")
	t.Log("Number of leagues:", len(allLeagues))
}

func TestBetsapiCrawler_GetEventStatsTrend(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	statsTrend := betsapi.GetEventStatsTrend("1981616")
	t.Logf("Stats trend: %+v", statsTrend)
}

func TestBetsapiCrawler_GetEndedEvents(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	endedEvents := betsapi.GetEndedEvents(types.SoccerId, "","","by","","1")
	t.Logf("Ended events: %+v", endedEvents)
}

func TestBetsapiCrawler_GetEndedEvents_Day(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	date := "20190610"
	endedEvents := betsapi.GetEndedEvents(types.SoccerId, "","","",date,"1")
	t.Logf("Number of ended events on %s: %d", date, len(endedEvents))
}

func TestBetsapiCrawler_GetEventOdds(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	eventOdds, err := betsapi.GetEventOdds("1989042")
	if err != nil{
		t.Fatal(err)
	}

	t.Logf("Event odds: %+v", eventOdds)
}

func TestBetsapiCrawler_GetEventHistory(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	eventHistory, err := betsapi.GetEventHistory("1989042", "")
	if err != nil{
		t.Fatal(err)
	}

	t.Logf("Event history: %+v", eventHistory)
}

func TestBetsapiCrawler_GetUpcomingEvents(t *testing.T) {
	betsapi := BetsapiCrawler{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
	log.SetLevel(log.DebugLevel)

	upcomingEvents := betsapi.GetEndedEvents(types.SoccerId, "","","by","","1")
	t.Logf("Upcoming events: %+v", upcomingEvents)
}