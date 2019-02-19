#!/bin/sh

docker build -t gnte .
go get github.com/karalabe/xgo
xgo --targets=linux/amd64 -out scripts/dosclient ../../
xgo --targets=linux/amd64 -out scripts/dosuser ../dosUser/
xgo --targets=linux/amd64 -out scripts/dosboot ../bootStrapNode/
Chmod 777 scripts/dosclient-linux-amd64
Chmod 777 scripts/dosuser-linux-amd64
Chmod 777 scripts/dosboot-linux-amd64
cp -r ../../testAccounts scripts/
cp ../dosUser/ama.json scripts/
cp ../../config.json scripts/
