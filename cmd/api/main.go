package main

import (
	"github.com/drankou/deep-odds/pkg/api"
	"github.com/drankou/deep-odds/pkg/api/types"
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
	deepOdds := &api.DeepOddsServer{}
	err := deepOdds.Init()
	if err != nil {
		log.Fatal(err)
	}

	types.RegisterDeepOddsServer(grpcServer, deepOdds)
	lis, err := net.Listen("tcp", "localhost"+os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("DeepOdds server is listening on localhost%s.", os.Getenv("SERVER_PORT"))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
