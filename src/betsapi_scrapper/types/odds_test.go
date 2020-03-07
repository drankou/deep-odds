package types

import (
	"betsapi_scrapper/utils"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"testing"
)

func TestRemoveDuplicitResultOdds(t *testing.T) {
	data, err := ioutil.ReadFile(path.Join(utils.GetAbsPathToRoot(), "mock_data", "odds.bson"))

	var odds Odds
	err = bson.Unmarshal(data, &odds)
	if err != nil {
		log.Fatal(err)
	}

	clean := RemoveDuplicitResultOdds(odds.FullTimeResult)
	if len(clean) >= len(odds.FullTimeResult) {
		t.Fatal("number of odds should be less after odds cleaning")
	}
	t.Logf("Initial number of odds ticks: %d", len(odds.FullTimeResult))
	t.Logf("Number of odds after dupicities removing: %d", len(clean))
}

func TestAddMissingResultOdds(t *testing.T) {
	data, err := ioutil.ReadFile(path.Join(utils.GetAbsPathToRoot(), "mock_data", "odds.bson"))

	var odds Odds
	err = bson.Unmarshal(data, &odds)
	if err != nil {
		log.Fatal(err)
	}

	fullTimeResultOdds := RemoveDuplicitResultOdds(odds.FullTimeResult)
	fullTimeResultOdds = AddMissingResultOdds(fullTimeResultOdds)
	//log.Print("Result odds length: ", len(odds))
	for _, odd := range fullTimeResultOdds {
		log.Print("minute: ", odd.Minute)
		log.Print("score: ", odd.Score)
		log.Printf("%+v", odd.ResultOdds)
		log.Print(strings.Repeat("-", 20))
	}
}

func TestAddMissingAsianResultOdds(t *testing.T) {
	data, err := ioutil.ReadFile(path.Join(utils.GetAbsPathToRoot(), "mock_data", "odds.bson"))

	var odds Odds
	err = bson.Unmarshal(data, &odds)
	if err != nil {
		log.Fatal(err)
	}

	asianHandicapResultOdds := RemoveDuplicitAsianHandicapResult(odds.AsianHandicap)
	asianHandicapResultOdds = AddMissingAsianResultOdds(asianHandicapResultOdds)
	for _, odd := range asianHandicapResultOdds {
		log.Print("minute: ", odd.Minute)
		log.Print("score: ", odd.Score)
		log.Printf("%+v", odd.AsianHandicapResultOdds)
		log.Print(strings.Repeat("-", 20))
	}
}

func TestAddMissingAsianTotalOdds(t *testing.T) {
	data, err := ioutil.ReadFile(path.Join(utils.GetAbsPathToRoot(), "mock_data", "odds.bson"))

	var odds Odds
	err = bson.Unmarshal(data, &odds)
	if err != nil {
		log.Fatal(err)
	}

	asianHandicapTotal := RemoveDuplicitAsianHandicapTotal(odds.GoalLineTotal)
	asianHandicapTotal = AddMissingAsianTotalOdds(asianHandicapTotal)
	for _, odds := range asianHandicapTotal {
		log.Print("minute: ", odds.Minute)
		log.Print("score: ", odds.Score)
		log.Printf("%+v", odds.AsianHandicapTotalOdds)
		log.Print(strings.Repeat("-", 20))
	}
}
