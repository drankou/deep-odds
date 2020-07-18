package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/drankou/deep-odds/pkg/betsapi/types"
	"github.com/drankou/deep-odds/pkg/betsapi/types/constants"
	"github.com/drankou/deep-odds/pkg/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type BetsapiService struct {
	Client          *http.Client
	RateLimiter     *utils.RateLimiter
	Cache           *utils.Cache
	excludedLeagues map[string]bool
}

func (s *BetsapiService) Init() error {
	s.Client = &http.Client{}

	//initialize cache
	s.Cache = &utils.Cache{}
	err := s.Cache.Init()
	if err != nil {
		return err
	}

	// init rate limiter for api requests
	s.RateLimiter = &utils.RateLimiter{}
	s.RateLimiter.Init(time.Second)

	s.excludedLeagues = make(map[string]bool)
	for _, excludedLeague := range constants.ExcludedLeagues {
		s.excludedLeagues[excludedLeague] = true
	}

	log.Info("Betsapi service initialized")
	return nil
}

func (s *BetsapiService) GetInPlayEvents(ctx context.Context, req *types.InPlayEventsRequest) (*types.EventsResponse, error) {
	response := &types.EventsResponse{}
	log.Debug("Getting In-Play events...")
	//check in-play events in cache
	if val, exist := s.Cache.Load(fmt.Sprintf("inplay_%s", req.GetSportId())); exist {
		log.Debugln("InPlay: Returning value from cache")

		events := val.([]*types.EventView)
		response.Events = events

		return response, nil
	} else {
		defer func() {
			log.Debugln("InPlay: Storing value to cache")
			s.Cache.Store(fmt.Sprintf("inplay_%s", req.GetSportId()), response.Events, &utils.StoreOption{Ttl: time.Minute})
		}()
	}

	httpReq, err := http.NewRequest("GET", constants.InPlayEventsUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("sport_id", req.GetSportId())
	httpReq.URL.RawQuery = q.Encode()

	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var inplayEventsResponse types.InplayEventsResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &inplayEventsResponse)
		if err != nil {
			return nil, err
		}

		if inplayEventsResponse.Success == 1 {
			for i := range inplayEventsResponse.Results {
				if !s.excludedLeagues[inplayEventsResponse.Results[i].GetLeague().GetId()] {
					response.Events = append(response.GetEvents(), &inplayEventsResponse.Results[i])
				}
			}

			return response, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /events/inplay", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /events/inplay", resp.StatusCode)
	}
}

func (s *BetsapiService) GetUpcomingEvents(ctx context.Context, req *types.UpcomingEventsRequest) (*types.EventsResponse, error) {
	response := &types.EventsResponse{}

	httpReq, err := http.NewRequest("GET", constants.UpcomingEventsUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("sport_id", req.GetSportId())
	q.Add("league_id", req.GetLeagueId())
	q.Add("team_id", req.GetTeamId())
	q.Add("cc", req.GetCountryCode())
	q.Add("day", req.GetDay())
	q.Add("page", req.GetPage())

	httpReq.URL.RawQuery = q.Encode()
	s.RateLimiter.RateBlock()
	resp, err := s.Client.Do(httpReq)
	if err != nil {
		log.Error(err)
	}

	var eventsPagerResponse types.BetsapEventsPagerResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &eventsPagerResponse)
		if err != nil {
			return nil, err
		}

		if eventsPagerResponse.Success == 1 {
			if eventsPagerResponse.Pager.Page*eventsPagerResponse.Pager.PerPage < eventsPagerResponse.Pager.Total {
				response.NextPage = eventsPagerResponse.Pager.Page + 1
			} else {
				response.NextPage = -1
			}

			var events []*types.EventView
			for i := range eventsPagerResponse.Results {
				if !s.excludedLeagues[eventsPagerResponse.Results[i].GetLeague().GetId()] {
					events = append(events, &eventsPagerResponse.Results[i])
				}
			}

			response.Events = events

			return response, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /events/upcoming", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /events/upcoming", resp.StatusCode)
	}
}

func (BetsapiService) GetStartingEvents(ctx context.Context, req *types.StartingEventsRequest) (*types.EventsResponse, error) {
	panic("implement me")
}

func (s *BetsapiService) GetEndedEvents(ctx context.Context, req *types.EndedEventsRequest) (*types.EventsResponse, error) {
	response := &types.EventsResponse{}

	httpReq, err := http.NewRequest("GET", constants.EndedEventsUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("sport_id", req.GetSportId())
	q.Add("league_id", req.GetLeagueId())
	q.Add("team_id", req.GetTeamId())
	q.Add("cc", req.GetCountryCode())
	q.Add("day", req.GetDay())
	q.Add("page", req.GetPage())

	httpReq.URL.RawQuery = q.Encode()

	s.RateLimiter.RateBlock()
	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var eventsPagerResponse types.BetsapEventsPagerResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &eventsPagerResponse)
		if err != nil {
			return nil, err
		}

		if eventsPagerResponse.Success == 1 {
			if eventsPagerResponse.Pager.Page*eventsPagerResponse.Pager.PerPage < eventsPagerResponse.Pager.Total {
				response.NextPage = eventsPagerResponse.Pager.Page + 1
			} else {
				response.NextPage = -1
			}

			var events []*types.EventView
			for i := range eventsPagerResponse.Results {
				if !s.excludedLeagues[eventsPagerResponse.Results[i].GetLeague().GetId()] && eventsPagerResponse.Results[i].GetTimeStatus() == "3"{
					events = append(events, &eventsPagerResponse.Results[i])
				}
			}

			response.Events = events
			return response, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /events/ended", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /events/ended", resp.StatusCode)
	}
}

