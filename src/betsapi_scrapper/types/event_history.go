package types

type EventHistory struct {
	H2H  []*Event `json:"h2h,omitempty" bson:"h2h"`
	Home []*Event `json:"home,omitempty" bson:"home"`
	Away []*Event `json:"away,omitempty" bson:"away"`
}

func (eventHistory *EventHistory) Clean() {
	//for _, event := range eventHistory.H2H {
	//	event.Clean()
	//}
	//
	//for _, event := range eventHistory.Home {
	//	event.Clean()
	//}
	//
	//for _, event := range eventHistory.Away {
	//	event.Clean()
	//}
}
