package types

import (
	"sort"
	"strconv"
)

type Odds struct {
	FullTimeResult         []*Result              `json:"1_1" bson:"full_time"`
	AsianHandicap          []*AsianHandicapResult `json:"1_2" bson:"asian_handicap"`
	GoalLineTotal          []*AsianHandicapTotal  `json:"1_3" bson:"total"`
	AsianCorners           []*AsianHandicapTotal  `json:"1_4" bson:"asian_corners"`
	FirstHalfAsianHandicap []*AsianHandicapResult `json:"1_5" bson:"first_half_asian_handicap"`
	FirstHalfGoalLineTotal []*AsianHandicapTotal  `json:"1_6" bson:"first_half_total"`
	FirstHalfAsianCorners  []*AsianHandicapTotal  `json:"1_7" bson:"first_half_asian_corners"`
	FirstHalfResult        []*Result              `json:"1_8" bson:"first_half"`
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
	Id          string `json:"id,omitempty" bson:"-"`
	*ResultOdds `bson:"odds"`
	*OddsInfo   `bson:"odds_info"`
}

type ResultOdds struct {
	HomeOdd float64 `json:"home_od,string" bson:"home_odds"`
	DrawOdd float64 `json:"draw_od,string" bson:"draw_odds"`
	AwayOdd float64 `json:"away_od,string" bson:"away_odds"`
}

func RemoveDuplicitResultOdds(resultOdds []*Result) []*Result {
	var resultList []*Result
	keys := make(map[int64]bool)
	for i, entry := range resultOdds {
		if _, value := keys[entry.Minute]; !value && (entry.HomeOdd != -1 || entry.AwayOdd != -1 || entry.DrawOdd != -1) {
			keys[entry.Minute] = true
			resultList = append(resultList, resultOdds[i])
		}
	}

	return resultList
}

func AddMissingResultOdds(resultOdds []*Result) []*Result {
	var resultList []*Result

	minuteOdds := make(map[int64]*Result)
	var minutes []int64
	for i, entry := range resultOdds {
		minute := entry.Minute
		minutes = append(minutes, minute)
		minuteOdds[minute] = resultOdds[i]
	}

	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	if len(minutes) == 0 {
		uknownOdds := &ResultOdds{
			HomeOdd: -1,
			DrawOdd: -1,
			AwayOdd: -1,
		}

		firstMinuteOdds := &Result{
			Id:         "-1",
			ResultOdds: uknownOdds,
			OddsInfo: &OddsInfo{
				Score:  "0-0",
				Minute: 0,
			},
		}

		minuteOdds[0] = firstMinuteOdds
		minutes = append(minutes, 0)
	}

	//check if first available minute is 0 - start of the match
	if len(minutes) > 0 && minutes[0] != 0 {
		// create first minute odds from the first available minute data
		nextMinuteWithOdds := minuteOdds[minutes[0]]
		var firstMinuteOdds *Result

		if nextMinuteWithOdds != nil {
			unixTime := nextMinuteWithOdds.AddTime
			unixTime -= minutes[0] * 60

			if nextMinuteWithOdds.Score == "0-0" {
				firstMinuteOdds = &Result{
					Id:         nextMinuteWithOdds.Id + "0",
					ResultOdds: nextMinuteWithOdds.ResultOdds,
					OddsInfo: &OddsInfo{
						Score:   nextMinuteWithOdds.Score,
						Minute:  0,
						AddTime: unixTime,
					},
				}
			} else {
				uknownOdds := &ResultOdds{
					HomeOdd: -1,
					DrawOdd: -1,
					AwayOdd: -1,
				}

				firstMinuteOdds = &Result{
					Id:         nextMinuteWithOdds.Id + "0",
					ResultOdds: uknownOdds,
					OddsInfo: &OddsInfo{
						Score:   "0-0",
						Minute:  0,
						AddTime: unixTime,
					},
				}
			}
		} else {
			uknownOdds := &ResultOdds{
				HomeOdd: -1,
				DrawOdd: -1,
				AwayOdd: -1,
			}

			firstMinuteOdds = &Result{
				Id:         "-1",
				ResultOdds: uknownOdds,
				OddsInfo: &OddsInfo{
					Score:  "0-0",
					Minute: 0,
				},
			}
		}

		minuteOdds[0] = firstMinuteOdds
		minutes = append(minutes, 0)
	}

	lastMinute := int64(90)
	//TODO get last minute of the match
	for i := int64(1); i <= lastMinute; i++ {
		//check if minute presented in the minute->odds map
		if minuteOdds[i] == nil {
			//create minute odds with data from previous minute
			previousMinute := i - 1
			previousMinuteOdds := minuteOdds[previousMinute]

			unixTime := previousMinuteOdds.AddTime
			unixTime += 60

			newMinuteOdds := &Result{
				Id:         previousMinuteOdds.Id + strconv.FormatInt(i, 10),
				ResultOdds: previousMinuteOdds.ResultOdds,
				OddsInfo: &OddsInfo{
					Score:   previousMinuteOdds.Score,
					Minute:  i,
					AddTime: unixTime,
				},
			}

			minuteOdds[i] = newMinuteOdds
			minutes = append(minutes, i)
		}
	}

	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	for _, minute := range minutes {
		resultList = append(resultList, minuteOdds[minute])
	}

	return resultList
}

