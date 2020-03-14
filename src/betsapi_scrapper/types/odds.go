package types

func (m *Odds) Clean() {
	//result
	//m.FullTime = RemoveDuplicitResultOdds(m.GetFullTime())
	//m.FirstHalf = RemoveDuplicitResultOdds(m.GetFirstHalf())
	//
	////handicap
	//m.AsianHandicap = RemoveDuplicitAsianHandicapResult(m.GetAsianHandicap())
	//m.FirstHalfAsianHandicap = RemoveDuplicitAsianHandicapResult(m.GetFirstHalfAsianHandicap())
	//
	////totals
	//m.Total = RemoveDuplicitAsianHandicapTotal(m.GetTotal())
	//m.AsianCorners = RemoveDuplicitAsianHandicapTotal(m.GetAsianCorners())
	//m.FirstHalfTotal = RemoveDuplicitAsianHandicapTotal(m.GetFirstHalfTotal())
	//m.FirstHalfAsianCorners = RemoveDuplicitAsianHandicapTotal(m.GetFirstHalfAsianCorners())
}

//func RemoveDuplicitResultOdds(resultOdds []*Result) []*Result {
//	var resultList []*Result
//	keys := make(map[int64]bool)
//	for i, entry := range resultOdds {
//		if _, value := keys[entry.TimeStr]; !value && (entry.HomeOd != -1 || entry.AwayOd != -1 || entry.DrawOd != -1) {
//			keys[entry.TimeStr] = true
//			resultList = append(resultList, resultOdds[i])
//		}
//	}
//
//	return resultList
//}

