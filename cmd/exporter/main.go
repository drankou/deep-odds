package main

import (
	"github.com/drankou/deep-odds/pkg/exporter"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	//Load environmental variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//initialize betsapi exporter
	betsapiExporter := exporter.Exporter{}
	err = betsapiExporter.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Exporter initialized")

	//from 01.06.2020 to now
	betsapiExporter.ExportFootballEventsFromDate(1583625600, 1595096336)
}