type AsianHandicapResult struct {
	Id                       string `json:"id,omitempty" bson:"-"`
	*AsianHandicapResultOdds `bson:"odds"`
	*OddsInfo                `bson:"odds_info"`
}

type AsianHandicapResultOdds struct {
	HomeOdd  float64 `json:"home_od,string" bson:"home_odds"`
	Handicap string  `json:"handicap" bson:"handicap"`
	AwayOdd  float64 `json:"away_od,string" bson:"away_odds"`
}

func RemoveDuplicitAsianHandicapResult(resultOdds []*AsianHandicapResult) []*AsianHandicapResult {
	var resultList []*AsianHandicapResult
	keys := make(map[int64]bool)
	for i, entry := range resultOdds {
		if _, value := keys[entry.Minute]; !value && (entry.HomeOdd != -1 || entry.AwayOdd != -1) {
			keys[entry.Minute] = true
			resultList = append(resultList, resultOdds[i])
		}
	}

	return resultList
}

func AddMissingAsianResultOdds(asianResultOdds []*AsianHandicapResult) []*AsianHandicapResult {
	var resultList []*AsianHandicapResult

	minuteOdds := make(map[int64]*AsianHandicapResult)
	var minutes []int64
	for i, entry := range asianResultOdds {
		minute := entry.Minute
		minutes = append(minutes, minute)
		minuteOdds[minute] = asianResultOdds[i]
	}

	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	//check if first available minute is 0 - start of the match
	if len(minutes) > 0 && minutes[0] != 0 {
		// create first minute odds from the first available minute data
		nextMinuteWithOdds := minuteOdds[minutes[0]]
		unixTime := nextMinuteWithOdds.AddTime
		unixTime -= minutes[0] * 60

		var firstMinuteOdds *AsianHandicapResult
		if nextMinuteWithOdds.Score == "0-0" {
			firstMinuteOdds = &AsianHandicapResult{
				Id:                      nextMinuteWithOdds.Id + "0",
				AsianHandicapResultOdds: nextMinuteWithOdds.AsianHandicapResultOdds,
				OddsInfo: &OddsInfo{
					Score:   nextMinuteWithOdds.Score,
					Minute:  0,
					AddTime: unixTime,
				},
			}
		} else {
			uknownOdds := &AsianHandicapResultOdds{
				HomeOdd:  -1,
				Handicap: "-1",
				AwayOdd:  -1,
			}

			firstMinuteOdds = &AsianHandicapResult{
				Id:                      nextMinuteWithOdds.Id + "0",
				AsianHandicapResultOdds: uknownOdds,
				OddsInfo: &OddsInfo{
					Score:   "0-0",
					Minute:  0,
					AddTime: unixTime,
				},
			}
		}

		minuteOdds[0] = firstMinuteOdds
		minutes = append(minutes, 0)
	}

	lastMinute := int64(90)
	//TODO get last minute of the match
	for i := int64(1); i <= lastMinute; i++ {
		//check if minute presented in the minute->odds map
		if minuteOdds[i] == nil {
			//create minute odds with data from previous minute
			previousMinute := i - 1
			previousMinuteOdds := minuteOdds[previousMinute]

			unixTime := previousMinuteOdds.AddTime
			unixTime += 60

			newMinuteOdds := &AsianHandicapResult{
				Id:                      previousMinuteOdds.Id + strconv.FormatInt(i, 10),
				AsianHandicapResultOdds: previousMinuteOdds.AsianHandicapResultOdds,
				OddsInfo: &OddsInfo{
					Score:   previousMinuteOdds.Score,
					Minute:  i,
					AddTime: unixTime,
				},
			}

			minuteOdds[i] = newMinuteOdds
			minutes = append(minutes, i)
		}
	}

	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	for _, minute := range minutes {
		resultList = append(resultList, minuteOdds[minute])
	}

	return resultList
}

