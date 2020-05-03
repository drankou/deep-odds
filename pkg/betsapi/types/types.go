package types

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
	Results []League `json:"results"`
}

// Team response
type BetsapiTeamResponse struct {
	// Response success
	Success int `json:"success"`
	// Paginating response results
	Pager Pager `json:"pager"`
	// List of all leagues
	Results []Team `json:"results"`
}

// Pagination structure
type Pager struct {
	// Current page
	Page int32 `json:"page"`
	// Number of results per page
	PerPage int32 `json:"per_page"`
	// Total number of results
	Total int32 `json:"total"`
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
