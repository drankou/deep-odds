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

	clean := RemoveDuplicitResultOdds(odds.GetFullTime())
	if len(clean) >= len(odds.GetFullTime()) {
		t.Fatal("number of odds should be less after odds cleaning")
	}
	t.Logf("Initial number of odds ticks: %d", len(odds.GetFullTime()))
	t.Logf("Number of odds after dupicities removing: %d", len(clean))
}

func TestAddMissingResultOdds(t *testing.T) {
	data, err := ioutil.ReadFile(path.Join(utils.GetAbsPathToRoot(), "mock_data", "odds.bson"))

	var odds Odds
	err = bson.Unmarshal(data, &odds)
	if err != nil {
		log.Fatal(err)
	}

	fullTimeResultOdds := RemoveDuplicitResultOdds(odds.GetFullTime())
	fullTimeResultOdds = AddMissingResultOdds(fullTimeResultOdds)
	//log.Print("Result odds length: ", len(odds))
	for _, odd := range fullTimeResultOdds {
		log.Print("minute: ", odd.GetMinute())
		log.Print("score: ", odd.GetScore())
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
		log.Print("minute: ", odd.GetMinute())
		log.Print("score: ", odd.GetScore())
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

	asianHandicapTotal := RemoveDuplicitAsianHandicapTotal(odds.GetTotal())
	asianHandicapTotal = AddMissingAsianTotalOdds(asianHandicapTotal)
	for _, odds := range asianHandicapTotal {
		log.Print("minute: ", odds.GetMinute())
		log.Print("score: ", odds.GetScore())
		log.Print(strings.Repeat("-", 20))
	}
}
