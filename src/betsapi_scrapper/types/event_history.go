package types

func (m *EventHistory) Clean() {
	for _, event := range m.GetH2H() {
		event.Clean()
	}

	for _, event := range m.GetHome() {
		event.Clean()
	}

	for _, event := range m.GetAway() {
		event.Clean()
	}
}
