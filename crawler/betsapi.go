package crawler

import (
	"betsapiScrapper/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

const API_URL string = "https://api.betsapi.com" //https://app.bsportsfan.com
const API_TOKEN string = "25493-YejsbOXuvVD4Ue"

const (
	//https://betsapi.com/docs/events/inplay.html
	InPlayEventsUrl = API_URL + "/v1/events/inplay"

	//https://betsapi.com/docs/events/upcoming.html
	UpcomingEventsUrl = API_URL + "/v2/events/upcoming"

	//https://betsapi.com/docs/events/ended.html
	EndedEventsUrl = API_URL + "/v2/events/ended"

	//https://betsapi.com/docs/events/search.html
	EventsSearchUrl = API_URL + "/v1/events/search"

	//https://betsapi.com/docs/events/view.html
	EventViewUrl = API_URL + "/v1/event/view"

	//https://betsapi.com/docs/events/history.html
	EventHistoryUrl = API_URL + "/v1/event/history"

	//https://betsapi.com/docs/events/odds_summary.html
	EventOddsSummaryUrl = API_URL + "/v2/event/odds/summary"

	//https://betsapi.com/docs/events/odds.html
	EventOddsUrl = API_URL + "/v2/event/odds"

	//https://betsapi.com/docs/events/stats_trend.html
	EventStatsTrendUrl = API_URL + "/v1/event/stats_trend"

	//https://betsapi.com/docs/events/lineup.html
	EventLineUpUrl = API_URL + "/v1/event/lineup?id="

	//https://betsapi.com/docs/events/videos.html
	EventVideosUrl = API_URL + "/v1/event/videos?id="

	//Returns all leagues
	//https://betsapi.com/docs/events/league.html
	LeagueUrl = API_URL + "/v1/league"

	//https://betsapi.com/docs/events/league_table.html
	LeagueTable = API_URL + "/v2/league/table?league_id="

	//https://betsapi.com/docs/events/league_toplist.html
	LeagueTopList = API_URL + "/v1/league/toplist?league_id="

	//Returns all teams
	//https://betsapi.com/docs/events/team.html
	TeamUrl = API_URL + "/v1/team?sport_id"

	//https://betsapi.com/docs/events/team_squad.html
	TeamSquadUrl = API_URL + "/v1/team/squad?team_id="

	//https://betsapi.com/docs/events/player.html
	PlayerUrl = API_URL + "/v1/player?player_id="

	//https://betsapi.com/docs/events/tennis_ranking.html
	TennisRanking = API_URL + "/v1/tennis/ranking?type_id="
)

type BetsapiCrawler struct {
	Client           *http.Client
	Context          context.Context
	SportId          string
	Cache            *types.Cache
	InPlayEvents     []types.Event
	FootballEvents   []types.FootballEvent
	TennisEvents     []types.TennisEvent
	BasketballEvents []types.BasketballEvent
}

func (betsapi *BetsapiCrawler) Init() error {
	betsapi.Client = &http.Client{}
	betsapi.Context = context.Background()

	//initialize cache
	betsapi.Cache = &types.Cache{}
	betsapi.Cache.ResetInterval = "every 1m"
	betsapi.Cache.ResetFunc = betsapi.Cache.Clear

	return nil
}

//Returns in-play events defined by sportId,
func (betsapi *BetsapiCrawler) GetInPlayEvents(sportId string) []types.Event {
	log.Info("Getting In-Play events...")
	var inPlayEvents []types.Event

	//check in-play events in cache
	if val, exist := betsapi.Cache.Load(fmt.Sprintf("inplay_%s", sportId)); exist {
		log.Debugln("InPlay: Returning value from cache")
		return val.([]types.Event)
	}

	req, err := http.NewRequest("GET", InPlayEventsUrl, nil)
	if err != nil {
		log.Error(err)
	}

	//encode query parameters
	q := req.URL.Query()
	q.Add("token", API_TOKEN)
	q.Add("sport_id", sportId)
	req.URL.RawQuery = q.Encode()

	resp, err := betsapi.Client.Do(req)
	if err != nil {
		log.Error(err)
	}

	var betsapiResponse types.BetsapiInplayResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			log.Error(err)
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			log.Error(err)
		}

		if betsapiResponse.Success == 1 {
			//save successfull response to cache
			log.Debugln("InPlay: Storing value to cache")
			betsapi.Cache.Store(fmt.Sprintf("inplay_%s", sportId), betsapiResponse.Results)

			return betsapiResponse.Results
		} else {
			log.Errorf("Error: %d: unsuccessful API response: /events/inplay", resp.StatusCode)
		}
	} else {
		log.Errorf("Error: %d: request: /events/inplay", resp.StatusCode)
	}

	return inPlayEvents
}

