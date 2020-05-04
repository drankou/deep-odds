#!/usr/bin/env bash

protoc -I=. -I=./vendor --gofast_out=plugins=grpc,paths=source_relative:. ./pkg/betsapi/types/types.proto