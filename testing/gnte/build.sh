#!/bin/sh

docker build -t gnte .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./scripts/peernode ../peerNode/main.go
Chmod 777 ./scripts/peernode
