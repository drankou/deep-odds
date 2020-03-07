package storage

import (
	"betsapi_scrapper/types"
	"betsapi_scrapper/utils"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestMongoWrapper_Connect(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	mongo := MongoWrapper{}
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	mongo.Connect(mongoConnectionString, DefaultMongoDatabase)
}

func TestMongoWrapper_ReadOne(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	mongo := GetMongoWrapper()

	filter := make(map[string]interface{})
	data, err := mongo.ReadOne("football_event", reflect.TypeOf(types.FootballEvent{}), filter)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := data.(*types.FootballEvent); !ok {
		t.Fatal("unsuccessful type assertion")
	}
}

func TestMongoWrapper_Insert(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	os.Setenv("ENVIRONMENT", "dev")

	mongo := GetMongoWrapper()

	footballEvent := &types.FootballEvent{
		Event: &types.Event{
			HomeTeam: &types.Team{
				Name:        "Naftan Novopolotsk",
				CountryCode: "by",
			},
			AwayTeam: &types.Team{
				Name:        "BATE Borisov",
				CountryCode: "by",
			},
		},
		History:    &types.EventHistory{},
		Odds:       &types.Odds{},
		StatsTrend: &types.StatsTrend{},
	}
	_, err := mongo.Insert("football_event", footballEvent)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMongoWrapper_ReadAll(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	mongo := GetMongoWrapper()

	filter := make(map[string]interface{})
	footballEvents, err := mongo.GetFootballEvents(filter)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Total number of football events: ", len(footballEvents))
}

func TestGetMongoWrapper_ReadAll_Where(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	mongo := GetMongoWrapper()

	filter := map[string]interface{}{
		"stats_trend.yellow_cards": map[string]interface{}{
			"$ne": nil,
		},
		"event.league.id": "874",
	}
	footballEvents, err := mongo.ReadAll("football_event", reflect.TypeOf(types.FootballEvent{}), filter)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Total number of football events with stats: ", len(footballEvents))
}

func TestMongoWrapper_StreamAll(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	os.Setenv("ENVIRONMENT", "prod")

	mongo := GetMongoWrapper()

	filter := map[string]interface{}{
		"stats_trend.attacks": map[string]interface{}{
			"$ne": nil,
		},
	}
	eventsChan, err := mongo.StreamAll("football_event", reflect.TypeOf(types.FootballEvent{}), filter)
	if err != nil {
		t.Fatal(err)
	}

	total := 0
	for footballEvent := range eventsChan {
		total++
		log.Print(footballEvent.(*types.FootballEvent).Event.Id)
	}

	t.Log("Total number of football events with stats: ", total)
}

func TestDumpBsonStatsTrendForMock(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	mongo := GetMongoWrapper()

	filter := map[string]interface{}{
		"stats_trend.attacks.home": map[string]interface{}{
			"$ne": nil,
		},
	}

	entry, err := mongo.ReadOne("football_event", reflect.TypeOf(types.FootballEvent{}), filter)
	if err != nil {
		t.Fatal(err)
	}

	footballEvent := entry.(*types.FootballEvent)
	statsTrend := footballEvent.StatsTrend

	//log.Print("EventId: ", footballEvent.Event.Id)
	data, err := bson.Marshal(statsTrend)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(path.Join(utils.GetAbsPathToRoot(), "mock_data", "stats_trend.bson"), data, 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDumpBsonOddsForMock(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	mongo := GetMongoWrapper()

	filter := map[string]interface{}{
		"odds": map[string]interface{}{
			"$ne": nil,
		},
	}

	entry, err := mongo.ReadOne("football_event", reflect.TypeOf(types.FootballEvent{}), filter)
	if err != nil {
		t.Fatal(err)
	}

	footballEvent := entry.(*types.FootballEvent)
	odds := footballEvent.Odds

	//log.Print("EventId: ", footballEvent.Event.Id)
	data, err := bson.Marshal(odds)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(path.Join(utils.GetAbsPathToRoot(), "mock_data", "odds.bson"), data, 0644)
	if err != nil {
		t.Fatal(err)
	}
}