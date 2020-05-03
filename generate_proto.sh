#!/usr/bin/env bash

protoc -I=. -I=./vendor --go_out=paths=source_relative:. ./pkg/betsapi/types/types.proto