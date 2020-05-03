package exporter

import (
	"github.com/drankou/deep-odds/pkg/betsapi/types"
	"github.com/drankou/deep-odds/pkg/storage"
)

//represents endpoint for all requests on betsapi data
type Exporter struct {
	//TODO cache
	Mongo   *storage.MongoWrapper
	Betsapi *types.BetsapiClient
}

func (e *Exporter) Init() error {
	e.Mongo = storage.GetMongoWrapper()
	//e.Betsapi = betsapi.GetBetsapiWrapper()

	return nil
}

//func (e *Exporter) ExportFootballEventsByLeague(leagueId string) {
//	events, err := e.Betsapi.GetEndedEvents(types.SoccerId, leagueId, "", "", "", "")
//	if err != nil {
//		log.Errorf("Exporter: ended events: %s", err)
//	}
//
//	log.Infof("Crawling ended events for league: %s", leagueId)
//	log.Infof("Number of ended events for league: %d", len(events))
//
//	//fill additional info about events
//	for _, event := range events {
//		footballEvent := e.GetFootballEventById(event.Id)
//
//		if footballEvent != nil {
//			//save event to mongo
//			err := e.SaveFootballEventToMongo(footballEvent)
//			if err != nil {
//				log.Errorf("Mongo error: %s", err)
//			}
//		}
//	}
//}
//
//func (e *Exporter) GetFootballEventById(eventId string) *types.FootballEvent {
//	//calling event view because ended events could return event with missing data
//	event, err := e.Betsapi.GetEventView(eventId)
//	if err != nil {
//		log.Errorf("Exporter: event view: %s", err)
//		return nil
//	}
//
//	//fill additional info about event
//	footballEvent := e.EventToFootballEvent(event)
//
//	//clean unnecessary data
//	footballEvent.Clean()
//
//	return footballEvent
//}
//
////Request for stats, odds and history of given event
//func (e *Exporter) EventToFootballEvent(event *types.Event) *types.FootballEvent {
//	statsTrend, err := e.Betsapi.GetEventStatsTrend(event.Id)
//	if err != nil {
//		log.Errorf("Exporter: event stats: %s: eventId: %s", err, event.Id)
//	}
//
//	//get event's live odds
//	odds, err := e.Betsapi.GetEventOdds(event.Id)
//	if err != nil {
//		log.Errorf("Exporter: event odds: %s: eventId:%s", err, event.Id)
//	}
//
//	//get event's history
//	history, err := e.Betsapi.GetEventHistory(event.Id, "20")
//	if err != nil {
//		log.Errorf("Exporter: event history: %s: eventId: %s", err, event.Id)
//	}
//
//	footballEvent := &types.FootballEvent{
//		Event:      event,
//		History:    history,
//		Odds:       odds,
//		StatsTrend: statsTrend,
//	}
//
//	return footballEvent
//}
//
//func (e *Exporter) ExportFootballEventsFromDate(from int64, to int64) {
//	//format YYYYMMDD, eg: 20180814 (min 20160901)
//
//	for timestamp := from; timestamp < to; timestamp += 86400 {
//		year, month, day := time.Unix(timestamp, 0).Truncate(24 * time.Hour).Date()
//
//		var monthStr string
//		var dayStr string
//
//		if month < 10 {
//			monthStr = fmt.Sprintf("0%d", month)
//		} else {
//			monthStr = fmt.Sprintf("%d", month)
//		}
//
//		if day < 10 {
//			dayStr = fmt.Sprintf("0%d", day)
//		} else {
//			dayStr = fmt.Sprintf("%d", day)
//		}
//
//		betsapiDay := fmt.Sprintf("%d%s%s", year, monthStr, dayStr)
//		log.Print(betsapiDay)
//
//		events, err := e.Betsapi.GetEndedEvents(types.SoccerId, "", "", "", betsapiDay, "")
//		if err != nil {
//			log.Errorf("Exporter: ended events: %s", err)
//		}
//
//		log.Infof("Number of ended events for %s: %d", betsapiDay, len(events))
//
//		//fill additional info about events
//		for _, event := range events {
//			footballEvent := e.GetFootballEventById(event.Id)
//
//			if footballEvent != nil {
//				//save event to mongo
//				err := e.SaveFootballEventToMongo(footballEvent)
//				if err != nil {
//					log.Errorf("Mongo error: %s", err)
//				}
//			}
//		}
//	}
//}
//
//func (e *Exporter) UpdateFootballEventStatsTrend(footballEvent *types.FootballEvent) error {
//	statsTrend, err := e.Betsapi.GetEventStatsTrend(footballEvent.Event.Id)
//	if err != nil {
//		return errors.Errorf("Exporter: event stats: %s: eventId: %s", err, footballEvent.Event.Id)
//	}
//
//	log.Info("mongo: entry updated: ", footballEvent.Event.Id)
//	footballEvent.StatsTrend = statsTrend
//
//	return nil
//}
//
//func (e *Exporter) SaveFootballEventToMongo(footballEvent *types.FootballEvent) error {
//	_, err := e.Mongo.Insert("football_event", footballEvent)
//	if err != nil {
//		return err
//	}
//
//	//log.Info("mongo: entry inserted: ", footballEvent.Event.Id)
//
//	return nil
//}
