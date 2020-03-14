#!/usr/bin/env bash

protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gofast_out=plugins=grpc:./src/betsapi_scrapper/types ./src/betsapi_scrapper/types/types.proto