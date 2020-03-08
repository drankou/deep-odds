package main

import (
	"betsapi_scrapper/storage"
	"betsapi_scrapper/types"
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	os.Setenv("ENVIRONMENT", "prod")
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	BasicCsvDataset()
}

func BasicCsvDataset() {
	file, err := os.Create("dataset.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	header := []string{"home_team", "away_team", "league_id", "cc", "minute", "home_goals", "away_goals", "home_attacks", "away_attacks", "home_dang_attacks", "away_dang_attacks", "home_off_target", "away_off_target", "home_on_target", "away_on_target", "home_corners", "away_corners", "home_yellow_cards", "away_yellow_cards", "home_red_cards", "away_red_cards", "home_substitutions", "away_substitutions", "home_odds", "draw_odds", "away_odds", "home_win", "draw", "away_win"}
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

	eventChan, err := mongo.StreamAll("football_event", reflect.TypeOf(types.FootballEvent{}), where)
	if err != nil {
		log.Fatal(err)
	}

	for event := range eventChan {
		var base []string
		footballEvent := event.(*types.FootballEvent)
		statsTrend := types.AddMissingStatsTrend(footballEvent.StatsTrend)
		var odds []*types.Result
		if footballEvent.Odds != nil {
			odds = types.RemoveDuplicitResultOdds(footballEvent.Odds.FullTimeResult)
			odds = types.AddMissingResultOdds(odds)
		}

		if len(odds) == 0 {
			for i := 0; i < 90; i++ {
				emptyResult := &types.Result{
					ResultOdds: &types.ResultOdds{
						HomeOdd: -1,
						DrawOdd: -1,
						AwayOdd: -1,
					},
				}
				odds = append(odds, emptyResult)
			}
		}

		score := footballEvent.Event.Score
		goals := strings.Split(score, "-")
		if len(goals) != 2 {
			continue
		}

		homeTeamGoals, _ := strconv.Atoi(goals[0])
		awayTeamGoals, _ := strconv.Atoi(goals[1])

		base = append(base, footballEvent.Event.HomeTeam.Name)
		base = append(base, footballEvent.Event.AwayTeam.Name)
		base = append(base, footballEvent.Event.League.Id)
		base = append(base, footballEvent.Event.League.CountryCode)

		for i := 0; i < 90; i++ {
			var record []string

			record = append(record, base...)
			record = append(record, strconv.Itoa(i))

			record = append(record, strconv.FormatInt(statsTrend.Goals.Home[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.Goals.Away[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.Attacks.Home[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.Attacks.Away[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.DangerousAttacks.Home[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.DangerousAttacks.Away[i].Value, 10))
			//record = append(record, strconv.FormatInt(statsTrend.Possession.Away[i].Value, 10))
			//record = append(record, string(statsTrend.Possession.Away[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.OffTarget.Home[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.OffTarget.Away[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.OnTarget.Home[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.OnTarget.Away[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.Corners.Home[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.Corners.Away[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.YellowCards.Home[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.YellowCards.Away[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.RedCards.Home[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.RedCards.Away[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.Substitutions.Away[i].Value, 10))
			record = append(record, strconv.FormatInt(statsTrend.Substitutions.Away[i].Value, 10))

			record = append(record, strconv.FormatFloat(odds[i].HomeOdd, 'f', 3, 64))
			record = append(record, strconv.FormatFloat(odds[i].DrawOdd, 'f', 3, 64))
			record = append(record, strconv.FormatFloat(odds[i].AwayOdd, 'f', 3, 64))

			if homeTeamGoals > awayTeamGoals {
				record = append(record, "1", "0", "0")
			} else if homeTeamGoals < awayTeamGoals {
				record = append(record, "0", "0", "1")
			} else {
				record = append(record, "0", "1", "0")
			}

			err := csvWriter.Write(record)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