func (betsapi *BetsapiCrawler) GetUpcomingEvents(sportId, leagueId, teamId, countryCode, day, page string) ([]types.Event, error) {
	var upcomingEvents []types.Event
	req, err := http.NewRequest("GET", UpcomingEventsUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := req.URL.Query()
	q.Add("token", API_TOKEN)
	q.Add("sport_id", sportId)
	q.Add("league_id", leagueId)
	q.Add("team_id", teamId)
	q.Add("cc", countryCode)
	q.Add("day", day)
	q.Add("page", page)

	req.URL.RawQuery = q.Encode()
	resp, err := betsapi.Client.Do(req)
	if err != nil {
		log.Error(err)
	}

	var betsapiResponse types.BetsapEventsPagerResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			return nil, err
		}

		if betsapiResponse.Success == 1 {
			log.Info("Total events: ", betsapiResponse.Pager.Total)
			upcomingEvents = append(upcomingEvents, betsapiResponse.Results...)

			actualPage := betsapiResponse.Pager.Page
			perPage := betsapiResponse.Pager.PerPage
			total := betsapiResponse.Pager.Total

			if actualPage == 100 {
				return upcomingEvents, errors.Errorf("Error %d: max page limit", resp.StatusCode)
			}

			if actualPage*perPage < total {
				nextPage := strconv.Itoa(betsapiResponse.Pager.Page + 1)
				nextPageEvents, err := betsapi.GetUpcomingEvents(sportId, leagueId, teamId, countryCode, day, nextPage)
				if err != nil {
					return upcomingEvents, err
				}

				upcomingEvents = append(upcomingEvents, nextPageEvents...)
			}

			return upcomingEvents, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /events/upcoming", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /events/upcoming", resp.StatusCode)
	}
}

//Check events that not started or where max 15m lasted
func (betsapi *BetsapiCrawler) GetStartingEvents(sportId string, minuteThreshold int64) []types.Event {
	log.Info("Getting starting events...")
	var startingEvents []types.Event
	var startingEventBySportId []types.Event

	//check starting events in cache
	if val, exist := betsapi.Cache.Load(fmt.Sprintf("starting_%d", minuteThreshold)); exist {
		log.Debugln("Starting: Returning value from cache")
		allStartingEvents := val.([]types.Event)
		for _, event := range allStartingEvents {
			if event.SportId == sportId {
				startingEventBySportId = append(startingEventBySportId, event)
			}
		}

		return startingEventBySportId
	}

	inPlayEvents := betsapi.GetInPlayEvents(sportId)
	for _, event := range inPlayEvents {
		minutes, _ := event.Timer.Minutes.Int64()
		if event.TimeStatus == "0" || minutes < minuteThreshold {
			startingEvents = append(startingEvents, event)
			if event.SportId == sportId {
				startingEventBySportId = append(startingEventBySportId, event)
			}
		}
	}

	//store to cache
	log.Debugln("Starting: storing value to cache")
	betsapi.Cache.Store(fmt.Sprintf("starting_%d", minuteThreshold), startingEvents)

	return startingEvents
}

func (betsapi *BetsapiCrawler) GetEventView(eventId string) interface{} {
	req, err := http.NewRequest("GET", EventViewUrl, nil)
	if err != nil {
		log.Error(err)
	}

	//encode query parameters
	q := req.URL.Query()
	q.Add("token", API_TOKEN)
	q.Add("event_id", eventId)
	req.URL.RawQuery = q.Encode()

	resp, err := betsapi.Client.Do(req)
	if err != nil {
		log.Error(err)
	}

	var betsapiResponse types.BetsapiStatsResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			log.Error(err)
			return nil
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			log.Error(err)
			return nil
		}

		if betsapiResponse.Success == 1 {
			if len(betsapiResponse.Results) > 0 {
				event := betsapiResponse.Results[0]

				switch event.SportId {
				case types.SoccerId:
					var footballEventView types.BetsapiFootballStatsResponse
					err = json.Unmarshal(data, &footballEventView)
					if err != nil {
						log.Error(err)
						return nil
					}

					if len(footballEventView.Results) > 0 {
						return footballEventView.Results[0]
					}
				case types.BasketballId:
					return nil
				case types.TennisId:
					return nil
				default:
					log.Error("Unsupported sport id for event view")
				}
			}

		} else {
			log.Errorf("Error: %d: unsuccessful API response: /events/view", resp.StatusCode)
		}
	} else {
		log.Errorf("Error: %d: request: /event/view", resp.StatusCode)
	}

	return nil
}

