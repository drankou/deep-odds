package types

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
	Success int                    `json:"success"`
	Results []FootballEventResults `json:"results"`
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
	Id string `json:"id" bson:"id"`
	//League name
	Name string `json:"name" bson:"name"`
	//League country code
	Cc string `json:"cc" bson:"country_code"`
	// League has table flag
	HasLeagueTable int `json:"has_league_table" bson:"has_league_table"`
	// League has toplist flag
	HasLeagueToplist int `json:"has_league_toplist" bson:"has_league_toplist"`
}

//Football statistics
type FootballStatistics struct {
	Goals            [2]string `json:"goals" bson:"goals"`
	Attacks          [2]string `json:"attacks" bson:"attacks"`
	DangerousAttacks [2]string `json:"dangerous_attacks" bson:"dangerous_attacks"`
	Corners          [2]string `json:"corners" bson:"corners"`
	OnTarget         [2]string `json:"on_target" bson:"on_target"`
	OffTarget        [2]string `json:"off_target" bson:"off_target"`
}

type TennisStatistics struct {
}

type BasketballStatistics struct {
}

type FootballEventResults struct {
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

// ------------- STATS TREND ----------------------
type BetsapiStatsTrendResponse struct {
	Success int        `json:"success"`
	Results StatsTrend `json:"results"`
}

//---------------- ODDS --------------------------
type BetsapiEventOddsResponse struct {
	Success int `json:"success"`
	Results struct {
		Odds Odds `json:"odds"`
	} `json:"results"`
}

//-------------------- EVENT HISTORY---------------------------

type BetsapiEventHistoryResponse struct {
	Success int          `json:"success"`
	Results EventHistory `json:"results"`
}
