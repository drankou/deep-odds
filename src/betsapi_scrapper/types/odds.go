package types

import (
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
)

type Odds struct {
	FullTimeResult         []Result              `json:"1_1" bson:"full_time"`
	AsianHandicap          []AsianHandicapResult `json:"1_2" bson:"asian_handicap"`
	GoalLineTotal          []AsianHandicapTotal  `json:"1_3" bson:"total"`
	AsianCorners           []AsianHandicapTotal  `json:"1_4" bson:"asian_corners"`
	FirstHalfAsianHandicap []AsianHandicapResult `json:"1_5" bson:"first_half_asian_handicap"`
	FirstHalfGoalLineTotal []AsianHandicapTotal  `json:"1_6" bson:"first_half_total"`
	FirstHalfAsianCorners  []AsianHandicapTotal  `json:"1_7" bson:"first_half_asian_corners"`
	FirstHalfResult        []Result              `json:"1_8" bson:"first_half"`
}

func (odds *Odds) ToNew() *NewOdds {
	if odds == nil {
		return nil
	}

	return &NewOdds{
		FullTimeResult:         ResultSliceToPointers(odds.FullTimeResult),
		AsianHandicap:          AsianHandicapResultSliceToPointers(odds.AsianHandicap),
		GoalLineTotal:          AsianHandicapTotalSliceToPointers(odds.GoalLineTotal),
		AsianCorners:           AsianHandicapTotalSliceToPointers(odds.AsianCorners),
		FirstHalfAsianHandicap: AsianHandicapResultSliceToPointers(odds.FirstHalfAsianHandicap),
		FirstHalfGoalLineTotal: AsianHandicapTotalSliceToPointers(odds.FirstHalfGoalLineTotal),
		FirstHalfAsianCorners:  AsianHandicapTotalSliceToPointers(odds.FirstHalfAsianCorners),
		FirstHalfResult:        ResultSliceToPointers(odds.FirstHalfResult),
	}
}

func ResultSliceToPointers(results []Result) []*NewResult {
	var res []*NewResult
	for _, r := range results {
		res = append(res, r.ToNew())
	}

	return res
}

func AsianHandicapResultSliceToPointers(asianResult []AsianHandicapResult) []*NewAsianHandicapResult {
	var res []*NewAsianHandicapResult
	for _, r := range asianResult {
		res = append(res, r.ToNew())
	}

	return res
}

func AsianHandicapTotalSliceToPointers(asianTotals []AsianHandicapTotal) []*NewAsianHandicapTotal {
	var res []*NewAsianHandicapTotal
	for _, r := range asianTotals {
		res = append(res, r.ToNew())
	}

	return res
}

type NewOdds struct {
	FullTimeResult         []*NewResult              `json:"1_1" bson:"full_time"`
	AsianHandicap          []*NewAsianHandicapResult `json:"1_2" bson:"asian_handicap"`
	GoalLineTotal          []*NewAsianHandicapTotal  `json:"1_3" bson:"total"`
	AsianCorners           []*NewAsianHandicapTotal  `json:"1_4" bson:"asian_corners"`
	FirstHalfAsianHandicap []*NewAsianHandicapResult `json:"1_5" bson:"first_half_asian_handicap"`
	FirstHalfGoalLineTotal []*NewAsianHandicapTotal  `json:"1_6" bson:"first_half_total"`
	FirstHalfAsianCorners  []*NewAsianHandicapTotal  `json:"1_7" bson:"first_half_asian_corners"`
	FirstHalfResult        []*NewResult              `json:"1_8" bson:"first_half"`
}

func (odds *NewOdds) Clean() {
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
	Minute  string `json:"time_str,omitempty" bson:"minute"`
	AddTime string `json:"add_time,omitempty" bson:"add_time"`
}

func (result *Result) ToNew() *NewResult {
	minute, err := strconv.ParseInt(result.Minute, 10, 64)
	if err != nil {
		logrus.Error("Result.minute:", err)
	}

	addTime, err := strconv.ParseInt(result.AddTime, 10, 64)
	if err != nil {
		logrus.Error("Result.addtime:", err)
	}

	return &NewResult{
		Id:            result.Id,
		NewResultOdds: result.ResultOdds.ToNew(),
		NewOddsInfo: &NewOddsInfo{
			Score:   result.Score,
			Minute:  minute,
			AddTime: addTime,
		},
	}
}

type NewResult struct {
	Id             string `json:"id,omitempty" bson:"-"`
	*NewResultOdds `bson:"odds"`
	*NewOddsInfo   `bson:"odds_info"`
}

