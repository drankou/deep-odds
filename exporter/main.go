package main

import (
	"betsapiScrapper/betsapi"
	"betsapiScrapper/types"
	log "github.com/sirupsen/logrus"
	"time"
)

const STARTING_DATE = "20170610"
const dateLayout = "20060102"

type FootballEvent struct {
	EventInfo  types.Event        `json:"event_info"`
	History    types.EventHistory `json:"history"`
	Odds       types.Odds         `json:"odds"`
	StatsTrend types.StatsTrend   `json:"stats_trend"`
}

func main() {
	betsapiCrawler := betsapi.BetsapiCrawler{}
	err := betsapiCrawler.Init()
	if err != nil {
		log.Fatal(err)
	}

	leagues, err := betsapiCrawler.GetLeagues(types.SoccerId, "", "")
	if err != nil {
		log.Fatal(err)
	}

	for _, league := range leagues {
		log.Infof("Getting matches for %s", league.Name)
		var resultEvents []FootballEvent
		now := time.Now().UTC()
		startDate, err := time.Parse(dateLayout, STARTING_DATE)
		if err != nil {
			log.Error(err)
		}

		for ; startDate.Unix() < now.Unix(); startDate = startDate.AddDate(0, 0, 1) {
			dateStr := startDate.Format(dateLayout)
			log.Infof("Parsing matches for %s", startDate.UTC())
			events, err := betsapiCrawler.GetEndedEvents(types.SoccerId, league.Id, "", "", dateStr, "")
			if err != nil {
				log.Errorf("Exporter: ended events: %s", err)
				continue
			}

			for _, event := range events {
				//get event's stats trend
				statsTrend, err := betsapiCrawler.GetEventStatsTrend(event.Id)
				if err != nil {
					log.Errorf("Exporter: ended events: %s", err)
				}

				//get event's live odds
				odds, err := betsapiCrawler.GetEventOdds(event.Id)
				if err != nil {
					log.Errorf("Exporter: ended events: %s", err)
				}

				//get event's history
				history, err := betsapiCrawler.GetEventHistory(event.Id, "20")
				if err != nil {
					log.Errorf("Exporter: ended events: %s", err)
				}

				resultEvents = append(resultEvents, FootballEvent{
					EventInfo:  event,
					History:    *history,
					Odds:       *odds,
					StatsTrend: *statsTrend,
				})
			}
		}
	}
}