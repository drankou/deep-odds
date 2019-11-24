package exporter

import (
	"log"
	"os"
	"testing"
)

func TestExporter_Init(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb+srv://alex:rffBz9nRs8H8N@betsapi-data-ksacs.gcp.mongodb.net/test?retryWrites=true&w=majority")

	e := Exporter{}
	err := e.Init()
	if err != nil {
		t.Fatal(err)
	}
}

func TestExporter_GetFootballEventById(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb+srv://alex:rffBz9nRs8H8N@betsapi-data-ksacs.gcp.mongodb.net/test?retryWrites=true&w=majority")

	e := Exporter{}
	err := e.Init()
	if err != nil {
		t.Fatal(err)
	}

	footballEvent := e.GetFootballEventById("196381")
	if footballEvent == nil {
		t.Fatal("Empty result")
	}

	footballEvent.Clean()
	log.Printf("Event:%+v", footballEvent.Event)
	log.Printf("Odds:%+v", footballEvent.Odds)
	log.Printf("History:%+v", footballEvent.History)
	log.Printf("StatsTrend:%+v", footballEvent.StatsTrend)
	e.SaveFootballEventToMongo(footballEvent)
}

func TestExporter_GetFootballEventsByLeague(t *testing.T) {
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")

	e := Exporter{}
	err := e.Init()
	if err != nil {
		t.Fatal(err)
	}

	e.ExportFootballEventsByLeague("94")
}