type ResultOdds struct {
	HomeOdd string `json:"home_od,omitempty" bson:"home_odds"`
	DrawOdd string `json:"draw_od,omitempty" bson:"draw_odds"`
	AwayOdd string `json:"away_od,omitempty" bson:"away_odds"`
}

func (resultOdds *ResultOdds) ToNew() *NewResultOdds {
	homeOdds, err := strconv.ParseFloat(resultOdds.HomeOdd, 64)
	if err != nil {
		homeOdds = -1
	}

	awayOdds, err := strconv.ParseFloat(resultOdds.AwayOdd, 64)
	if err != nil {
		awayOdds = -1
	}

	drawOdds, err := strconv.ParseFloat(resultOdds.DrawOdd, 64)
	if err != nil {
		drawOdds = -1
	}

	return &NewResultOdds{
		HomeOdd: homeOdds,
		DrawOdd: drawOdds,
		AwayOdd: awayOdds,
	}
}

type NewResultOdds struct {
	HomeOdd float64 `json:"home_od,string" bson:"home_odds"`
	DrawOdd float64 `json:"draw_od,string" bson:"draw_odds"`
	AwayOdd float64 `json:"away_od,string" bson:"away_odds"`
}

func RemoveDuplicitResultOdds(resultOdds []*NewResult) []*NewResult {
	var resultList []*NewResult
	keys := make(map[int64]bool)
	for i, entry := range resultOdds {
		if _, value := keys[entry.Minute]; !value && (entry.HomeOdd != -1 || entry.AwayOdd != -1 || entry.DrawOdd != -1) {
			keys[entry.Minute] = true
			resultList = append(resultList, resultOdds[i])
		}
	}

	return resultList
}

