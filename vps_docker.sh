#!/bin/bash
mkdir -p $(pwd)/vault
startClient(){
  docker run -it --mount type=bind,source=$(pwd)/vault,target=/vault \
          dosnode:latest /client wallet create
  echo -n Enter password :
  read -s password
  docker run -it -d -p 7946:7946 -p 8080:8080 -p 9501:9501 \
          --mount type=bind,source=$(pwd),target=/config  \
          --mount type=bind,source=$(pwd)/vault,target=/vault  \
          -e CONFIGPATH=config -e PASSPHRASE=$password \
          dosnode:latest /client start
}

status(){
  curl http://localhost:8080/
}

stop(){
  docker stop $(docker ps -a -q)
}

case "$1" in
  "start")
    startClient
    ;;
  "status")
    status
    ;;
  "stop")
    stop
    ;;
  *)
    echo "Usage: bash vps_docker.sh [OPTION]"
    echo "OPTION:"
    echo "  run           Run the client from Docker Hub"
    echo "  stop          Stop the client"
    ;;
esac
