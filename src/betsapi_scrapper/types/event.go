package types

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

type Event struct {
	Id         string           `json:"id" bson:"id"`
	Time       string           `json:"time" bson:"time"`
	SportId    string           `json:"sport_id" bson:"sport_id"`
	TimeStatus string           `json:"time_status" bson:"time_status"`
	Score      string           `json:"ss" bson:"score"`
	HomeTeam   Team             `json:"home,omitempty" bson:"home_team"`
	AwayTeam   Team             `json:"away,omitempty" bson:"away_team"`
	League     League           `json:"league,omitempty" bson:"league"`
	Timer      Timer            `json:"timer,omitempty" bson:"-"`
	ExtraInfo  ExtraInfo        `json:"extra,omitempty" bson:"extra_info"`
	Events     []EventViewEvent `json:"events,omitempty"`
	HasLineup  int              `json:"has_lineup" bson:"has_lineup"`
}

func (event *Event) ToNew() *NewEvent {
	timestamp, err := strconv.ParseInt(event.Time, 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	var eventViewEvents []*EventViewEvent
	for _, eventView := range event.Events {
		eventViewEvents = append(eventViewEvents, &eventView)
	}

	return &NewEvent{
		Id:         event.Id,
		Time:       timestamp,
		SportId:    event.SportId,
		TimeStatus: event.TimeStatus,
		Score:      event.Score,
		HomeTeam:   &event.HomeTeam,
		AwayTeam:   &event.AwayTeam,
		League:     &event.League,
		Timer:      event.Timer.ToNew(),
		ExtraInfo:  event.ExtraInfo.ToNew(),
		Events:     eventViewEvents,
		HasLineup:  event.HasLineup,
	}
}

type NewEvent struct {
	Id         string            `json:"id" bson:"id"`
	Time       int64             `json:"time,string" bson:"time"`
	SportId    string            `json:"sport_id" bson:"sport_id"`
	TimeStatus string            `json:"time_status" bson:"time_status"`
	Score      string            `json:"ss" bson:"score"`
	HomeTeam   *Team             `json:"home" bson:"home_team"`
	AwayTeam   *Team             `json:"away" bson:"away_team"`
	League     *League           `json:"league" bson:"league"`
	Timer      *NewTimer         `json:"timer" bson:"-"`
	ExtraInfo  *NewExtraInfo     `json:"extra,omitempty" bson:"extra_info"`
	Events     []*EventViewEvent `json:"events,omitempty"`
	HasLineup  int               `json:"has_lineup,omitempty" bson:"has_lineup"`
}

func (event *Event) Clean() {
	if event.HomeTeam.CountryCode == "" {
		event.HomeTeam.CountryCode = event.League.CountryCode
	}

	if event.AwayTeam.CountryCode == "" {
		event.AwayTeam.CountryCode = event.League.CountryCode
	}
}

type Team struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	//ImageId     json.Number `json:"image_id" bson:"-"`
	CountryCode string `json:"cc" bson:"country_code"`
}

type League struct {
	Id          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	CountryCode string `json:"cc" bson:"country_code"`
}

type Timer struct {
	Active    string `json:"tt,omitempty" bson:"-"`
	Minutes   string `json:"tm,omitempty" bson:"minutes"`
	Seconds   string `json:"ts,omitempty" bson:"seconds"`
	AddedTime int    `json:"ta,omitempty" bson:"added_time"`
}

func (timer *Timer) ToNew() *NewTimer {
	active, err := strconv.ParseInt(timer.Active, 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	minutes, err := strconv.ParseInt(timer.Minutes, 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	seconds, err := strconv.ParseInt(timer.Seconds, 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	return &NewTimer{
		Active:    active,
		Minutes:   minutes,
		Seconds:   seconds,
		AddedTime: int64(timer.AddedTime),
	}
}

type NewTimer struct {
	Active    int64 `json:"tt,string" bson:"-"`
	Minutes   int64 `json:"tm,string" bson:"minutes"`
	Seconds   int64 `json:"ts,string" bson:"seconds"`
	AddedTime int64 `json:"ta" bson:"added_time"`
}

type EventViewEvent struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type ExtraInfo struct {
	HomeManager  Manager `json:"home_manager"`
	AwayManager  Manager `json:"away_manager"`
	Referee      Referee `json:"referee"`
	Pitch        string  `json:"pitch" bson:"pitch"`
	Weather      string  `json:"weather" bson:"weather"`
	Stadium      string  `json:"stadium" bson:"stadium"`
	HomePosition string  `json:"home_pos,omitempty" bson:"home_position"`
	AwayPosition string  `json:"away_pos,omitempty" bson:"away_position"`
}

func (info *ExtraInfo) ToNew() *NewExtraInfo {
	var homeManager *Manager
	if info.HomeManager.Name != "" {
		homeManager = &info.HomeManager
	}

	var awayManager *Manager
	if info.AwayManager.Name != "" {
		awayManager = &info.AwayManager
	}

	var referee *Referee
	if info.Referee.Name != "" {
		referee = &info.Referee
	}

	homePos, err := strconv.ParseInt(info.HomePosition, 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	awayPos, err := strconv.ParseInt(info.AwayPosition, 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	return &NewExtraInfo{
		HomeManager:  homeManager,
		AwayManager:  awayManager,
		Referee:      referee,
		Pitch:        info.Pitch,
		Weather:      info.Weather,
		Stadium:      info.Stadium,
		HomePosition: homePos,
		AwayPosition: awayPos,
	}
}

type NewExtraInfo struct {
	HomeManager  *Manager `json:"home_manager,omitempty" bson:"home_manager,omitempty"`
	AwayManager  *Manager `json:"away_manager,omitempty" bson:"away_manager,omitempty"`
	Referee      *Referee `json:"referee,omitempty" bson:"referee,omitempty"`
	Pitch        string   `json:"pitch,omitempty" bson:"pitch,omitempty"`
	Weather      string   `json:"weather,omitempty" bson:"weather,omitempty"`
	Stadium      string   `json:"stadium,omitempty" bson:"stadium,omitempty"`
	HomePosition int64    `json:"home_pos,string" bson:"home_position,omitempty"`
	AwayPosition int64    `json:"away_pos,string" bson:"away_position,omitempty"`
}

type Manager struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	CountryCode string `json:"cc,omitempty" bson:"country_code,omitempty"`
}

type Referee struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	CountryCode string `json:"cc,omitempty" bson:"country_code,omitempty"`
}
