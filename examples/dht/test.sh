#!/usr/bin/env bash

portNumber=57909

go run dhtmain.go -port=$portNumber -target &

while(($portNumber<57929))
do
    sleep 1
    let "portNumber++"
    go run dhtmain.go -port=$portNumber -address=127.0.0.1:57909 -target &
done

sleep 5
go run dhtmain.go -address=127.0.0.1:57909 -find &

read

killall -9 dhtmain