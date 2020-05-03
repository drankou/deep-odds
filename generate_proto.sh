#!/usr/bin/env bash

protoc -I=. -I=./vendor --go_out=./pkg/betsapi/types ./pkg/betsapi/types/types.proto