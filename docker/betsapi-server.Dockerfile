# gcr.io/deep-odds/betsapi-server
FROM golang:latest as builder

WORKDIR /go/src/github.com/drankou/deep-odds

COPY ./cmd/betsapi /go/src/github.com/drankou/deep-odds/cmd/betsapi
COPY ./pkg/betsapi /go/src/github.com/drankou/deep-odds/pkg/betsapi
COPY ./pkg/storage /go/src/github.com/drankou/deep-odds/pkg/storage
COPY ./pkg/utils /go/src/github.com/drankou/deep-odds/pkg/utils

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build /go/src/github.com/drankou/deep-odds/cmd/betsapi/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/drankou/deep-odds/main /main

CMD ["/main"]