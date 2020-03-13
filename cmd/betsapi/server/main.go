package main

import (
	"betsapi_scrapper/service"
	"betsapi_scrapper/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

const defaultBetsapiServicePort = ":50001"

func main() {
	os.Setenv("BETSAPI_TOKEN", "25493-JGWujvhpW6upWr")
	grpcServer := grpc.NewServer()

	betsapiService := &service.BetsapiService{}
	err := betsapiService.Init()
	if err != nil {
		log.Fatal(err)
	}

	types.RegisterBetsapiServer(grpcServer, betsapiService)
	log.Info("BetsapiService is ready.")

	lis, err := net.Listen("tcp", defaultBetsapiServicePort)
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
