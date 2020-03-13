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
