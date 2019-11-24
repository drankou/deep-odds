package storage

import (
	"betsapi_scrapper/types"
	"os"
	"testing"
)

func TestMongoWrapper_Connect(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	mongo := MongoWrapper{}
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	mongo.Connect(mongoConnectionString, DefaultMongoDatabase)
}

func TestMongoWrapper_Insert(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
	mongo := MongoWrapper{}
	mongoConnectionString := os.Getenv("MONGO_CONNECTION_STRING")

	mongo.Connect(mongoConnectionString, DefaultMongoDatabase)

	footbalEvent := &types.FootballEvent{
		Event:      &types.Event{},
		History:    &types.EventHistory{},
		Odds:       &types.Odds{},
		StatsTrend: &types.StatsTrend{},
	}
	_, err := mongo.Insert("football_event", footbalEvent)
	if err != nil {
		t.Fatal(err)
	}
}