func (betsapi *BetsapiCrawler) GetEventHistory(eventId string, qty string) (*types.EventHistory, error) {
	req, err := http.NewRequest("GET", EventHistoryUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := req.URL.Query()
	q.Add("token", API_TOKEN)
	q.Add("event_id", eventId)
	q.Add("qty", qty)

	req.URL.RawQuery = q.Encode()
	resp, err := betsapi.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var betsapiResponse types.BetsapiEventHistoryResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			return nil, err
		}

		if betsapiResponse.Success == 1 {
			return &betsapiResponse.Results, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /event/history", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /event/history", resp.StatusCode)
	}
}

func (betsapi *BetsapiCrawler) GetEventOdds(eventId string) (*types.Odds, error) {
	req, err := http.NewRequest("GET", EventOddsUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := req.URL.Query()
	q.Add("token", API_TOKEN)
	q.Add("event_id", eventId)

	req.URL.RawQuery = q.Encode()
	resp, err := betsapi.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var betsapiResponse types.BetsapiEventOddsResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			return nil, err
		}

		if betsapiResponse.Success == 1 {
			return &betsapiResponse.Results.Odds, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /event/odds", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /event/odds", resp.StatusCode)
	}
}

func (betsapi *BetsapiCrawler) GetEventOddsSummary() {
	//Ignore, prematch odds by bookmakers
}

func (betsapi *BetsapiCrawler) GetEndedEvents(sportId, leagueId, teamId, countryCode, day, page string) []types.Event {
	var endedEvents []types.Event
	req, err := http.NewRequest("GET", EndedEventsUrl, nil)
	if err != nil {
		log.Error(err)
	}

	//encode query parameters
	q := req.URL.Query()
	q.Add("token", API_TOKEN)
	q.Add("sport_id", sportId)
	q.Add("league_id", leagueId)
	q.Add("team_id", teamId)
	q.Add("cc", countryCode)
	q.Add("day", day)
	q.Add("page", page)

	req.URL.RawQuery = q.Encode()
	resp, err := betsapi.Client.Do(req)
	if err != nil {
		log.Error(err)
	}

	var betsapiResponse types.BetsapEventsPagerResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			log.Error(err)
			return nil
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			log.Error(err)
			return nil
		}

		if betsapiResponse.Success == 1 {
			log.Info("Total events: ", betsapiResponse.Pager.Total)
			endedEvents = append(endedEvents, betsapiResponse.Results...)

			actualPage := betsapiResponse.Pager.Page
			perPage := betsapiResponse.Pager.PerPage
			total := betsapiResponse.Pager.Total

			if actualPage == 100 {
				log.Warn("Warning: max page limit", resp.StatusCode)
				return endedEvents
			}

			if actualPage*perPage < total {
				nextPage := strconv.Itoa(betsapiResponse.Pager.Page + 1)
				endedEvents = append(endedEvents, betsapi.GetEndedEvents(sportId, leagueId, teamId, countryCode, day, nextPage)...)
			}
		} else {
			log.Errorf("Error: %d: unsuccessful API response: /events/ended", resp.StatusCode)
		}
	} else {
		log.Errorf("Error: %d: request: /events/ended", resp.StatusCode)
	}

	return endedEvents
}

//soccer only
func (betsapi *BetsapiCrawler) GetEventStatsTrend(eventId string) *types.StatsTrend {
	req, err := http.NewRequest("GET", EventStatsTrendUrl, nil)
	if err != nil {
		log.Error(err)
	}

	//encode query parameters
	q := req.URL.Query()
	q.Add("token", API_TOKEN)
	q.Add("event_id", eventId)

	req.URL.RawQuery = q.Encode()
	resp, err := betsapi.Client.Do(req)
	if err != nil {
		log.Error(err)
	}

	var betsapiResponse types.BetsapiStatsTrendResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			log.Error(err)
			return nil
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			log.Error(err)
			return nil
		}

		if betsapiResponse.Success == 1 {
			return &betsapiResponse.Results
		} else {
			log.Errorf("Error: %d: unsuccessful API response: /event/stats_trend", resp.StatusCode)
		}
	} else {
		log.Errorf("Error: %d: request: /event/stats_trend", resp.StatusCode)
	}

	return nil
}

func (betsapi *BetsapiCrawler) GetEventLineup(eventId string) {

}

func (betsapi *BetsapiCrawler) GetEventVideos(eventId string) {

}

