package types

func (event *EventView) Clean() {
	if event.GetHomeTeam().GetCountryCode() == "" {
		event.GetHomeTeam().CountryCode = event.GetLeague().GetCountryCode()
	}

	if event.GetAwayTeam().GetCountryCode() == "" {
		event.GetAwayTeam().CountryCode = event.GetLeague().GetCountryCode()
	}
}
