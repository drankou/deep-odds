package types

type StatsTrend struct {
	Attacks          StatsTrendValue `json:"attacks" bson:"attacks"`
	DangerousAttacks StatsTrendValue `json:"dangerous_attacks" bson:"dangerous_attacks"`
	Possession       StatsTrendValue `json:"possesion" bson:"possesion"`
	OffTarget        StatsTrendValue `json:"off_target" bson:"off_target"`
	OnTarget         StatsTrendValue `json:"on_target" bson:"on_target"`
	Corners          StatsTrendValue `json:"corners" bson:"corners"`
	Goals            StatsTrendValue `json:"goals" bson:"goals"`
	YellowCards      StatsTrendValue `json:"yellow_cards" bson:"yellow_cards"`
	RedCards         StatsTrendValue `json:"redcards" bson:"red_cards"`
	Substitutions    StatsTrendValue `json:"substitutions" bson:"substitutions"`
}

func (stats *StatsTrend) Clean() {

}

type StatsTrendValue struct {
	Home []StatsTrendTick `json:"home"`
	Away []StatsTrendTick `json:"away"`
}

type StatsTrendTick struct {
	Time  string `json:"time_str"`
	Value string `json:"val"`
}