func (betsapi *BetsapiCrawler) GetLeagues(sportId string, countryCode string, page string) []types.FootballLeague {
	var leagues []types.FootballLeague

	req, err := http.NewRequest("GET", LeagueUrl, nil)
	if err != nil {
		log.Error(err)
	}

	//encode query parameters
	q := req.URL.Query()
	q.Add("cc", countryCode)
	q.Add("page", page)
	q.Add("token", API_TOKEN)
	q.Add("sport_id", sportId)

	req.URL.RawQuery = q.Encode()
	resp, err := betsapi.Client.Do(req)
	if err != nil {
		log.Error(err)
	}

	var betsapiResponse types.BetsapiLeagueResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			log.Error(err)
			return nil
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			log.Error(err)
			return nil
		}

		if betsapiResponse.Success == 1 {
			leagues = append(leagues, betsapiResponse.Results...)
			actualPage := betsapiResponse.Pager.Page
			perPage := betsapiResponse.Pager.PerPage
			total := betsapiResponse.Pager.Total

			if actualPage*perPage < total {
				nextPage := strconv.Itoa(betsapiResponse.Pager.Page + 1)
				leagues = append(leagues, betsapi.GetLeagues(sportId, countryCode, nextPage)...)
			}
		} else {
			log.Errorf("Error: %d: unsuccessful API response: /event/league", resp.StatusCode)
		}
	} else {
		log.Errorf("Error: %d: request: /event/league", resp.StatusCode)
	}

	return leagues
}

func (betsapi *BetsapiCrawler) GetLeagueTable(leagueId string) {

}

func (betsapi *BetsapiCrawler) GetLeagueTopList(leagueId string) {

}

func (betsapi *BetsapiCrawler) GetTeams(sportId string) {

}

func (betsapi *BetsapiCrawler) GetTeamSquad(teamId string) {

}

func (betsapi *BetsapiCrawler) GetPlayerInfo(playerId string) {

}

func (betsapi *BetsapiCrawler) GetTennisRankong(typeId string) {

}

func (betsapi *BetsapiCrawler) GetHotFootballEvents() {
	for _, event := range betsapi.FootballEvents {
		if isDangerousAttacks(event.FootballStatistics.DangerousAttacks) && isShots(event.OnTarget, event.OffTarget) {
			prettifyFootballEventOutput(event)
		}
	}
}

func prettifyFootballEventOutput(event types.FootballEvent) {
	fmt.Printf("League: %s\n", event.League.Name)
	fmt.Printf("Match: %s - %s\n", event.HomeTeam.Name, event.AwayTeam.Name)
	fmt.Printf("Score: %s\n", event.Score)
	fmt.Printf("========== Statistics =========\n")
	fmt.Printf("Attacks: %s - %s\n", event.Attacks[0], event.Attacks[1])
	fmt.Printf("Dangerous attacks: %s - %s\n", event.DangerousAttacks[0], event.DangerousAttacks[1])
	fmt.Printf("Shots on target: %s - %s\n", event.OnTarget[0], event.OnTarget[1])
	fmt.Printf("Shots off target: %s - %s\n", event.OffTarget[0], event.OffTarget[1])
	fmt.Printf("Corners: %s - %s\n", event.Corners[0], event.Corners[1])
	fmt.Printf("-------------------------------\n")
}

func isDangerousAttacks(dangerousAttacks [2]string) bool {
	homeDangAttacks, _ := strconv.Atoi(dangerousAttacks[0])
	awayDangAttacks, _ := strconv.Atoi(dangerousAttacks[1])

	if ((homeDangAttacks - awayDangAttacks) >= 7) || ((awayDangAttacks - homeDangAttacks) >= 7) {
		return true
	} else {
		return false
	}
}

func isShots(shotsOnTarget [2]string, shotsOffTarget [2]string) bool {
	homeOnTarget, _ := strconv.Atoi(shotsOnTarget[0])
	awayOnTarget, _ := strconv.Atoi(shotsOnTarget[1])

	homeOffTarget, _ := strconv.Atoi(shotsOffTarget[0])
	awayOffTarget, _ := strconv.Atoi(shotsOffTarget[1])

	if ((homeOnTarget - awayOnTarget) >= 2) || ((awayOnTarget - homeOnTarget) >= 2) ||
		(((homeOnTarget - awayOnTarget) >= 1) || ((awayOnTarget - homeOnTarget) >= 1) &&
			((homeOffTarget-awayOffTarget >= 3) || (awayOffTarget-homeOffTarget >= 3))) {
		return true
	} else {
		return false
	}
}
