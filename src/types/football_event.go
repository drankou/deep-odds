package types

//aggregated data about football event
type FootballEvent struct {
	Event      *Event        `json:"event_info" bson:"event_info"`
	History    *EventHistory `json:"history" bson:"event_history"`
	Odds       *Odds         `json:"odds" bson:"odds"`
	StatsTrend *StatsTrend   `json:"stats_trend" bson:"stats_trend"`
}

func (f *FootballEvent) Clean() {
	f.Event.Clean()
	f.History.Clean()
	f.Odds.Clean()
	f.StatsTrend.Clean()
}
