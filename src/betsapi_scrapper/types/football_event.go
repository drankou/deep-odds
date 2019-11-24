package types

//aggregated data about football event
type FootballEvent struct {
	Event      *Event        `json:"event" bson:"event"`
	History    *EventHistory `json:"history" bson:"history"`
	Odds       *Odds         `json:"odds" bson:"odds"`
	StatsTrend *StatsTrend   `json:"stats_trend" bson:"stats_trend"`
}

func (f *FootballEvent) Clean() {
	f.Event.Clean()
	f.History.Clean()
	f.Odds.Clean()
	f.StatsTrend.Clean()
}