//func AddMissingResultOdds(resultOdds []*Result) []*Result {
//	var resultList []*Result
//
//	minuteOdds := make(map[int64]*Result)
//	var minutes []int64
//	for i, entry := range resultOdds {
//		minute := entry.TimeStr
//		minutes = append(minutes, minute)
//		minuteOdds[minute] = resultOdds[i]
//	}
//
//	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })
//
//	if len(minutes) == 0 {
//		firstMinuteOdds := &Result{
//			Id:      "-1",
//			HomeOd:  -1,
//			DrawOd:  -1,
//			AwayOd:  -1,
//			Ss:      "0-0",
//			TimeStr: 0,
//		}
//
//		minuteOdds[0] = firstMinuteOdds
//		minutes = append(minutes, 0)
//	}
//
//	//check if first available minute is 0 - start of the match
//	if len(minutes) > 0 && minutes[0] != 0 {
//		// create first minute odds from the first available minute data
//		nextMinuteWithOdds := minuteOdds[minutes[0]]
//		var firstMinuteOdds *Result
//
//		if nextMinuteWithOdds != nil {
//			unixTime := nextMinuteWithOdds.AddTime
//			unixTime -= minutes[0] * 60
//
//			if nextMinuteWithOdds.GetSs() == "0-0" {
//				firstMinuteOdds = &Result{
//					Id:      nextMinuteWithOdds.GetId() + "0",
//					HomeOd:  nextMinuteWithOdds.GetHomeOd(),
//					DrawOd:  nextMinuteWithOdds.GetDrawOd(),
//					AwayOd:  nextMinuteWithOdds.GetAwayOd(),
//					Ss:      nextMinuteWithOdds.GetSs(),
//					TimeStr: 0,
//					AddTime: unixTime,
//				}
//			} else {
//				firstMinuteOdds = &Result{
//					Id:      nextMinuteWithOdds.Id + "0",
//					HomeOd:  -1,
//					DrawOd:  -1,
//					AwayOd:  -1,
//					Ss:      "0-0",
//					TimeStr: 0,
//					AddTime: unixTime,
//				}
//			}
//		} else {
//			firstMinuteOdds = &Result{
//				Id:      "-1",
//				HomeOd:  -1,
//				DrawOd:  -1,
//				AwayOd:  -1,
//				Ss:      "0-0",
//				TimeStr: 0,
//			}
//		}
//
//		minuteOdds[0] = firstMinuteOdds
//		minutes = append(minutes, 0)
//	}
//
//	lastMinute := int64(90)
//	//TODO get last minute of the match
//	for i := int64(1); i <= lastMinute; i++ {
//		//check if minute presented in the minute->odds map
//		if minuteOdds[i] == nil {
//			//create minute odds with data from previous minute
//			previousMinute := i - 1
//			previousMinuteOdds := minuteOdds[previousMinute]
//
//			unixTime := previousMinuteOdds.AddTime
//			unixTime += 60
//
//			newMinuteOdds := &Result{
//				Id:      previousMinuteOdds.Id + strconv.FormatInt(i, 10),
//				HomeOd:  previousMinuteOdds.GetHomeOd(),
//				DrawOd:  previousMinuteOdds.GetDrawOd(),
//				AwayOd:  previousMinuteOdds.GetAwayOd(),
//				Ss:      previousMinuteOdds.GetSs(),
//				TimeStr: i,
//				AddTime: unixTime,
//			}
//
//			minuteOdds[i] = newMinuteOdds
//			minutes = append(minutes, i)
//		}
//	}
//
//	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })
//
//	for _, minute := range minutes {
//		resultList = append(resultList, minuteOdds[minute])
//	}
//
//	return resultList
//}
//
//func RemoveDuplicitAsianHandicapResult(resultOdds []*AsianHandicapResult) []*AsianHandicapResult {
//	var resultList []*AsianHandicapResult
//	keys := make(map[int64]bool)
//	for i, entry := range resultOdds {
//		if _, value := keys[entry.GetTimeStr()]; !value && (entry.GetHomeOd() != -1 || entry.GetAwayOd() != -1) {
//			keys[entry.GetTimeStr()] = true
//			resultList = append(resultList, resultOdds[i])
//		}
//	}
//
//	return resultList
//}
//
//func AddMissingAsianResultOdds(asianResultOdds []*AsianHandicapResult) []*AsianHandicapResult {
//	var resultList []*AsianHandicapResult
//
//	minuteOdds := make(map[int64]*AsianHandicapResult)
//	var minutes []int64
//	for i, entry := range asianResultOdds {
//		minute := entry.GetTimeStr()
//		minutes = append(minutes, minute)
//		minuteOdds[minute] = asianResultOdds[i]
//	}
//
//	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })
//
//	//check if first available minute is 0 - start of the match
//	if len(minutes) > 0 && minutes[0] != 0 {
//		// create first minute odds from the first available minute data
//		nextMinuteWithOdds := minuteOdds[minutes[0]]
//		unixTime := nextMinuteWithOdds.AddTime
//		unixTime -= minutes[0] * 60
//
//		var firstMinuteOdds *AsianHandicapResult
//		if nextMinuteWithOdds.GetSs() == "0-0" {
//			firstMinuteOdds = &AsianHandicapResult{
//				Id:       nextMinuteWithOdds.Id + "0",
//				HomeOd:   nextMinuteWithOdds.GetHomeOd(),
//				Handicap: nextMinuteWithOdds.GetHandicap(),
//				AwayOd:   nextMinuteWithOdds.GetAwayOd(),
//				Ss:       nextMinuteWithOdds.GetSs(),
//				TimeStr:  0,
//				AddTime:  unixTime,
//			}
//		} else {
//			firstMinuteOdds = &AsianHandicapResult{
//				Id:       nextMinuteWithOdds.Id + "0",
//				HomeOd:   -1,
//				Handicap: "-1",
//				AwayOd:   -1,
//				Ss:       "0-0",
//				TimeStr:  0,
//				AddTime:  unixTime,
//			}
//		}
//
//		minuteOdds[0] = firstMinuteOdds
//		minutes = append(minutes, 0)
//	}
//
//	lastMinute := int64(90)
//	//TODO get last minute of the match
//	for i := int64(1); i <= lastMinute; i++ {
//		//check if minute presented in the minute->odds map
//		if minuteOdds[i] == nil {
//			//create minute odds with data from previous minute
//			previousMinute := i - 1
//			previousMinuteOdds := minuteOdds[previousMinute]
//
//			unixTime := previousMinuteOdds.AddTime
//			unixTime += 60
//
//			newMinuteOdds := &AsianHandicapResult{
//				Id:       previousMinuteOdds.Id + strconv.FormatInt(i, 10),
//				HomeOd:   previousMinuteOdds.GetHomeOd(),
//				Handicap: previousMinuteOdds.GetHandicap(),
//				AwayOd:   previousMinuteOdds.GetAwayOd(),
//				Ss:       previousMinuteOdds.GetSs(),
//				TimeStr:  i,
//				AddTime:  unixTime,
//			}
//
//			minuteOdds[i] = newMinuteOdds
//			minutes = append(minutes, i)
//		}
//	}
//
//	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })
//
//	for _, minute := range minutes {
//		resultList = append(resultList, minuteOdds[minute])
//	}
//
//	return resultList
//}
//
//func RemoveDuplicitAsianHandicapTotal(resultOdds []*AsianHandicapTotal) []*AsianHandicapTotal {
//	var resultList []*AsianHandicapTotal
//	keys := make(map[int64]bool)
//	for _, entry := range resultOdds {
//		if _, value := keys[entry.GetTimeStr()]; !value && (entry.GetOverOd() != -1 || entry.GetUnderOd() != -1) {
//			keys[entry.GetTimeStr()] = true
//			resultList = append(resultList, entry)
//		}
//	}
//
//	return resultList
//}
//
//func AddMissingAsianTotalOdds(asianTotalOdds []*AsianHandicapTotal) []*AsianHandicapTotal {
//	var resultList []*AsianHandicapTotal
//
//	minuteOdds := make(map[int64]*AsianHandicapTotal)
//	var minutes []int64
//	for i, entry := range asianTotalOdds {
//		minute := entry.GetTimeStr()
//		minutes = append(minutes, minute)
//		minuteOdds[minute] = asianTotalOdds[i]
//	}
//
//	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })
//
//	//check if first available minute is 0 - start of the match
//	if len(minutes) > 0 && minutes[0] != 0 {
//		// create first minute odds from the first available minute data
//		nextMinuteWithOdds := minuteOdds[minutes[0]]
//		unixTime := nextMinuteWithOdds.AddTime
//		unixTime -= minutes[0] * 60
//
//		var firstMinuteOdds *AsianHandicapTotal
//		if nextMinuteWithOdds.GetSs() == "0-0" {
//			firstMinuteOdds = &AsianHandicapTotal{
//				Id:       nextMinuteWithOdds.Id + "0",
//				OverOd:   nextMinuteWithOdds.GetOverOd(),
//				Handicap: nextMinuteWithOdds.GetHandicap(),
//				UnderOd:  nextMinuteWithOdds.GetUnderOd(),
//				Ss:       nextMinuteWithOdds.GetSs(),
//				TimeStr:  0,
//				AddTime:  unixTime,
//			}
//		} else {
//			firstMinuteOdds = &AsianHandicapTotal{
//				Id:       nextMinuteWithOdds.Id + "0",
//				OverOd:   -1,
//				Handicap: "-1",
//				UnderOd:  -1,
//				Ss:       "0-0",
//				TimeStr:  0,
//				AddTime:  unixTime,
//			}
//		}
//
//		minuteOdds[0] = firstMinuteOdds
//		minutes = append(minutes, 0)
//	}
//
//	lastMinute := int64(90)
//	//TODO get last minute of the match
//	for i := int64(1); i <= lastMinute; i++ {
//		//check if minute presented in the minute->odds map
//		if minuteOdds[i] == nil {
//			//create minute odds with data from previous minute
//			previousMinute := i - 1
//			previousMinuteOdds := minuteOdds[previousMinute]
//
//			unixTime := previousMinuteOdds.AddTime
//			unixTime += 60
//
//			newMinuteOdds := &AsianHandicapTotal{
//				Id:       previousMinuteOdds.Id + strconv.FormatInt(i, 10),
//				OverOd:   previousMinuteOdds.GetOverOd(),
//				Handicap: previousMinuteOdds.GetHandicap(),
//				UnderOd:  previousMinuteOdds.GetUnderOd(),
//				Ss:       previousMinuteOdds.GetSs(),
//				TimeStr:  i,
//				AddTime:  unixTime,
//			}
//
//			minuteOdds[i] = newMinuteOdds
//			minutes = append(minutes, i)
//		}
//	}
//
//	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })
//
//	for _, minute := range minutes {
//		resultList = append(resultList, minuteOdds[minute])
//	}
//
//	return resultList
//}
