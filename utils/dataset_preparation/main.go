package main

import (
	"encoding/csv"
	"github.com/drankou/deep-odds/pkg/betsapi/types"
	"github.com/drankou/deep-odds/pkg/storage"
	log "github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	os.Setenv("ENVIRONMENT", "prod")
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	BasicCsvDataset()
}

func BasicCsvDataset() {
	file, err := os.Create("football_events.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	header := []string{
		"event.id",
		"event.start_time",
		"team.home", "team.away",
		"league.name", "league.id",
		"country.code",
		"minute",
		"goals.home", "goals.away",
		"attacks.home", "attacks.away",
		"dang_attacks.home", "dang_attacks.away",
		//"possession.home", "possession.away",
		"off_target.home", "off_target.away",
		"on_target.home", "on_target.away",
		"corners.home", "corners.away",
		"yellow_cards.home", "yellow_cards.away",
		"red_cards.home", "red_cards.away",
		"substitutions.home", "substitutions.away",
		"odds.home", "odds.draw", "odds.away",
		"start_odds.home", "start_odds.draw", "start_odds.away",
		"final_score",
		"result",
	}
	//TODO check possession and to to the result dataset

	err = csvWriter.Write(header)
	if err != nil {
		log.Error(err)
	}

	mongo := storage.GetMongoWrapper()
	where := map[string]interface{}{
		"stats_trend.attacks.home": map[string]interface{}{
			"$ne": nil,
		},
	}

	eventChan, err := mongo.StreamAll("football_event-v2", reflect.TypeOf(types.FootballEvent{}), where)
	if err != nil {
		log.Fatal(err)
	}

	for event := range eventChan {
		var base []string
		footballEvent := event.(*types.FootballEvent)
		statsTrend := types.AddMissingStatsTrend(footballEvent.GetStatsTrend())
		//TODO Use map minute -> odds
		odds := types.AddMissingResultOdds(footballEvent.GetOdds().GetFullTime())
		finalScore := footballEvent.GetEvent().GetScore()

		base = append(base, footballEvent.GetEvent().GetId())
		base = append(base, time.Unix(footballEvent.GetEvent().GetTime(), 0).UTC().String())
		base = append(base, footballEvent.GetEvent().GetHomeTeam().GetName())
		base = append(base, footballEvent.GetEvent().GetAwayTeam().GetName())
		base = append(base, footballEvent.GetEvent().GetLeague().GetName())
		base = append(base, footballEvent.GetEvent().GetLeague().GetId())
		base = append(base, footballEvent.GetEvent().GetLeague().GetCountryCode())

		for i := 0; i <= 90; i++ {
			var record []string

			//event info
			record = append(record, base...)

			//stats
			record = append(record, strconv.Itoa(i)) //minute
			record = append(record, strconv.FormatInt(statsTrend.GetGoals().GetHome()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetGoals().GetAway()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetAttacks().GetHome()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetAttacks().GetAway()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetDangerousAttacks().GetHome()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetDangerousAttacks().GetAway()[i].GetValue(), 10))
			//record = append(record, strconv.FormatInt(statsTrend.Possession.GetAway()[i].GetValue(), 10))
			//record = append(record, string(statsTrend.Possession.GetAway()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetOffTarget().GetHome()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetOffTarget().GetAway()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetOnTarget().GetHome()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetOnTarget().GetAway()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetCorners().GetHome()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetCorners().GetAway()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetYellowCards().GetHome()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetYellowCards().GetAway()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetRedCards().GetHome()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetRedCards().GetAway()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetSubstitutions().GetHome()[i].GetValue(), 10))
			record = append(record, strconv.FormatInt(statsTrend.GetSubstitutions().GetAway()[i].GetValue(), 10))

			//actual odds
			record = append(record, strconv.FormatFloat(odds[i].GetHomeOdds(), 'f', 3, 64))
			record = append(record, strconv.FormatFloat(odds[i].GetDrawOdds(), 'f', 3, 64))
			record = append(record, strconv.FormatFloat(odds[i].GetAwayOdds(), 'f', 3, 64))
			//start odds
			record = append(record, strconv.FormatFloat(odds[0].GetHomeOdds(), 'f', 3, 64))
			record = append(record, strconv.FormatFloat(odds[0].GetDrawOdds(), 'f', 3, 64))
			record = append(record, strconv.FormatFloat(odds[0].GetAwayOdds(), 'f', 3, 64))

			record = append(record, finalScore)

			goals := strings.Split(finalScore, "-")
			goalsHome, err := strconv.Atoi(goals[0])
			if err != nil{
				log.Print(footballEvent.Event.Id)
				continue
			}
			goalsAway, err := strconv.Atoi(goals[1])
			if err != nil{
				log.Print(footballEvent.Event.Id)
				continue
			}

			if goalsHome > goalsAway {
				record = append(record, []string{"1"}...)
			} else if goalsHome < goalsAway {
				record = append(record, []string{"2"}...)
			} else {
				record = append(record, []string{"0"}...)
			}

			err = csvWriter.Write(record)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
