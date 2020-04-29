package main

import (
	"betsapi_scrapper/service"
	"betsapi_scrapper/types"
	"betsapi_scrapper/utils"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	if err := godotenv.Load(utils.PathFromRoot("cmd", "betsapi",".env")); err != nil {
		log.Panicf("Error loading .env file. #%v", err)
	}

	grpcServer := grpc.NewServer()
	betsapiService := &service.BetsapiService{}
	err := betsapiService.Init()
	if err != nil {
		log.Fatal(err)
	}

	types.RegisterBetsapiServer(grpcServer, betsapiService)
	lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("BetsapiService is listening on localhost%s.", os.Getenv("SERVER_PORT"))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
