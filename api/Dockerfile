FROM golang:1.16.3

ENV GO111MODULE=on

WORKDIR /go/src/github.com/mashita1023/clean-architecture

RUN go get -u github.com/cosmtrek/air
# RUN go mod tidy
CMD air
