package types

// InPlay events response
type InplayEventsResponse struct {
	// Response success
	Success int         `json:"success"`
	Results []EventView `json:"results"`
}

// EventView response
type EventViewResponse struct {
	// Response success
	Success int         `json:"success"`
	Results []EventView `json:"results"`
}

type BetsapEventsPagerResponse struct {
	Success int         `json:"success"`
	Pager   Pager       `json:"pager"`
	Results []EventView `json:"results"`
}

// League response
type LeagueResponse struct {
	// Response success
	Success int `json:"success"`
	// Paginating response results
	Pager Pager `json:"pager"`
	// List of all leagues
	Results []League `json:"results"`
}

// Team response
type TeamResponse struct {
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

// ------------- STATS TREND ----------------------
type EventStatsTrendResponse struct {
	Success int             `json:"success"`
	Results EventStatsTrend `json:"results"`
}

//---------------- ODDS --------------------------
type EventOddsResponse struct {
	Success int `json:"success"`
	Results struct {
		Odds EventOdds `json:"odds"`
	} `json:"results"`
}

//-------------------- EVENT HISTORY---------------------------

type EventHistoryResponse struct {
	Success int          `json:"success"`
	Results EventHistory `json:"results"`
}
