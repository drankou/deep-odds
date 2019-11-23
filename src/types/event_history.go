package types

type EventHistory struct {
	H2H  []Event `json:"h2h" bson:"h2h"`
	Home []Event `json:"home" bson:"home"`
	Away []Event `json:"away" bson:"away"`
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