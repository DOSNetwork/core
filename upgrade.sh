#!/bin/bash

export DOSTAG="v1.0.2-m"

sudo rm -f vault/dosclient.pid
sudo docker stop $(docker ps -a -q)
sudo docker rm -f $(docker ps -a -q)
sudo rm -f vault/doslog.txt
sudo rm -f vault/doslog.txt.*

docker pull dosnetwork/dosnode:'$DOSTAG'
docker run -it -d -p 7946:7946 -p 8080:8080 -p 9501:9501 --mount type=bind,source=/home/"$USER"/,target=/config --mount type=bind,source=/home/"$USER"/vault,target=/vault --hostname dos --name dosclient --env CONFIGPATH=config --env PASSPHRASE="$PASSWORD" --env APPSESSION="$DOSTAG-mainnet" --env LOG_LEVEL=debug dosnetwork/dosnode:"$DOSTAG" /dosclient start
