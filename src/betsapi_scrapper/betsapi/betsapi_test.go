package betsapi

import (
	"betsapi_scrapper/types"
	"os"
	"testing"
)

func TestBetsapiCrawler_Init(t *testing.T) {
	betsapi := BetsapiWrapper{}
	err := betsapi.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBetsapiCrawler_GetInPlayEvents(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	sportId := types.SoccerId
	inPlayEvents, err := betsapi.GetInPlayEvents(sportId)
	if err != nil {
		t.Fatal(err)
	}
	if len(inPlayEvents) == 0 {
		t.Fatal("There are no in play events for given sport id")
	}

	t.Logf("Number of in-play events: %d", len(inPlayEvents))
}

func TestBetsapiCrawler_GetInPlayEvents_Cached(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	inPlayEvents, err := betsapi.GetInPlayEvents(types.SoccerId)
	if err != nil {
		t.Fatal(err)
	}
	if len(inPlayEvents) == 0 {
		t.Errorf("There are no in play events for given sport id")
	}

	t.Logf("Number of in-play events: %d", len(inPlayEvents))

	inPlayEventsCached, err := betsapi.GetInPlayEvents(types.SoccerId)
	if err != nil {
		t.Fatal(err)
	}
	if len(inPlayEventsCached) == 0 {
		t.Errorf("There are should be cached in-play events")
	}

	t.Logf("Number of cached in-play events: %d", len(inPlayEventsCached))
}

func TestBetsapiCrawler_GetStartingEvents(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	sportId := types.SoccerId
	startingEvents, err := betsapi.GetStartingEvents(sportId, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(startingEvents) == 0 {
		t.Errorf("There are no starting events for given sport id")
	}

	t.Logf("Number of starting events: %d", len(startingEvents))
}

func TestBetsapiCrawler_GetStartingEvents_Cached(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	sportId := types.SoccerId
	startingEvents, err := betsapi.GetStartingEvents(sportId, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(startingEvents) == 0 {
		t.Errorf("There are no starting events for given sport id")
	}

	t.Logf("Number of starting events: %d", len(startingEvents))

	startingEventsCached, err := betsapi.GetStartingEvents(sportId, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(startingEventsCached) == 0 {
		t.Errorf("There are should be cached starting events")
	}

	t.Logf("Number of cached starting events: %d", len(startingEventsCached))
}

func TestBetsapiCrawler_GetEventView(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	eventId := "92149"
	result, err := betsapi.GetEventView(eventId)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}

func TestBetsapiCrawler_GetLeagues(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	allLeagues, err := betsapi.GetLeagues(types.SoccerId, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Number of leagues:", len(allLeagues))
}

func TestBetsapiCrawler_GetEventStatsTrend(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	statsTrend, err := betsapi.GetEventStatsTrend("1981616")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Stats trend: %+v", statsTrend)
}

func TestBetsapiCrawler_GetEndedEvents(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	endedEvents, err := betsapi.GetEndedEvents(types.SoccerId, "", "", "kw", "", "1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ended events: %+v", endedEvents)
}

func TestBetsapiCrawler_GetEndedEvents_Day(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	date := "20191123"
	endedEvents, err := betsapi.GetEndedEvents(types.SoccerId, "", "", "kw", date, "1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Number of ended events on %s: %d", date, len(endedEvents))
	for _, event := range endedEvents {
		event.Clean()
		t.Logf("%+v", event)
	}
}

func TestBetsapiCrawler_GetEventOdds(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	eventOdds, err := betsapi.GetEventOdds("1989042")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Event odds: %+v", eventOdds)
}

func TestBetsapiCrawler_GetEventHistory(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	eventHistory, err := betsapi.GetEventHistory("1989042", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Event history: %+v", eventHistory)
}

func TestBetsapiCrawler_GetUpcomingEvents(t *testing.T) {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	betsapi := GetBetsapiWrapper()

	upcomingEvents, err := betsapi.GetUpcomingEvents(types.SoccerId, "", "", "by", "", "1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Upcoming events: %+v", upcomingEvents)
}
