package types

type EventHistory struct {
	H2H  []Event `json:"h2h,omitempty" bson:"h2h"`
	Home []Event `json:"home,omitempty" bson:"home"`
	Away []Event `json:"away,omitempty" bson:"away"`
}

func (eventHistory *EventHistory) ToNew() *NewEventHistory {
	if eventHistory == nil{
		return nil
	}

	var h2h []*NewEvent
	for i := range eventHistory.H2H {
		h2h = append(h2h, eventHistory.H2H[i].ToNew())
	}

	var home []*NewEvent
	for i := range eventHistory.Home {
		home = append(h2h, eventHistory.Home[i].ToNew())
	}

	var away []*NewEvent
	for i := range eventHistory.Away {
		away = append(h2h, eventHistory.Away[i].ToNew())
	}

	return &NewEventHistory{
		H2H:  h2h,
		Home: home,
		Away: away,
	}
}

type NewEventHistory struct {
	H2H  []*NewEvent `json:"h2h" bson:"h2h"`
	Home []*NewEvent `json:"home" bson:"home"`
	Away []*NewEvent `json:"away" bson:"away"`
}

func (eventHistory *EventHistory) Clean() {
	for _, event := range eventHistory.H2H {
		event.Clean()
	}

	for _, event := range eventHistory.Home {
		event.Clean()
	}

	for _, event := range eventHistory.Away {
		event.Clean()
	}
}
