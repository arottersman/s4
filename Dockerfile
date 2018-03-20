FROM golang:1.9-alpine

RUN apk add --no-cache wget curl git build-base

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/go-redis/redis