func AddMissingResultOdds(resultOdds []*NewResult) []*NewResult {
	var resultList []*NewResult

	minuteOdds := make(map[int64]*NewResult)
	var minutes []int64
	for i, entry := range resultOdds {
		minute := entry.Minute
		minutes = append(minutes, minute)
		minuteOdds[minute] = resultOdds[i]
	}

	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })

	//check if first available minute is 0 - start of the match
	if len(minutes) > 0 && minutes[0] != 0 {
		// create first minute odds from the first available minute data
		nextMinuteWithOdds := minuteOdds[minutes[0]]
		unixTime := nextMinuteWithOdds.AddTime
		unixTime -= minutes[0] * 60

		var firstMinuteOdds *NewResult
		if nextMinuteWithOdds.Score == "0-0" {
			firstMinuteOdds = &NewResult{
				Id:            nextMinuteWithOdds.Id + "0",
				NewResultOdds: nextMinuteWithOdds.NewResultOdds,
				NewOddsInfo: &NewOddsInfo{
					Score:   nextMinuteWithOdds.Score,
					Minute:  0,
					AddTime: unixTime,
				},
			}
		} else {
			uknownOdds := &NewResultOdds{
				HomeOdd: -1,
				DrawOdd: -1,
				AwayOdd: -1,
			}

			firstMinuteOdds = &NewResult{
				Id:            nextMinuteWithOdds.Id + "0",
				NewResultOdds: uknownOdds,
				NewOddsInfo: &NewOddsInfo{
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

			newMinuteOdds := &NewResult{
				Id:            previousMinuteOdds.Id + strconv.FormatInt(i, 10),
				NewResultOdds: previousMinuteOdds.NewResultOdds,
				NewOddsInfo: &NewOddsInfo{
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
	Id string `json:"id" bson:"id"`
	AsianHandicapResultOdds
	Score   string `json:"ss" bson:"score"`
	Minute  string `json:"time_str,omitempty" bson:"minute"`
	AddTime string `json:"add_time,omitempty" bson:"add_time"`
}

func (asianRes *AsianHandicapResult) ToNew() *NewAsianHandicapResult {
	minute, err := strconv.ParseInt(asianRes.Minute, 10, 64)
	if err != nil {
		logrus.Error("asianResult.minute:", err)
	}

	addTime, err := strconv.ParseInt(asianRes.AddTime, 10, 64)
	if err != nil {
		logrus.Error("asianResult.addtime:", err)
	}

	return &NewAsianHandicapResult{
		Id:                         asianRes.Id,
		NewAsianHandicapResultOdds: asianRes.AsianHandicapResultOdds.ToNew(),
		NewOddsInfo: &NewOddsInfo{
			Score:   asianRes.Score,
			Minute:  minute,
			AddTime: addTime,
		},
	}
}

type NewAsianHandicapResult struct {
	Id                          string `json:"id,omitempty" bson:"-"`
	*NewAsianHandicapResultOdds `bson:"odds"`
	*NewOddsInfo                `bson:"odds_info"`
}

type AsianHandicapResultOdds struct {
	HomeOdd  string `json:"home_od,omitempty" bson:"home_odds"`
	Handicap string `json:"handicap,omitempty" bson:"handicap"`
	AwayOdd  string `json:"away_od,omitempty" bson:"away_odds"`
}

func (odds *AsianHandicapResultOdds) ToNew() *NewAsianHandicapResultOdds {
	homeOdds, err := strconv.ParseFloat(odds.HomeOdd, 64)
	if err != nil {
		homeOdds = -1
	}

	awayOdds, err := strconv.ParseFloat(odds.AwayOdd, 64)
	if err != nil {
		awayOdds = -1
	}

	return &NewAsianHandicapResultOdds{
		HomeOdd:  homeOdds,
		Handicap: odds.Handicap,
		AwayOdd:  awayOdds,
	}
}

type NewAsianHandicapResultOdds struct {
	HomeOdd  float64 `json:"home_od,string" bson:"home_odds"`
	Handicap string  `json:"handicap" bson:"handicap"`
	AwayOdd  float64 `json:"away_od,string" bson:"away_odds"`
}

func RemoveDuplicitAsianHandicapResult(resultOdds []*NewAsianHandicapResult) []*NewAsianHandicapResult {
	var resultList []*NewAsianHandicapResult
	keys := make(map[int64]bool)
	for i, entry := range resultOdds {
		if _, value := keys[entry.Minute]; !value && (entry.HomeOdd != -1 || entry.AwayOdd != -1) {
			keys[entry.Minute] = true
			resultList = append(resultList, resultOdds[i])
		}
	}

	return resultList
}

func AddMissingAsianResultOdds(asianResultOdds []*NewAsianHandicapResult) []*NewAsianHandicapResult {
	var resultList []*NewAsianHandicapResult

	minuteOdds := make(map[int64]*NewAsianHandicapResult)
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

		var firstMinuteOdds *NewAsianHandicapResult
		if nextMinuteWithOdds.Score == "0-0" {
			firstMinuteOdds = &NewAsianHandicapResult{
				Id:                         nextMinuteWithOdds.Id + "0",
				NewAsianHandicapResultOdds: nextMinuteWithOdds.NewAsianHandicapResultOdds,
				NewOddsInfo: &NewOddsInfo{
					Score:   nextMinuteWithOdds.Score,
					Minute:  0,
					AddTime: unixTime,
				},
			}
		} else {
			uknownOdds := &NewAsianHandicapResultOdds{
				HomeOdd:  -1,
				Handicap: "-1",
				AwayOdd:  -1,
			}

			firstMinuteOdds = &NewAsianHandicapResult{
				Id:                         nextMinuteWithOdds.Id + "0",
				NewAsianHandicapResultOdds: uknownOdds,
				NewOddsInfo: &NewOddsInfo{
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

			newMinuteOdds := &NewAsianHandicapResult{
				Id:            previousMinuteOdds.Id + strconv.FormatInt(i, 10),
				NewAsianHandicapResultOdds: previousMinuteOdds.NewAsianHandicapResultOdds,
				NewOddsInfo: &NewOddsInfo{
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
	Id string `json:"id" bson:"id"`
	AsianHandicapTotalOdds
	Score   string `json:"ss" bson:"score"`
	Minute  string `json:"time_str,omitempty" bson:"minute"`
	AddTime string `json:"add_time,omitempty" bson:"add_time"`
}

func (asianTotal *AsianHandicapTotal) ToNew() *NewAsianHandicapTotal {
	minute, err := strconv.ParseInt(asianTotal.Minute, 10, 64)
	if err != nil {
		logrus.Error("asianTotal.minute:", err)
	}

	addTime, err := strconv.ParseInt(asianTotal.AddTime, 10, 64)
	if err != nil {
		logrus.Error("asianTotal.addtime:", err)
	}

	return &NewAsianHandicapTotal{
		Id:                        asianTotal.Id,
		NewAsianHandicapTotalOdds: asianTotal.AsianHandicapTotalOdds.ToNew(),
		NewOddsInfo: &NewOddsInfo{
			Score:   asianTotal.Score,
			Minute:  minute,
			AddTime: addTime,
		},
	}
}

type NewAsianHandicapTotal struct {
	Id                         string `json:"id" bson:"-"`
	*NewAsianHandicapTotalOdds `bson:"odds"`
	*NewOddsInfo               `bson:"odds_info"`
}

type OddsInfo struct {
	Score   string `json:"ss" bson:"score"`
	Minute  string `json:"time_str,omitempty" bson:"minute"`
	AddTime string `json:"add_time,omitempty" bson:"add_time"`
}

func (oddsInfo *OddsInfo) ToNew() *NewOddsInfo {
	minute, err := strconv.ParseInt(oddsInfo.Minute, 10, 64)
	if err != nil {
		logrus.Error("oddsInfo.minute", err)
	}

	addTime, err := strconv.ParseInt(oddsInfo.AddTime, 10, 64)
	if err != nil {
		logrus.Error("oddsInfo.addtime", err)
	}

	return &NewOddsInfo{
		Score:   oddsInfo.Score,
		Minute:  minute,
		AddTime: addTime,
	}
}

type NewOddsInfo struct {
	Score   string `json:"ss" bson:"score"`
	Minute  int64  `json:"time_str,string" bson:"minute"`
	AddTime int64  `json:"add_time,string" bson:"add_time"`
}

type AsianHandicapTotalOdds struct {
	OverOdd  string `json:"over_od,omitempty" bson:"over_odds"`
	Handicap string `json:"handicap,omitempty" bson:"handicap"`
	UnderOdd string `json:"under_od,omitempty" bson:"under_odds"`
}

func (odds *AsianHandicapTotalOdds) ToNew() *NewAsianHandicapTotalOdds {
	homeOdds, err := strconv.ParseFloat(odds.OverOdd, 64)
	if err != nil {
		homeOdds = -1
	}

	awayOdds, err := strconv.ParseFloat(odds.UnderOdd, 64)
	if err != nil {
		homeOdds = -1
	}

	return &NewAsianHandicapTotalOdds{
		OverOdd:  homeOdds,
		Handicap: odds.Handicap,
		UnderOdd: awayOdds,
	}
}

type NewAsianHandicapTotalOdds struct {
	OverOdd  float64 `json:"over_od,string" bson:"over_odds"`
	Handicap string  `json:"handicap" bson:"handicap"`
	UnderOdd float64 `json:"under_od,string" bson:"under_odds"`
}

func RemoveDuplicitAsianHandicapTotal(resultOdds []*NewAsianHandicapTotal) []*NewAsianHandicapTotal {
	var resultList []*NewAsianHandicapTotal
	keys := make(map[int64]bool)
	for _, entry := range resultOdds {
		if _, value := keys[entry.Minute]; !value && (entry.OverOdd != -1 || entry.UnderOdd != -1) {
			keys[entry.Minute] = true
			resultList = append(resultList, entry)
		}
	}

	return resultList
}

func AddMissingAsianTotalOdds(asianTotalOdds []*NewAsianHandicapTotal) []*NewAsianHandicapTotal {
	var resultList []*NewAsianHandicapTotal

	minuteOdds := make(map[int64]*NewAsianHandicapTotal)
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

		var firstMinuteOdds *NewAsianHandicapTotal
		if nextMinuteWithOdds.Score == "0-0" {
			firstMinuteOdds = &NewAsianHandicapTotal{
				Id:                         nextMinuteWithOdds.Id + "0",
				NewAsianHandicapTotalOdds: nextMinuteWithOdds.NewAsianHandicapTotalOdds,
				NewOddsInfo: &NewOddsInfo{
					Score:   nextMinuteWithOdds.Score,
					Minute:  0,
					AddTime: unixTime,
				},
			}
		} else {
			uknownOdds := &NewAsianHandicapTotalOdds{
				OverOdd:  -1,
				Handicap: "-1",
				UnderOdd:  -1,
			}

			firstMinuteOdds = &NewAsianHandicapTotal{
				Id:                         nextMinuteWithOdds.Id + "0",
				NewAsianHandicapTotalOdds: uknownOdds,
				NewOddsInfo: &NewOddsInfo{
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

			newMinuteOdds := &NewAsianHandicapTotal{
				Id:            previousMinuteOdds.Id + strconv.FormatInt(i, 10),
				NewAsianHandicapTotalOdds: previousMinuteOdds.NewAsianHandicapTotalOdds,
				NewOddsInfo: &NewOddsInfo{
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