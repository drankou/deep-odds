package types

import (
	"sort"
)

type StatsTrend struct {
	Attacks          *StatsTrendValue `json:"attacks" bson:"attacks"`
	DangerousAttacks *StatsTrendValue `json:"dangerous_attacks" bson:"dangerous_attacks"`
	Possession       *StatsTrendValue `json:"possesion" bson:"possesion"`
	OffTarget        *StatsTrendValue `json:"off_target" bson:"off_target"`
	OnTarget         *StatsTrendValue `json:"on_target" bson:"on_target"`
	Corners          *StatsTrendValue `json:"corners" bson:"corners"`
	Goals            *StatsTrendValue `json:"goals" bson:"goals"`
	YellowCards      *StatsTrendValue `json:"yellowcards" bson:"yellow_cards"`
	RedCards         *StatsTrendValue `json:"redcards" bson:"red_cards"`
	Substitutions    *StatsTrendValue `json:"substitutions" bson:"substitutions"`
}

func (stats *StatsTrend) Clean() {

}

type StatsTrendValue struct {
	Home []*StatsTrendTick `json:"home"`
	Away []*StatsTrendTick `json:"away"`
}

type StatsTrendTick struct {
	Time  int64 `json:"time_str,string" bson:"time"`
	Value int64 `json:"val,string" bson:"value"`
}

func AddMissingStatsTrend(statsTrend *StatsTrend) *StatsTrend {
	return &StatsTrend{
		Attacks:          addMissingStatsTrendValues(statsTrend.Attacks),
		DangerousAttacks: addMissingStatsTrendValues(statsTrend.DangerousAttacks),
		Possession:       addMissingStatsTrendValues(statsTrend.Possession),
		OffTarget:        addMissingStatsTrendValues(statsTrend.OffTarget),
		OnTarget:         addMissingStatsTrendValues(statsTrend.OnTarget),
		Corners:          addMissingStatsTrendValues(statsTrend.Corners),
		Goals:            addMissingStatsTrendValues(statsTrend.Goals),
		YellowCards:      addMissingStatsTrendValues(statsTrend.YellowCards),
		RedCards:         addMissingStatsTrendValues(statsTrend.RedCards),
		Substitutions:    addMissingStatsTrendValues(statsTrend.Substitutions),
	}
}

func addMissingStatsTrendValues(value *StatsTrendValue) *StatsTrendValue {
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
		minutes = append(minutes, ticks[i].Time)
		minuteValue[ticks[i].Time] = ticks[i].Value
	}
	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	//check if first available minute is 0 - start of the match
	if len(minutes) > 0 && minutes[0] != 0 {
		// add first minute value
		minuteValue[0] = 0
		minutes = append(minutes, 0)
	}

	lastMinute := int64(90)
	//TODO get last minute of the match
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
