package types

func (event *Event) Clean() {
	if event.HomeTeam.CountryCode == "" {
		event.HomeTeam.CountryCode = event.League.CountryCode
	}

	if event.AwayTeam.CountryCode == "" {
		event.AwayTeam.CountryCode = event.League.CountryCode
	}
}
