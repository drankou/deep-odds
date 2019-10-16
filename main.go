package main

import (
	"betsapiScrapper/crawler"
	"log"
)

func main() {
	betsapiCrawler := crawler.BetsapiCrawler{}
	err := betsapiCrawler.Init()
	if err != nil{
		log.Fatal(err)
	}
	betsapiCrawler.GetFootballEventsStats()
	betsapiCrawler.GetHotFootballEvents()
}
