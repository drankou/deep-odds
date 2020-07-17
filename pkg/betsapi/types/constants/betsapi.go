package constants

const API_URL string = "https://api.betsapi.com"

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

var ExcludedLeagues = []string{"22614", "22821", "22724", "22537", "22958", "22686", "22622", "22791", "22808", "23003",
	"22957", "22959", "23000", "22764", "22708", "23063", "23114", "23118", "23085", "22584", "22779", "22585", "22661",
	"22922", "22724", "22734", "22787", "22750", "23009", "22581", "12", "4742", "23246", "22605", "22687", "22631",
	"11375", "5384", "3122", "3260", "21760", "3262", "2650", "2655", "11197", "21760", "22652", "22656", "22648",
	"22611", "22598", "5284", "6179", "6203", "16333", "16376", "11895"}
