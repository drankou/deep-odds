package main

import (
	"betsapiScrapper/betsapi"
	"betsapiScrapper/storage"
	"betsapiScrapper/types"
	log "github.com/sirupsen/logrus"
	"time"
)

var startDate = time.Date(2017, 6, 10, 0, 0, 0, 0, time.UTC)

const dateLayout = "20060102"

func main() {
	//get matches for given league id
	ExportFootballEventsByLeague("8107")
}

func ExportFootballEventsByLeague(leagueId string) {
	//connect to mongodb
	mongo := storage.MongoWrapper{}
	mongoConnectionString := "mongodb://localhost:27017"
	mongo.Connect(mongoConnectionString, storage.DefaultMongoDatabase)

	betsapiCrawler := betsapi.GetBetsapiCrawler()

	events, err := betsapiCrawler.GetEndedEvents(types.SoccerId, leagueId, "", "", "", "")
	if err != nil {
		log.Errorf("Exporter: ended events: %s", err)
	}

	log.Info("Number of ended events for league: ", len(events))

	//fill additional info about events
	for _, event := range events {
		//get event's stats trend
		statsTrend, err := betsapiCrawler.GetEventStatsTrend(event.Id)
		if err != nil {
			log.Errorf("Exporter: event stats: %s: eventId: %s", err, event.Id)
		}

		//get event's live odds
		odds, err := betsapiCrawler.GetEventOdds(event.Id)
		if err != nil {
			log.Errorf("Exporter: event odds: %s: eventId:%s", err, event.Id)
		}

		//get event's history
		history, err := betsapiCrawler.GetEventHistory(event.Id, "20")
		if err != nil {
			log.Errorf("Exporter: ended events: %s: eventId: %s", err, event.Id)
		}

		footballEvent := types.FootballEvent{
			Event:      &event,
			History:    history,
			Odds:       odds,
			StatsTrend: statsTrend,
		}

		_, err = mongo.Insert("football_event", footballEvent)
		if err != nil {
			log.Error(err)
		}

		log.Info("mongo: entry inserted: ", event.Id)
	}
}