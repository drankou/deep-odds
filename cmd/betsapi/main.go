package main

import (
	"github.com/drankou/deep-odds/pkg/betsapi/service"
	"github.com/drankou/deep-odds/pkg/betsapi/types"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	if os.Getenv("ENVIRONMENT") == "dev" {
		if err := godotenv.Load(); err != nil {
			log.Panicf("Error loading .env file. #%v", err)
		}
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
