package main

import (
	"github.com/drankou/deep-odds/pkg/exporter"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const LeaguesTotal = 1665

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

	//start, err := strconv.Atoi(os.Getenv("START_LEAGUE"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//end, err := strconv.Atoi(os.Getenv("END_LEAGUE"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i := start; i <= end; i++ {
	//	leagueId := strconv.FormatInt(int64(i), 10)
	//	betsapiExporter.ExportFootballEventsByLeague(leagueId)
	//}
}
