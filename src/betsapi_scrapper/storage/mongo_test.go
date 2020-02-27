package storage

import (
	"betsapi_scrapper/types"
	"encoding/json"
	"log"
	"os"
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
	mongo := MongoWrapper{}
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	mongo.Connect(mongoConnectionString, DefaultMongoDatabase)

	filter := make(map[string]interface{})
	data, err := mongo.ReadOne("football_event", reflect.TypeOf(types.FootballEvent{}), filter)
	if err != nil{
		t.Fatal(err)
	}

	jsonOld, err := json.Marshal(data)
	if err != nil{
		t.Fatal(err)
	}

	log.Print(string(jsonOld))

	var newFootballEvent types.NewFootballEvent

	err = json.Unmarshal(jsonOld, &newFootballEvent)
	if err != nil{
		t.Fatal(err)
	}

	_, err = mongo.Insert("football_event", newFootballEvent)
	if err != nil{
		t.Fatal(err)
	}

	jsonNew, err := json.Marshal(newFootballEvent)
	if err != nil{
		t.Fatal(err)
	}

	log.Print(string(jsonNew))
}

func TestMongoWrapper_Insert(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	os.Setenv("ENVIRONMENT", "dev")

	mongo := MongoWrapper{}
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	mongo.Connect(mongoConnectionString, DefaultMongoDatabase)

	footballEvent := &types.FootballEvent{
		Event: &types.Event{
			HomeTeam: types.Team{
				Name:        "Naftan Novopolotsk",
				CountryCode: "by",
			},
			AwayTeam: types.Team{
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
	os.Setenv("ENVIRONMENT", "prod")

	mongo := MongoWrapper{}
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	mongo.Connect(mongoConnectionString, DefaultMongoDatabase)

	filter := make(map[string]interface{})
	footballEvents, err := mongo.GetFootballEvents(filter)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Total number of football events: ", len(footballEvents))
}

func TestGetMongoWrapper_ReadAll_Where(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	os.Setenv("ENVIRONMENT", "prod")

	mongo := MongoWrapper{}
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	mongo.Connect(mongoConnectionString, DefaultMongoDatabase)

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

	mongo := MongoWrapper{}
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	mongo.Connect(mongoConnectionString, DefaultMongoDatabase)

	filter := map[string]interface{}{
		"stats_trend.attacks": map[string]interface{}{
			"$ne": nil,
		},
	}
	eventsChan, err := mongo.StreamAll("football_event", reflect.TypeOf(types.FootballEvent{}), filter)
	if err != nil {
		t.Fatal(err)
	}

	for footballEvent := range eventsChan {
		log.Print(footballEvent.(*types.FootballEvent).Event.Id)
	}

	//t.Log("Total number of football events with stats: ", len(footballEvents))
}
