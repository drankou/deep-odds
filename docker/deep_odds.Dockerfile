# gcr.io/deep-odds/deep-odds
FROM golang:latest as builder

WORKDIR /go/src/github.com/drankou/deep-odds

COPY . /go/src/github.com/drankou/deep-odds
RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build /go/src/github.com/drankou/deep-odds/cmd/api/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/drankou/deep-odds/main /main

CMD ["/main"]