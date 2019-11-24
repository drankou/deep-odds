package types

import (
	"encoding/json"
)

type Event struct {
	Id         string           `json:"id" bson:"id"`
	Time       string           `json:"time" bson:"time"`
	SportId    string           `json:"sport_id" bson:"sport_id"`
	TimeStatus string           `json:"time_status" bson:"time_status"`
	Score      string           `json:"ss" bson:"score"`
	HomeTeam   Team             `json:"home" bson:"home_team"`
	AwayTeam   Team             `json:"away" bson:"away_team"`
	League     League           `json:"league" bson:"league"`
	Timer      Timer            `json:"timer" bson:"-"`
	ExtraInfo  ExtraInfo        `json:"extra" bson:"extra_info"`
	Events     []EventViewEvent `json:"events"`
	HasLineup  int              `json:"has_lineup" bson:"has_lineup"`
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
	Id          string      `json:"id" bson:"id"`
	Name        string      `json:"name" bson:"name"`
	ImageId     json.Number `json:"image_id" bson:"-"`
	CountryCode string      `json:"cc" bson:"country_code"`
}

type League struct {
	Id          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	CountryCode string `json:"cc" bson:"country_code"`
}

type Timer struct {
	Active    json.Number `json:"tt" bson:"-"`
	Minutes   json.Number `json:"tm" bson:"minutes"`
	Seconds   json.Number `json:"ts" bson:"seconds"`
	AddedTime int         `json:"ta" bson:"added_time"`
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
	HomePosition string  `json:"home_pos" bson:"home_position"`
	AwayPosition string  `json:"away_pos" bson:"away_position"`
}

type Manager struct {
	Name        string `json:"name" bson:"name"`
	CountryCode string `json:"cc" bson:"country_code"`
}

type Referee struct {
	Name        string `json:"name" bson:"name"`
	CountryCode string `json:"cc" bson:"country_code"`
}
