package types

import (
	"encoding/json"
)

//sportID description
const (
	SoccerId     string = "1"
	TennisId     string = "13"
	BasketballId string = "18"
	AllSports    string = "000"
)

//Time Status Error Value Description
const (
	NotStarted  string = "0"
	Inplay      string = "1"
	ToBeFixed   string = "2"
	Ended       string = "3"
	Postponed   string = "4"
	Canceled    string = "5"
	Walkover    string = "6"
	Interrupted string = "7"
	Abandoned   string = "8"
	Retired     string = "9"
	Removed     string = "10"
)

// InPlay events response
type BetsapiInplayResponse struct {
	// Response success
	Success int     `json:"success"`
	Results []Event `json:"results"`
}

// default stats response
type BetsapiStatsResponse struct {
	// Response success
	Success int     `json:"success"`
	Results []Event `json:"results"`
}

type BetsapEventsPagerResponse struct {
	Success int     `json:"success"`
	Pager   Pager   `json:"pager"`
	Results []Event `json:"results"`
}

// Stats response
type BetsapiFootballStatsResponse struct {
	// Response success
	Success int             `json:"success"`
	Results []FootballEvent `json:"results"`
}

// League response
type BetsapiLeagueResponse struct {
	// Response success
	Success int `json:"success"`
	// Paginating response results
	Pager Pager `json:"pager"`
	// List of all leagues
	Results []FootballLeague `json:"results"`
}

// Pagination structure
type Pager struct {
	// Current page
	Page int `json:"page"`
	// Number of results per page
	PerPage int `json:"per_page"`
	// Total number of results
	Total int `json:"total"`
}

// Football League
type FootballLeague struct {
	// League id
	Id string `json:"id"`
	//League name
	Name string `json:"name"`
	//League country code
	Cc string `json:"cc"`
	// League has table flag
	HasLeagueTable int `json:"has_league_table"`
	// League has toplist flag
	HasLeagueToplist int `json:"has_league_toplist"`
}

//Football statistics
type FootballStatistics struct {
	Goals            [2]string `json:"goals"`
	Attacks          [2]string `json:"attacks"`
	DangerousAttacks [2]string `json:"dangerous_attacks"`
	Corners          [2]string `json:"corners"`
	OnTarget         [2]string `json:"on_target"`
	OffTarget        [2]string `json:"off_target"`
}

type TennisStatistics struct {
}

type BasketballStatistics struct {
}

type Event struct {
	Id         string    `json:"id"`
	Time       string    `json:"time"`
	SportId    string    `json:"sport_id"`
	TimeStatus string    `json:"time_status"`
	Score      string    `json:"ss"`
	HomeTeam   Team      `json:"home"`
	AwayTeam   Team      `json:"away"`
	League     League    `json:"league"`
	Timer      Timer     `json:"timer"`
	ExtraInfo  ExtraInfo `json:"extra"`
	HasLineup  int       `json:"has_lineup"`
}

type FootballEvent struct {
	Event
	FootballStatistics `json:"stats"`
	Bet365Id           string `json:"bet365_id"`
}

type BasketballEvent struct {
	Event
	BasketballStatistics
}

type TennisEvent struct {
	Event
	TennisStatistics
}

type Team struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	ImageId     json.Number `json:"image_id"`
	CountryCode string      `json:"cc"`
}

type League struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Timer struct {
	Active    json.Number `json:"tt"`
	Minutes   json.Number `json:"tm"`
	Seconds   json.Number `json:"ts"`
	AddedTime int         `json:"ta"`
}

type ExtraInfo struct {
	Pitch        string `json:"pitch"`
	Weather      string `json:"weather"`
	Stadium      string `json:"stadium"`
	HomePosition string `json:"home_pos"`
	AwayPosition string `json:"away_pos"`
}

// ------------- STATS TREND ----------------------
type BetsapiStatsTrendResponse struct {
	Success int        `json:"success"`
	Results StatsTrend `json:"results"`
}

type StatsTrend struct {
	Attacks          StatsTrendValue `json:"attacks"`
	DangerousAttacks StatsTrendValue `json:"dangerous_attacks"`
	Possession       StatsTrendValue `json:"possesion"`
	OffTarget        StatsTrendValue `json:"off_target"`
	OnTarget         StatsTrendValue `json:"on_target"`
	Corners          StatsTrendValue `json:"corners"`
	Goals            StatsTrendValue `json:"goals"`
	YellowCards      StatsTrendValue `json:"yellow_cards"`
	Redcards         StatsTrendValue `json:"redcards"`
	Substitutions    StatsTrendValue `json:"substitutions"`
}

type StatsTrendValue struct {
	Home []StatsTrendTick `json:"home"`
	Away []StatsTrendTick `json:"away"`
}

type StatsTrendTick struct {
	Time  string `json:"time_str"`
	Value string `json:"val"`
}

//---------------- ODDS --------------------------
type BetsapiEventOddsResponse struct {
	Success int `json:"success"`
	Results struct {
		Odds Odds `json:"odds"`
	} `json:"results"`
}

type Odds struct {
	FullTimeResult         []Result              `json:"1_1"`
	AsianHandicap          []AsianHandicapResult `json:"1_2"`
	GoalLineTotal          []AsianHandicapTotal  `json:"1_3"`
	AsianCorners           []AsianHandicapTotal  `json:"1_4"`
	FirstHalfAsianHandicap []AsianHandicapResult `json:"1_5"`
	FirstHalfGoalLineTotal []AsianHandicapTotal  `json:"1_6"`
	FirstHalfAsianCorners  []AsianHandicapTotal  `json:"1_7"`
	FirstHalfResult        []Result              `json:"1_8"`
}

type Result struct {
	Id      string `json:"id"`
	HomeOdd string `json:"home_od"`
	DrawOdd string `json:"draw_od"`
	AwayOdd string `json:"away_od"`
	Score   string `json:"ss"`
	Minute  string `json:"time_str"`
	AddTime string `json:"add_time"`
}

type AsianHandicapResult struct {
	Id       string `json:"id"`
	HomeOdd  string `json:"home_od"`
	Handicap string `json:"handicap"`
	AwayOdd  string `json:"away_od"`
	Score    string `json:"ss"`
	Minute   string `json:"time_str"`
	AddTime  string `json:"add_time"`
}

type AsianHandicapTotal struct {
	Id       string `json:"id"`
	OverOdd  string `json:"over_od"`
	Handicap string `json:"handicap"`
	UnderOdd string `json:"under_od"`
	Score    string `json:"ss"`
	Minute   string `json:"time_str"`
	AddTime  string `json:"add_time"`
}

//-------------------- EVENT HISTORY---------------------------

type BetsapiEventHistoryResponse struct {
	Success int          `json:"success"`
	Results EventHistory `json:"results"`
}

type EventHistory struct {
	H2H  []Event `json:"h2h"`
	Home []Event `json:"home"`
	Away []Event `json:"away"`
}
