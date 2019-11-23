package main

import (
	"betsapiScrapper/betsapi"
	"log"
)

func main() {
	betsapiCrawler := betsapi.BetsapiCrawler{}
	err := betsapiCrawler.Init()
	if err != nil {
		log.Fatal(err)
	}
}