package main

import (
	"betsapiScrapper/crawler"
	"log"
)

func main() {
	betsapiCrawler := betsapi.BetsapiCrawler{}
	err := betsapiCrawler.Init()
	if err != nil{
		log.Fatal(err)
	}
}
