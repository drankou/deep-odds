package types

type Event struct {
	Id         string            `json:"id" bson:"id,omitempty"`
	Time       int64             `json:"time,string" bson:"time,omitempty"`
	SportId    string            `json:"sport_id" bson:"sport_id,omitempty"`
	TimeStatus string            `json:"time_status" bson:"time_status,omitempty"`
	Score      string            `json:"ss" bson:"score,omitempty"`
	HomeTeam   *Team             `json:"home" bson:"home_team,omitempty"`
	AwayTeam   *Team             `json:"away" bson:"away_team,omitempty"`
	League     *League           `json:"league" bson:"league,omitempty"`
	Timer      *Timer            `json:"timer" bson:"-,omitempty"`
	ExtraInfo  *ExtraInfo        `json:"extra,omitempty" bson:"extra_info,omitempty"`
	Events     []*EventViewEvent `json:"events,omitempty" bson:"events,omitempty"`
	HasLineup  int               `json:"has_lineup,omitempty" bson:"has_lineup,omitempty"`
}

func (event *Event) Clean() {
	//if event.HomeTeam.CountryCode == "" {
	//	event.HomeTeam.CountryCode = event.League.CountryCode
	//}
	//
	//if event.AwayTeam.CountryCode == "" {
	//	event.AwayTeam.CountryCode = event.League.CountryCode
	//}
}

type Team struct {
	Id   string `json:"id" bson:"id,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`
	//ImageId     json.Number `json:"image_id" bson:"-"`
	CountryCode string `json:"cc" bson:"country_code,omitempty"`
}

type League struct {
	Id          string `json:"id" bson:"id,omitempty"`
	Name        string `json:"name" bson:"name,omitempty"`
	CountryCode string `json:"cc" bson:"country_code,omitempty"`
}

type Timer struct {
	Active    int64 `json:"tt,string" bson:"-"`
	Minutes   int64 `json:"tm" bson:"minutes,omitempty"`
	Seconds   int64 `json:"ts" bson:"seconds,omitempty"`
	AddedTime int64 `json:"ta" bson:"added_time,omitempty"`
}

type ExtraInfo struct {
	HomeManager  *Manager     `json:"home_manager,omitempty" bson:"home_manager,omitempty"`
	AwayManager  *Manager     `json:"away_manager,omitempty" bson:"away_manager,omitempty"`
	Referee      *Referee     `json:"referee,omitempty" bson:"referee,omitempty"`
	StadiumData  *StadiumData `json:"stadium_data,omitempty" bson:"stadium_data,omitempty"`
	Length       int64        `json:"length,string,omitempty" bson:"length,omitempty"`
	Pitch        string       `json:"pitch,omitempty" bson:"pitch,omitempty"`
	Weather      string       `json:"weather,omitempty" bson:"weather,omitempty"`
	Stadium      string       `json:"stadium,omitempty" bson:"stadium,omitempty"`
	HomePosition int64        `json:"home_pos,string,omitempty" bson:"home_position,omitempty"`
	AwayPosition int64        `json:"away_pos,string,omitempty" bson:"away_position,omitempty"`
	Round        string       `json:"round,omitempty" bson:"round,omitempty"`
}

type Manager struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	CountryCode string `json:"cc,omitempty" bson:"country_code,omitempty"`
}

type Referee struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	CountryCode string `json:"cc,omitempty" bson:"country_code,omitempty"`
}

type StadiumData struct {
	Id           string `json:"id,omitempty" bson:"id,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	City         string `json:"city,omitempty" bson:"city,omitempty"`
	Country      string `json:"country,omitempty" bson:"country,omitempty"`
	Capacity     int64  `json:"capacity,string,omitempty" bson:"capacity,omitempty"`
	GoogleCoords string `json:"googlecoords,omitempty" bson:"googlecoords,omitempty"`
}

type EventViewEvent struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
