package service

import (
	"betsapi_scrapper/types"
	"betsapi_scrapper/types/constants"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type BetsapiRateLimiter struct {
	*sync.Mutex
	rate time.Duration
	last time.Time
}

func (b *BetsapiRateLimiter) init() {
	b.Mutex = &sync.Mutex{}
	b.last = time.Now()
	b.rate = 1 * time.Second
}

func (b *BetsapiRateLimiter) rateBlock() {
	b.Lock()
	defer b.Unlock()

	if time.Since(b.last) < b.rate {
		<-time.After(b.last.Add(b.rate).Sub(time.Now()))
	}
	b.last = time.Now()
}

type BetsapiService struct {
	Client      *http.Client
	Context     context.Context
	RateLimiter *BetsapiRateLimiter
	SportId     string
	Cache       *types.Cache
}

func (s *BetsapiService) Init() error {
	s.Client = &http.Client{}
	s.Context = context.Background()

	//initialize cache
	s.Cache = &types.Cache{}
	s.Cache.ResetInterval = "every 1m"
	s.Cache.ResetFunc = s.Cache.Clear

	s.RateLimiter = &BetsapiRateLimiter{}
	s.RateLimiter.init()

	log.Info("Betsapi crawler initialized")
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
	var upcomingEvents []types.Event
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
	s.RateLimiter.rateBlock()
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
			//log.Info("Total events: ", betsapiResponse.Pager.Total)
			upcomingEvents = append(upcomingEvents, betsapiResponse.Results...)

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

func (BetsapiService) GetEndedEvents(*types.EndedEventsRequest, types.Betsapi_GetEndedEventsServer) error {
	panic("implement me")
}

func (BetsapiService) GetEventView(context.Context, *types.EventViewRequest) (*types.Event, error) {
	panic("implement me")
}

func (BetsapiService) GetEventHistory(context.Context, *types.EventHistoryRequest) (*types.EventHistory, error) {
	panic("implement me")
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

	httpReq.URL.RawQuery = q.Encode()

	s.RateLimiter.rateBlock()
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

	s.RateLimiter.rateBlock()
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

func (BetsapiService) GetLeagues(*types.LeaguesRequest, types.Betsapi_GetLeaguesServer) error {
	panic("implement me")
}
