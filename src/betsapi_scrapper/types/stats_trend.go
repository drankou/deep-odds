package types

import (
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
)

type StatsTrend struct {
	Attacks          StatsTrendValue `json:"attacks" bson:"attacks"`
	DangerousAttacks StatsTrendValue `json:"dangerous_attacks" bson:"dangerous_attacks"`
	Possession       StatsTrendValue `json:"possesion" bson:"possesion"`
	OffTarget        StatsTrendValue `json:"off_target" bson:"off_target"`
	OnTarget         StatsTrendValue `json:"on_target" bson:"on_target"`
	Corners          StatsTrendValue `json:"corners" bson:"corners"`
	Goals            StatsTrendValue `json:"goals" bson:"goals"`
	YellowCards      StatsTrendValue `json:"yellowcards" bson:"yellow_cards"`
	RedCards         StatsTrendValue `json:"redcards" bson:"red_cards"`
	Substitutions    StatsTrendValue `json:"substitutions" bson:"substitutions"`
}

func (stats *StatsTrend) ToNew() *NewStatsTrend {
	if stats == nil {
		return nil
	}

	return &NewStatsTrend{
		Attacks:          stats.Attacks.ToNew(),
		DangerousAttacks: stats.DangerousAttacks.ToNew(),
		Possession:       stats.Possession.ToNew(),
		OffTarget:        stats.OffTarget.ToNew(),
		OnTarget:         stats.OnTarget.ToNew(),
		Corners:          stats.Corners.ToNew(),
		Goals:            stats.Goals.ToNew(),
		YellowCards:      stats.YellowCards.ToNew(),
		RedCards:         stats.RedCards.ToNew(),
		Substitutions:    stats.Substitutions.ToNew(),
	}
}

type NewStatsTrend struct {
	Attacks          *NewStatsTrendValue `json:"attacks" bson:"attacks"`
	DangerousAttacks *NewStatsTrendValue `json:"dangerous_attacks" bson:"dangerous_attacks"`
	Possession       *NewStatsTrendValue `json:"possesion" bson:"possesion"`
	OffTarget        *NewStatsTrendValue `json:"off_target" bson:"off_target"`
	OnTarget         *NewStatsTrendValue `json:"on_target" bson:"on_target"`
	Corners          *NewStatsTrendValue `json:"corners" bson:"corners"`
	Goals            *NewStatsTrendValue `json:"goals" bson:"goals"`
	YellowCards      *NewStatsTrendValue `json:"yellowcards" bson:"yellow_cards"`
	RedCards         *NewStatsTrendValue `json:"redcards" bson:"red_cards"`
	Substitutions    *NewStatsTrendValue `json:"substitutions" bson:"substitutions"`
}

func (stats *StatsTrend) Clean() {

}

type StatsTrendValue struct {
	Home []StatsTrendTick `json:"home"`
	Away []StatsTrendTick `json:"away"`
}

func (stv *StatsTrendValue) ToNew() *NewStatsTrendValue {
	return &NewStatsTrendValue{
		Home: StatsTrendTickSliceToPointers(stv.Home),
		Away: StatsTrendTickSliceToPointers(stv.Away),
	}
}

func StatsTrendTickSliceToPointers(ticks []StatsTrendTick) []*NewStatsTrendTick {
	var res []*NewStatsTrendTick
	for _, tick := range ticks {
		res = append(res, tick.ToNew())
	}

	return res
}

type NewStatsTrendValue struct {
	Home []*NewStatsTrendTick `json:"home"`
	Away []*NewStatsTrendTick `json:"away"`
}

type StatsTrendTick struct {
	Time  string `json:"time_str,omitempty"`
	Value string `json:"val,omitempty"`
}

func (stt *StatsTrendTick) ToNew() *NewStatsTrendTick {
	time, err := strconv.ParseInt(stt.Time, 10, 64)
	if err != nil {
		logrus.Error("stt.time", err)
	}

	value, err := strconv.ParseInt(stt.Value, 10, 64)
	if err != nil {
		logrus.Error("stt.value", err)
	}

	return &NewStatsTrendTick{
		Time:  time,
		Value: value,
	}
}

type NewStatsTrendTick struct {
	Time  int64 `json:"time_str,string" bson:"time"`
	Value int64 `json:"val,string" bson:"value"`
}

func AddMissingStatsTrend(statsTrend *NewStatsTrend) *NewStatsTrend {
	return &NewStatsTrend{
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

func addMissingStatsTrendValues(value *NewStatsTrendValue) *NewStatsTrendValue {
	return &NewStatsTrendValue{
		Home: addMissingStatsTrendTicks(value.Home),
		Away: addMissingStatsTrendTicks(value.Away),
	}
}

func addMissingStatsTrendTicks(ticks []*NewStatsTrendTick) []*NewStatsTrendTick {
	var res []*NewStatsTrendTick

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
		tick := &NewStatsTrendTick{
			Time:  minute,
			Value: minuteValue[minute],
		}

		res = append(res, tick)
	}

	return res
}
