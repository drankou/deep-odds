package service

import (
	"betsapi_scrapper/types"
	"betsapi_scrapper/types/constants"
	"betsapi_scrapper/utils"
	"context"
	"encoding/json"
	"fmt"
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
	Client      *http.Client
	RateLimiter *utils.RateLimiter
	Cache       *types.Cache
}

func (s *BetsapiService) Init() error {
	s.Client = &http.Client{}

	//initialize cache
	s.Cache = &types.Cache{}
	s.Cache.ResetInterval = "@every 10m"
	s.Cache.ResetFunc = s.Cache.Clear
	err := s.Cache.Initialize()
	if err != nil {
		return err
	}

	// init rate limiter for api requests
	s.RateLimiter = &utils.RateLimiter{}
	s.RateLimiter.Init(time.Second)

	log.Info("Betsapi service initialized")
	return nil
}

func (s *BetsapiService) GetInPlayEvents(req *types.InPlayEventsRequest, stream types.Betsapi_GetInPlayEventsServer) error {
	log.Info("Getting In-Play events...")
	//check in-play events in cache
	if val, exist := s.Cache.Load(fmt.Sprintf("inplay_%s", req.GetSportId())); exist {
		events := val.([]types.Event)

		log.Debugln("InPlay: Returning value from cache")
		for _, event := range events {
			err := stream.Send(&event)
			if err != nil {
				return err
			}
		}
		return nil
	}

	httpReq, err := http.NewRequest("GET", constants.InPlayEventsUrl, nil)
	if err != nil {
		return err
	}

	//encode query parameters
	q := httpReq.URL.Query()
	q.Add("token", os.Getenv("BETSAPI_TOKEN"))
	q.Add("sport_id", req.GetSportId())
	httpReq.URL.RawQuery = q.Encode()

	resp, err := s.Client.Do(httpReq)
	if err != nil {
		return err
	}

	var betsapiResponse types.BetsapiInplayResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			return err
		}

		if betsapiResponse.Success == 1 {
			//save successfull response to cache
			log.Debugln("InPlay: Storing value to cache")
			s.Cache.Store(fmt.Sprintf("inplay_%s", req.GetSportId()), betsapiResponse.Results)

			for _, event := range betsapiResponse.Results {
				err := stream.Send(&event)
				if err != nil {
					return err
				}
			}

			return nil
		} else {
			return errors.Errorf("Error: %d: unsuccessful API response: /events/inplay", resp.StatusCode)
		}
	} else {
		return errors.Errorf("Error: %d: request: /events/inplay", resp.StatusCode)
	}
}

func (s *BetsapiService) GetUpcomingEvents(req *types.UpcomingEventsRequest, stream types.Betsapi_GetUpcomingEventsServer) error {
	httpReq, err := http.NewRequest("GET", constants.UpcomingEventsUrl, nil)
	if err != nil {
		return err
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

	var betsapiResponse types.BetsapEventsPagerResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			return err
		}

		if betsapiResponse.Success == 1 {
			for _, event := range betsapiResponse.Results {
				err := stream.Send(&event)
				if err != nil {
					return err
				}
			}

			return nil
		} else {
			return errors.Errorf("Error: %d: unsuccessful API response: /events/upcoming", resp.StatusCode)
		}
	} else {
		return errors.Errorf("Error: %d: request: /events/upcoming", resp.StatusCode)
	}
}

func (BetsapiService) GetStartingEvents(*types.StartingEventsRequest, types.Betsapi_GetStartingEventsServer) error {
	panic("implement me")
}

func (s *BetsapiService) GetEndedEvents(req *types.EndedEventsRequest, stream types.Betsapi_GetEndedEventsServer) error {
	httpReq, err := http.NewRequest("GET", constants.EndedEventsUrl, nil)
	if err != nil {
		return err
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
		return err
	}

	var betsapiResponse types.BetsapEventsPagerResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			return err
		}

		if betsapiResponse.Success == 1 {
			for _, event := range betsapiResponse.Results {
				err := stream.Send(&event)
				if err != nil {
					return err
				}
			}

			return nil
		} else {
			return errors.Errorf("Error: %d: unsuccessful API response: /events/ended", resp.StatusCode)
		}
	} else {
		return errors.Errorf("Error: %d: request: /events/ended", resp.StatusCode)
	}
}

func (s *BetsapiService) GetEventView(ctx context.Context, req *types.EventViewRequest) (*types.Event, error) {
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

	var betsapiResponse types.BetsapiStatsResponse
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
			if len(betsapiResponse.Results) > 0 {
				event := betsapiResponse.Results[0]

				switch event.SportId {
				case constants.SoccerId:
					var footballEventView types.BetsapiFootballStatsResponse
					err = json.Unmarshal(data, &footballEventView)
					if err != nil {
						return nil, err
					}

					if len(footballEventView.Results) > 0 {
						return &footballEventView.Results[0].Event, nil
					}
				case constants.BasketballId:
					return nil, nil
				case constants.TennisId:
					return nil, nil
				default:
					return nil, errors.New("Unsupported sport id for event view")
				}
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

func (s *BetsapiService) GetEventOdds(ctx context.Context, req *types.EventOddsRequest) (*types.Odds, error) {
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

	var betsapiResponse types.BetsapiEventOddsResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		//replace "-" in odds to get Unmarshaling compatibility
		data = []byte(strings.Replace(string(data), `"-"`, `"-1"`, -1))

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

func (s *BetsapiService) GetEventStatsTrend(ctx context.Context, req *types.EventStatsTrendRequest) (*types.StatsTrend, error) {
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

	var betsapiResponse types.BetsapiStatsTrendResponse
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
			return nil, errors.Errorf("Error: %d: unsuccessful API response: /event/stats_trend", resp.StatusCode)
		}
	} else {
		return nil, errors.Errorf("Error: %d: request: /event/stats_trend", resp.StatusCode)
	}
}

func (s *BetsapiService) GetLeagues(req *types.LeaguesRequest, stream types.Betsapi_GetLeaguesServer) error {
	httpReq, err := http.NewRequest("GET", constants.LeagueUrl, nil)
	if err != nil {
		return err
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
		return err
	}

	var betsapiResponse types.BetsapiLeagueResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			return err
		}

		if betsapiResponse.Success == 1 {
			for _, league := range betsapiResponse.Results {
				err := stream.Send(&league)
				if err != nil {
					return err
				}
			}

			return nil
		} else {
			return errors.Errorf("Error: %d: unsuccessful API response: /event/league", resp.StatusCode)
		}
	} else {
		return errors.Errorf("Error: %d: request: /event/league", resp.StatusCode)
	}
}

func (s *BetsapiService) GetTeams(req *types.TeamsRequest, stream types.Betsapi_GetTeamsServer) error {
	httpReq, err := http.NewRequest("GET", constants.TeamUrl, nil)
	if err != nil {
		return err
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
		return err
	}

	var betsapiResponse types.BetsapiTeamResponse
	if resp.StatusCode == 200 {
		body := resp.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &betsapiResponse)
		if err != nil {
			return err
		}

		if betsapiResponse.Success == 1 {
			for _, team := range betsapiResponse.Results {
				err := stream.Send(&team)
				if err != nil {
					return err
				}
			}

			return nil
		} else {
			return errors.Errorf("Error: %d: unsuccessful API response: /event/team", resp.StatusCode)
		}
	} else {
		return errors.Errorf("Error: %d: request: /event/team", resp.StatusCode)
	}
}
