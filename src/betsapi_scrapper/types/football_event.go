package types

//aggregated data about football event

type FootballEvent struct {
	Event      *Event        `json:"event,omitempty" bson:"event"`
	History    *EventHistory `json:"history,omitempty" bson:"history"`
	Odds       *Odds         `json:"odds,omitempty" bson:"odds"`
	StatsTrend *StatsTrend   `json:"stats_trend,omitempty" bson:"stats_trend"`
}

func (f *FootballEvent) Clean() {
	if f.Event != nil {
		f.Event.Clean()
	}

	if f.History != nil {
		f.History.Clean()
	}

	if f.Odds != nil {
		f.Odds.Clean()
	}

	//if f.StatsTrend != nil {
	//	f.StatsTrend.Clean()
	//}
}
