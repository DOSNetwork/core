# Usage:
# $ make == $ make build
# $ make install

.PHONY: dep client install 
.DEFAULT_GOAL := build

build: dep client

dep:
	dep ensure -vendor-only

client:
	go build -o client

install: dep client
	go install 
