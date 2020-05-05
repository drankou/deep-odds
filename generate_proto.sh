#!/usr/bin/env bash

protoc -I=. -I=$GOPATH/src --gofast_out=plugins=grpc,paths=source_relative:. ./pkg/betsapi/types/types.proto