package types

//aggregated data about football event

func (m *FootballEvent) Clean() {
	if m.GetEvent() != nil {
		m.GetEvent().Clean()
	}

	if m.GetHistory() != nil {
		m.GetHistory().Clean()
	}

	if m.GetOdds() != nil {
		m.GetOdds().Clean()
	}

	if m.GetStatsTrend() != nil {
		m.GetStatsTrend().Clean()
	}
}
