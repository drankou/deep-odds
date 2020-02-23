package types

import (
	"github.com/sirupsen/logrus"
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
		logrus.Error(err)
		return nil
	}

	value, err := strconv.ParseInt(stt.Value, 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil
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
