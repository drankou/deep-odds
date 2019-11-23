package exporter

import (
	"log"
	"os"
	"testing"
)

func TestExporter_Init(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")

	e := Exporter{}
	err := e.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestExporter_GetFootballEventById(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")

	e := Exporter{}
	err := e.Init()
	if err != nil {
		t.Fatal(err)
	}

	footballEvent := e.GetFootballEventById("1679316")
	if footballEvent == nil {
		t.Fatal("Empty result")
	}

	log.Printf("Event:%+v", footballEvent.Event)
	log.Printf("Odds:%+v", footballEvent.Odds)
	log.Printf("History:%+v", footballEvent.History)
	log.Printf("StatsTrend:%+v", footballEvent.StatsTrend)
}

func TestExporter_GetFootballEventsByLeague(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")

	e := Exporter{}
	err := e.Init()
	if err != nil {
		t.Fatal(err)
	}

	footballEvents := e.GetFootballEventsByLeague("176")
	if len(footballEvents) == 0 {
		t.Fatal("Empty result")
	}
}
