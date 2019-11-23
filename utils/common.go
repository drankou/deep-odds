package utils

import "betsapiScrapper/types"

func UpdateStringSlice(idxsToDelete []int, slice []string) []string {
	var newSlice []string
	var start int
	var end int
	for _, index := range idxsToDelete {
		end = index
		newSlice = append(newSlice, slice[start:end]...)
		start = end + 1
	}
	newSlice = append(newSlice, slice[start:]...)

	return newSlice
}

func UpdateIntSlice(idxsToDelete []int, slice []int) []int {
	var newSlice []int
	var start int
	var end int
	for _, index := range idxsToDelete {
		end = index
		newSlice = append(newSlice, slice[start:end]...)
		start = end + 1
	}
	newSlice = append(newSlice, slice[start:]...)

	return newSlice
}

// Remove duplicities and empty strings from list of keywords
func RemoveDuplicitOdds(odds []types.IOdds) []types.IOdds {
	var resultList []types.IOdds
	keys := make(map[string]bool)
	for _, entry := range odds {
		if _, value := keys[entry.GetTime()]; !value && entry.GetTime() != "" && entry.GetHome() != "-" && entry.GetAway() != "-" && entry.GetValue() != "-" {
			keys[entry.GetTime()] = true
			resultList = append(resultList, entry)
		}
	}

	return resultList
}

// Remove duplicities and empty strings from list of keywords
func RemoveDuplicities(stringSlice []string) []string {
	var resultList []string
	keys := make(map[string]bool)
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value && entry != "" && entry != " " {
			keys[entry] = true
			resultList = append(resultList, entry)
		}
	}

	return resultList
}

// Remove duplicities int values from list
func RemoveDuplicitiesInt(intSlice []int) []int {
	var resultList []int
	keys := make(map[int]bool)
	for _, entry := range intSlice {
		if _, ok := keys[entry]; !ok {
			keys[entry] = true
			resultList = append(resultList, entry)
		}
	}

	return resultList
}