type AsianHandicapTotal struct {
	Id                      string `json:"id" bson:"-"`
	*AsianHandicapTotalOdds `bson:"odds"`
	*OddsInfo               `bson:"odds_info"`
}

type OddsInfo struct {
	Score   string `json:"ss" bson:"score"`
	Minute  int64  `json:"time_str,string" bson:"minute"`
	AddTime int64  `json:"add_time,string" bson:"add_time"`
}

type AsianHandicapTotalOdds struct {
	OverOdd  float64 `json:"over_od,string" bson:"over_odds"`
	Handicap string  `json:"handicap" bson:"handicap"`
	UnderOdd float64 `json:"under_od,string" bson:"under_odds"`
}

func RemoveDuplicitAsianHandicapTotal(resultOdds []*AsianHandicapTotal) []*AsianHandicapTotal {
	var resultList []*AsianHandicapTotal
	keys := make(map[int64]bool)
	for _, entry := range resultOdds {
		if _, value := keys[entry.Minute]; !value && (entry.OverOdd != -1 || entry.UnderOdd != -1) {
			keys[entry.Minute] = true
			resultList = append(resultList, entry)
		}
	}

	return resultList
}

func AddMissingAsianTotalOdds(asianTotalOdds []*AsianHandicapTotal) []*AsianHandicapTotal {
	var resultList []*AsianHandicapTotal

	minuteOdds := make(map[int64]*AsianHandicapTotal)
	var minutes []int64
	for i, entry := range asianTotalOdds {
		minute := entry.Minute
		minutes = append(minutes, minute)
		minuteOdds[minute] = asianTotalOdds[i]
	}

	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	//check if first available minute is 0 - start of the match
	if len(minutes) > 0 && minutes[0] != 0 {
		// create first minute odds from the first available minute data
		nextMinuteWithOdds := minuteOdds[minutes[0]]
		unixTime := nextMinuteWithOdds.AddTime
		unixTime -= minutes[0] * 60

		var firstMinuteOdds *AsianHandicapTotal
		if nextMinuteWithOdds.Score == "0-0" {
			firstMinuteOdds = &AsianHandicapTotal{
				Id:                     nextMinuteWithOdds.Id + "0",
				AsianHandicapTotalOdds: nextMinuteWithOdds.AsianHandicapTotalOdds,
				OddsInfo: &OddsInfo{
					Score:   nextMinuteWithOdds.Score,
					Minute:  0,
					AddTime: unixTime,
				},
			}
		} else {
			uknownOdds := &AsianHandicapTotalOdds{
				OverOdd:  -1,
				Handicap: "-1",
				UnderOdd: -1,
			}

			firstMinuteOdds = &AsianHandicapTotal{
				Id:                     nextMinuteWithOdds.Id + "0",
				AsianHandicapTotalOdds: uknownOdds,
				OddsInfo: &OddsInfo{
					Score:   "0-0",
					Minute:  0,
					AddTime: unixTime,
				},
			}
		}

		minuteOdds[0] = firstMinuteOdds
		minutes = append(minutes, 0)
	}

	lastMinute := int64(90)
	//TODO get last minute of the match
	for i := int64(1); i <= lastMinute; i++ {
		//check if minute presented in the minute->odds map
		if minuteOdds[i] == nil {
			//create minute odds with data from previous minute
			previousMinute := i - 1
			previousMinuteOdds := minuteOdds[previousMinute]

			unixTime := previousMinuteOdds.AddTime
			unixTime += 60

			newMinuteOdds := &AsianHandicapTotal{
				Id:                     previousMinuteOdds.Id + strconv.FormatInt(i, 10),
				AsianHandicapTotalOdds: previousMinuteOdds.AsianHandicapTotalOdds,
				OddsInfo: &OddsInfo{
					Score:   previousMinuteOdds.Score,
					Minute:  i,
					AddTime: unixTime,
				},
			}

			minuteOdds[i] = newMinuteOdds
			minutes = append(minutes, i)
		}
	}

	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	for _, minute := range minutes {
		resultList = append(resultList, minuteOdds[minute])
	}

	return resultList
}
