package types

type Odds struct {
	FullTimeResult         []Result              `json:"1_1" bson:"full_time""`
	AsianHandicap          []AsianHandicapResult `json:"1_2" bson:"asian_handicap"`
	GoalLineTotal          []AsianHandicapTotal  `json:"1_3" bson:"total"`
	AsianCorners           []AsianHandicapTotal  `json:"1_4" bson:"asian_corners"`
	FirstHalfAsianHandicap []AsianHandicapResult `json:"1_5" bson:"first_half_asian_handicap"`
	FirstHalfGoalLineTotal []AsianHandicapTotal  `json:"1_6" bson:"first_half_total"`
	FirstHalfAsianCorners  []AsianHandicapTotal  `json:"1_7" bson:"first_half_asian_corners"`
	FirstHalfResult        []Result              `json:"1_8" bson:"first_half"`
}

func (odds *Odds) Clean() {
	//result
	odds.FullTimeResult = RemoveDuplicitResultOdds(odds.FullTimeResult)
	odds.FirstHalfResult = RemoveDuplicitResultOdds(odds.FirstHalfResult)

	//handicap
	odds.AsianHandicap = RemoveDuplicitAsianHandicapResult(odds.AsianHandicap)
	odds.FirstHalfAsianHandicap = RemoveDuplicitAsianHandicapResult(odds.FirstHalfAsianHandicap)

	//totals
	odds.GoalLineTotal = RemoveDuplicitAsianHandicapTotal(odds.GoalLineTotal)
	odds.AsianCorners = RemoveDuplicitAsianHandicapTotal(odds.GoalLineTotal)
	odds.FirstHalfGoalLineTotal = RemoveDuplicitAsianHandicapTotal(odds.GoalLineTotal)
	odds.FirstHalfAsianCorners = RemoveDuplicitAsianHandicapTotal(odds.GoalLineTotal)
}

type Result struct {
	Id string `json:"id" bson:"id"`
	ResultOdds
	Score   string `json:"ss" bson:"score"`
	Minute  string `json:"time_str" bson:"minute"`
	AddTime string `json:"add_time" bson:"add_time"`
}

type ResultOdds struct {
	HomeOdd string `json:"home_od" bson:"home_odds"`
	DrawOdd string `json:"draw_od" bson:"draw_odds"`
	AwayOdd string `json:"away_od" bson:"away_odds"`
}

func RemoveDuplicitResultOdds(resultOdds []Result) []Result {
	var resultList []Result
	keys := make(map[string]bool)
	for _, entry := range resultOdds {
		if _, value := keys[entry.Minute]; !value && entry.Minute != "" && entry.HomeOdd != "-" && entry.AwayOdd != "-" && entry.DrawOdd != "-" {
			keys[entry.Minute] = true
			resultList = append(resultList, entry)
		}
	}

	return resultList
}

//TODO missing minutes fixed with near values
func AddMissingResultOdds(resultOdds []Result) []Result {
	var resultList []Result
	minuteOdds := make(map[string]ResultOdds)
	for _, entry := range resultOdds {
		minuteOdds[entry.Minute] = entry.ResultOdds
	}

	return resultList
}

type AsianHandicapResult struct {
	Id string `json:"id" bson:"id"`
	AsianHandicapResultOdds
	Score   string `json:"ss" bson:"score"`
	Minute  string `json:"time_str" bson:"minute"`
	AddTime string `json:"add_time" bson:"add_time"`
}

type AsianHandicapResultOdds struct {
	HomeOdd  string `json:"home_od" bson:"home_odds"`
	Handicap string `json:"handicap" bson:"handicap"`
	AwayOdd  string `json:"away_od" bson:"away_odds"`
}

func RemoveDuplicitAsianHandicapResult(resultOdds []AsianHandicapResult) []AsianHandicapResult {
	var resultList []AsianHandicapResult
	keys := make(map[string]bool)
	for _, entry := range resultOdds {
		if _, value := keys[entry.Minute]; !value && entry.Minute != "" && entry.HomeOdd != "-" && entry.AwayOdd != "-" && entry.Handicap != "-" {
			keys[entry.Minute] = true
			resultList = append(resultList, entry)
		}
	}

	return resultList
}

type AsianHandicapTotal struct {
	Id string `json:"id" bson:"id"`
	AsianHandicapTotalOdds
	Score   string `json:"ss" bson:"score"`
	Minute  string `json:"time_str" bson:"minute"`
	AddTime string `json:"add_time" bson:"add_time"`
}

type AsianHandicapTotalOdds struct {
	OverOdd  string `json:"over_od" bson:"over_odds"`
	Handicap string `json:"handicap" bson:"handicap"`
	UnderOdd string `json:"under_od" bson:"under_odds"`
}

func RemoveDuplicitAsianHandicapTotal(resultOdds []AsianHandicapTotal) []AsianHandicapTotal {
	var resultList []AsianHandicapTotal
	keys := make(map[string]bool)
	for _, entry := range resultOdds {
		if _, value := keys[entry.Minute]; !value && entry.Minute != "" && entry.OverOdd != "-" && entry.UnderOdd != "-" && entry.Handicap != "-" {
			keys[entry.Minute] = true
			resultList = append(resultList, entry)
		}
	}

	return resultList
}