func (s *BetsapiService) GetEventView(ctx context.Context, req *types.EventViewRequest) (*types.EventView, error) {
	httpReq, err := http.NewRequest("GET", constants.EventViewUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("event_id", req.GetEventId())
	httpReq.URL.RawQuery = q.Encode()

	s.RateLimiter.RateBlock()
	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var eventViewResponse types.EventViewResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &eventViewResponse)
		if err != nil {
			return nil, err
		}

		if eventViewResponse.Success == 1 {
			if len(eventViewResponse.Results) > 0 {
				return &eventViewResponse.Results[0], nil
			}

		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /events/view", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /event/view", resp.StatusCode)
	}

	return nil, errors.Errorf("Error: %d: request: /event/view", resp.StatusCode)
}

func (s *BetsapiService) GetEventHistory(ctx context.Context, req *types.EventHistoryRequest) (*types.EventHistory, error) {
	httpReq, err := http.NewRequest("GET", constants.EventHistoryUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("event_id", req.GetEventId())
	if req.GetQty() != "" {
		q.Add("qty", req.GetQty())
	}

	httpReq.URL.RawQuery = q.Encode()

	s.RateLimiter.RateBlock()
	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var eventHistoryResponse types.EventHistoryResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &eventHistoryResponse)
		if err != nil {
			return nil, err
		}

		if eventHistoryResponse.Success == 1 {
			return &eventHistoryResponse.Results, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /event/history", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /event/history", resp.StatusCode)
	}
}

func (s *BetsapiService) GetEventOdds(ctx context.Context, req *types.EventOddsRequest) (*types.EventOdds, error) {
	httpReq, err := http.NewRequest("GET", constants.EventOddsUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("event_id", req.GetEventId())
	if req.GetSource() != "" {
		q.Add("source", req.GetSource())
	}
	if req.GetOddsMarket() != "" {
		q.Add("odds_market", req.GetOddsMarket())
	}
	if req.GetSinceTime() != 0 {
		q.Add("since_time", strconv.FormatInt(req.GetSinceTime(), 10))
	}

	httpReq.URL.RawQuery = q.Encode()

	s.RateLimiter.RateBlock()
	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var eventOddsResponse types.EventOddsResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		//replace "-" in odds to get Unmarshaling compatibility
		data = []byte(strings.Replace(string(data), `"-"`, `"-1"`, -1))

		err = json.Unmarshal(data, &eventOddsResponse)
		if err != nil {
			return nil, err
		}

		if eventOddsResponse.Success == 1 {
			return &eventOddsResponse.Results.Odds, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /event/odds", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /event/odds", resp.StatusCode)
	}
}

func (s *BetsapiService) GetEventStatsTrend(ctx context.Context, req *types.EventStatsTrendRequest) (*types.EventStatsTrend, error) {
	httpReq, err := http.NewRequest("GET", constants.EventStatsTrendUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("event_id", req.GetEventId())

	httpReq.URL.RawQuery = q.Encode()

	s.RateLimiter.RateBlock()
	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var statsTrendResponse types.EventStatsTrendResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &statsTrendResponse)
		if err != nil {
			return nil, err
		}

		if statsTrendResponse.Success == 1 {
			return &statsTrendResponse.Results, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /event/stats_trend", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /event/stats_trend", resp.StatusCode)
	}
}

func (s *BetsapiService) GetLeagues(ctx context.Context, req *types.LeaguesRequest) (*types.LeaguesResponse, error) {
	response := &types.LeaguesResponse{}

	httpReq, err := http.NewRequest("GET", constants.LeagueUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("cc", req.GetCountryCode())
	q.Add("page", req.GetPage())
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("sport_id", req.GetSportId())

	httpReq.URL.RawQuery = q.Encode()

	s.RateLimiter.RateBlock()
	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var leagueResponse types.LeagueResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &leagueResponse)
		if err != nil {
			return nil, err
		}

		if leagueResponse.Success == 1 {
			if leagueResponse.Pager.Page*leagueResponse.Pager.PerPage < leagueResponse.Pager.Total {
				response.NextPage = leagueResponse.Pager.Page + 1
			} else {
				response.NextPage = -1
			}

			var leagues []*types.League
			for i := range leagueResponse.Results {
				leagues = append(leagues, &leagueResponse.Results[i])
			}

			response.Leagues = leagues

			return response, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /event/league", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /event/league", resp.StatusCode)
	}
}

func (s *BetsapiService) GetTeams(ctx context.Context, req *types.TeamsRequest) (*types.TeamsResponse, error) {
	response := &types.TeamsResponse{}

	httpReq, err := http.NewRequest("GET", constants.TeamUrl, nil)
	if err != nil {
		return nil, err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("page", req.GetPage())
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("sport_id", req.GetSportId())

	httpReq.URL.RawQuery = q.Encode()

	s.RateLimiter.RateBlock()
	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var teamResponse types.TeamResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &teamResponse)
		if err != nil {
			return nil, err
		}

		if teamResponse.Success == 1 {
			if teamResponse.Pager.Page*teamResponse.Pager.PerPage < teamResponse.Pager.Total {
				response.NextPage = teamResponse.Pager.Page + 1
			} else {
				response.NextPage = -1
			}

			var teams []*types.Team
			for i := range teamResponse.Results {
				teams = append(teams, &teamResponse.Results[i])
			}

			response.Teams = teams

			return response, nil
		} else {
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /event/team", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /event/team", resp.StatusCode)
	}
}
