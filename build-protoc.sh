#!/bin/bash

cd src/vendor

protoc --go_out=. tensorflow/core/framework/*.proto
protoc --go_out=plugins=grpc:. tensorflow_serving/apis/*.proto
