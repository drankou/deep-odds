package main

import (
	"betsapiScrapper/src/exporter"
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
}
