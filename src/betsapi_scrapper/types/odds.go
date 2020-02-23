package types

import (
	"github.com/sirupsen/logrus"
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
	Minute  string `json:"time_str,omitempty" bson:"minute"`
	AddTime string `json:"add_time,omitempty" bson:"add_time"`
}

func (result *Result) ToNew() *NewResult {
	minute, err := strconv.ParseInt(result.Minute, 10, 64)
	if err != nil {
		logrus.Error(err)
	}

	addTime, err := strconv.ParseInt(result.AddTime, 10, 64)
	if err != nil {
		logrus.Error(err)
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
	Id string `json:"id,omitempty" bson:"-"`
	*NewResultOdds `bson:"odds"`
	*NewOddsInfo `bson:"odds_info"`
}

type ResultOdds struct {
	HomeOdd string `json:"home_od,omitempty" bson:"home_odds"`
	DrawOdd string `json:"draw_od,omitempty" bson:"draw_odds"`
	AwayOdd string `json:"away_od,omitempty" bson:"away_odds"`
}

func (resultOdds *ResultOdds) ToNew() *NewResultOdds {
	homeOdds, err := strconv.ParseFloat(resultOdds.HomeOdd, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	awayOdds, err := strconv.ParseFloat(resultOdds.AwayOdd, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	drawOdds, err := strconv.ParseFloat(resultOdds.DrawOdd, 64)
	if err != nil {
		logrus.Error(err)
		return nil
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

//func AddMissingResultOdds(resultOdds []Result) []Result {
//	var resultList []Result
//
//	minuteOdds := make(map[int]*Result)
//	var minutes []int
//	for i, entry := range resultOdds {
//		minute := int(entry.Minute)
//		minutes = append(minutes, minute)
//		minuteOdds[minute] = &resultOdds[i]
//	}
//	sort.Ints(minutes)
//
//	//check if first available minute is 0 - start of the match
//	if len(minutes) > 0 && minutes[0] != 0 {
//		// create first minute odds from the first available minute data
//		nextMinuteWithOdds := *minuteOdds[minutes[0]]
//		unixTime := nextMinuteWithOdds.AddTime
//		unixTime -= int64(minutes[0]) * 60
//
//		firstMinuteOdds := &Result{
//			Id:         nextMinuteWithOdds.Id + "0",
//			ResultOdds: nextMinuteWithOdds.ResultOdds,
//			OddsInfo: OddsInfo{
//				Score:   "0-0",
//				Minute:  0,
//				AddTime: unixTime,
//			},
//		}
//
//		minuteOdds[0] = firstMinuteOdds
//		minutes = append(minutes, 0)
//	}
//
//	lastMinute := 90
//	//TODO get last minute of the match
//	for i := 1; i <= lastMinute; i++ {
//		//check if minute presented in the minute->odds map
//		if minuteOdds[i] == nil {
//			//create minute odds with data from previous minute
//			previousMinute := i - 1
//			previousMinuteOdds := *minuteOdds[previousMinute]
//
//			unixTime := previousMinuteOdds.AddTime
//			unixTime += 60
//
//			newMinuteOdds := &Result{
//				Id:         previousMinuteOdds.Id + strconv.Itoa(i),
//				ResultOdds: previousMinuteOdds.ResultOdds,
//				OddsInfo: OddsInfo{
//					Score:   previousMinuteOdds.Score,
//					Minute:  int32(i),
//					AddTime: unixTime,
//				},
//			}
//
//			minuteOdds[i] = newMinuteOdds
//			minutes = append(minutes, i)
//		}
//	}
//
//	sort.Ints(minutes)
//
//	for _, minute := range minutes {
//		resultList = append(resultList, *minuteOdds[minute])
//	}
//
//	return resultList
//}

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
		logrus.Error(err)
	}

	addTime, err := strconv.ParseInt(asianRes.AddTime, 10, 64)
	if err != nil {
		logrus.Error(err)
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
	Id string `json:"id,omitempty" bson:"-"`
	*NewAsianHandicapResultOdds `bson:"odds"`
	*NewOddsInfo `bson:"odds_info"`
}

type AsianHandicapResultOdds struct {
	HomeOdd  string `json:"home_od,omitempty" bson:"home_odds"`
	Handicap string `json:"handicap,omitempty" bson:"handicap"`
	AwayOdd  string `json:"away_od,omitempty" bson:"away_odds"`
}

func (odds *AsianHandicapResultOdds) ToNew() *NewAsianHandicapResultOdds {
	homeOdds, err := strconv.ParseFloat(odds.HomeOdd, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	awayOdds, err := strconv.ParseFloat(odds.AwayOdd, 64)
	if err != nil {
		logrus.Error(err)
		return nil
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
	Minute  string `json:"time_str,omitempty" bson:"minute"`
	AddTime string `json:"add_time,omitempty" bson:"add_time"`
}

func (asianTotal *AsianHandicapTotal) ToNew() *NewAsianHandicapTotal {
	minute, err := strconv.ParseInt(asianTotal.Minute, 10, 64)
	if err != nil {
		logrus.Error(err)
	}

	addTime, err := strconv.ParseInt(asianTotal.AddTime, 10, 64)
	if err != nil {
		logrus.Error(err)
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
	Id string `json:"id" bson:"-"`
	*NewAsianHandicapTotalOdds `bson:"odds"`
	*NewOddsInfo `bson:"odds_info"`
}

type OddsInfo struct {
	Score   string `json:"ss" bson:"score"`
	Minute  string `json:"time_str,omitempty" bson:"minute"`
	AddTime string `json:"add_time,omitempty" bson:"add_time"`
}

func (oddsInfo *OddsInfo) ToNew() *NewOddsInfo {
	minute, err := strconv.ParseInt(oddsInfo.Minute, 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	addTime, err := strconv.ParseInt(oddsInfo.AddTime, 10, 64)
	if err != nil {
		logrus.Error(err)
		return nil
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
		logrus.Error(err)
		return nil
	}

	awayOdds, err := strconv.ParseFloat(odds.UnderOdd, 64)
	if err != nil {
		logrus.Error(err)
		return nil
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
