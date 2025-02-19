package types

import (
	"sort"
	"strconv"
	"strings"
)

func (m *EventStatsTrend) Clean() {
//
}

func AddMissingStatsTrend(statsTrend *EventStatsTrend) *EventStatsTrend {
	return &EventStatsTrend{
		Attacks:          addMissingStatsTrendValues(statsTrend.GetAttacks()),
		DangerousAttacks: addMissingStatsTrendValues(statsTrend.GetDangerousAttacks()),
		Possession:       addMissingStatsTrendValues(statsTrend.GetPossession()),
		OffTarget:        addMissingStatsTrendValues(statsTrend.GetOffTarget()),
		OnTarget:         addMissingStatsTrendValues(statsTrend.GetOnTarget()),
		Corners:          addMissingStatsTrendValues(statsTrend.GetCorners()),
		Goals:            addMissingStatsTrendValues(statsTrend.GetGoals()),
		YellowCards:      addMissingStatsTrendValues(statsTrend.GetYellowCards()),
		RedCards:         addMissingStatsTrendValues(statsTrend.GetRedCards()),
		Substitutions:    addMissingStatsTrendValues(statsTrend.GetSubstitutions()),
	}
}

func addMissingStatsTrendValues(value *StatsTrendValue) *StatsTrendValue {
	if value == nil {
		value = &StatsTrendValue{}
	}

	return &StatsTrendValue{
		Home: addMissingStatsTrendTicks(value.Home),
		Away: addMissingStatsTrendTicks(value.Away),
	}
}

func addMissingStatsTrendTicks(ticks []*StatsTrendTick) []*StatsTrendTick {
	var res []*StatsTrendTick

	minuteValue := make(map[int64]int64)
	var minutes []int64
	for i := range ticks {
		minutes = append(minutes, ticks[i].GetTime())
		minuteValue[ticks[i].GetTime()] = ticks[i].GetValue()
	}
	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	//check edge cases when stats trend time series data may have a sequence such as
	// [44 45 46 45] where 46 represents stoppage time, hence 45+1, and second 45 is a start of the second half
	//so it could lead that value of 45th minute is larger than value of 46th
	for i := int64(45); i <= 55; i++ {
		for j := i - 1; j >= 45; j-- {
			if minuteValue[i] != 0 && minuteValue[j] != 0 {
				if minuteValue[j] > minuteValue[i] {
					minuteValue[i] = minuteValue[j]
				}
			}
		}
	}

	//check if first available minute is 0 - start of the match
	if (len(minutes) > 0 && minutes[0] != 0) || len(minutes) == 0 {
		// add first minute value
		minuteValue[0] = 0
		minutes = append(minutes, 0)
	}

	lastMinute := int64(90)
	for i := int64(1); i <= lastMinute; i++ {
		if _, ok := minuteValue[i]; !ok {
			//add missing minute value from the previous minute
			minuteValue[i] = minuteValue[i-1]
			minutes = append(minutes, i)
		}
	}

	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	for _, minute := range minutes {
		tick := &StatsTrendTick{
			Time:  minute,
			Value: minuteValue[minute],
		}

		res = append(res, tick)
	}

	return res
}

//Extracting stats about yellow cards from match' text events
func YellowCardsStatsFromEvents(footballEvent *FootballEvent) *StatsTrendValue {
	result := &StatsTrendValue{
		Home: []*StatsTrendTick{},
		Away: []*StatsTrendTick{},
	}

	for _, event := range footballEvent.Event.Events {
		minuteStr := strings.Split(event.Text, `'`)[0]
		var minute int64
		if strings.Contains(minuteStr, "+") {
			minute, _ = strconv.ParseInt(strings.Split(minuteStr, `+`)[0], 10, 64)
		} else {
			minute, _ = strconv.ParseInt(minuteStr, 10, 64)
		}

		if strings.Contains(event.Text, "Yellow Card") {
			newTick := &StatsTrendTick{
				Time: minute,
			}

			teamName := strings.Split(event.Text, `-`)[2]
			if strings.Contains(teamName, footballEvent.Event.HomeTeam.Name) {
				result.Home = append(result.Home, newTick)
			} else {
				result.Away = append(result.Away, newTick)
			}
		}
	}

	result.Home = sortAndFillValue(result.Home)
	result.Away = sortAndFillValue(result.Away)

	return result
}

func RedCardsStatsFromEvents(footballEvent *FootballEvent) *StatsTrendValue {
	result := &StatsTrendValue{
		Home: []*StatsTrendTick{},
		Away: []*StatsTrendTick{},
	}

	for _, event := range footballEvent.Event.Events {
		minuteStr := strings.Split(event.Text, `'`)[0]
		var minute int64
		if strings.Contains(minuteStr, "+") {
			minute, _ = strconv.ParseInt(strings.Split(minuteStr, `+`)[0], 10, 64)
		} else {
			minute, _ = strconv.ParseInt(minuteStr, 10, 64)
		}

		if strings.Contains(event.Text, "Red Card") {
			newTick := &StatsTrendTick{
				Time: minute,
			}

			teamName := strings.Split(event.Text, `-`)[2]
			if strings.Contains(teamName, footballEvent.Event.HomeTeam.Name) {
				result.Home = append(result.Home, newTick)
			} else {
				result.Away = append(result.Away, newTick)
			}
		}
	}

	result.Home = sortAndFillValue(result.Home)
	result.Away = sortAndFillValue(result.Away)

	return result
}

//Extracting stats about corners from match' text events
func CornersStatsFromEvents(footballEvent *FootballEvent) *StatsTrendValue {
	result := &StatsTrendValue{
		Home: []*StatsTrendTick{},
		Away: []*StatsTrendTick{},
	}

	for _, event := range footballEvent.Event.Events {
		minuteStr := strings.Split(event.Text, `'`)[0]
		var minute int64
		if strings.Contains(minuteStr, "+") {
			minute, _ = strconv.ParseInt(strings.Split(minuteStr, `+`)[0], 10, 64)
		} else {
			minute, _ = strconv.ParseInt(minuteStr, 10, 64)
		}

		if strings.Contains(event.Text, " Corner ") {
			newTick := &StatsTrendTick{
				Time: minute,
			}

			teamName := strings.Split(event.Text, `-`)[2]
			if strings.Contains(teamName, footballEvent.Event.HomeTeam.Name) {
				result.Home = append(result.Home, newTick)
			} else {
				result.Away = append(result.Away, newTick)
			}
		}
	}

	result.Home = sortAndFillValue(result.Home)
	result.Away = sortAndFillValue(result.Away)

	return result
}

func sortAndFillValue(ticks []*StatsTrendTick) []*StatsTrendTick {
	sort.Slice(ticks, func(i, j int) bool { return ticks[i].GetTime() < ticks[j].GetTime() })

	for i := range ticks {
		ticks[i].Value = int64(i + 1)
	}

	return ticks
}