#!/bin/bash
RELEASE="v1.0-beta.18"
findNodeIP() {
  if [ -z "$nodeIP" ]; then
    nodeIP=$(host myip.opendns.com resolver1.opendns.com | grep "myip.opendns.com has address" | awk -F " " '{print $NF}')
    echo "node public IP :"$nodeIP
    echo -n "Press [Enter] to continue or reply no if IP is not correct:"
    read reply
    if [ "$reply" == "no" ]; then
      echo -n "Enter your public IP [Enter]: "
      read nodeIP
    fi
  fi
}

startClient() {
  mkdir -p $(pwd)/vault
  echo "1) Check if dosclient exists"
  if [ ! -f $(pwd)/dosclient ]; then
    URL="https://github.com/DOSNetwork/core/releases/download/"$RELEASE"/dosclient"
    echo "-->Downloading dosnode from "$URL
    wget $URL
    chmod +x dosclient
  fi
  echo "2) Check if config.json includes Infura API key"
  if grep -q "#REPLACE-WITH-INFURA-APIKEY" $(pwd)/config.json; then
    echo -n "Enter your infura api key [Enter]: "
    read apikey
    sed -i -e 's?#REPLACE-WITH-INFURA-APIKEY?'$apikey'?g' $(pwd)/config.json
  fi
  echo "3) Check if config.json includes a node public IP"
  if grep -q "#REPLACE-WITH-IP" $(pwd)/config.json; then
    findNodeIP
    sed -i -e 's?#REPLACE-WITH-IP?'$nodeIP'?g' $(pwd)/config.json
  fi
  echo "4) Check if node keystore exists"
  $(pwd)/dosclient wallet create
  echo -n Keystore Password [Enter]:
  read -s password
  echo ""
  export PASSPHRASE=$password
  echo "5) Run dosclient"
  nohup ./dosclient start >dos.out 2>&1 &
  sleep 3
  pgrep -x dosclient >/dev/null && echo "dosclient is running" || echo "dosclient is not running"
}

status() {
  curl http://localhost:8080/
  if [ "$?" = "7" ]; then
    log
  fi
}

stop() {
  ./dosclient stop
}

log() {
  #cat $(pwd)/vault/doslog.txt
  result="$(
    curl http://localhost:80/api/explorer/search?text= &
    pageSize=20 &
    pageIndex=0
  )"
  if [ "$?" = "7" ]; then
    echo "fail"
  fi
  echo "result: "$result
}

case "$1" in
"start")
  startClient
  ;;
"stop")
  stop
  ;;
"status")
  status
  ;;
"log")
  log
  ;;
*)
  echo "Usage: bash dos.sh [OPTION]"
  echo "OPTION:"
  echo "  start         Start dosclient from Docker Hub"
  echo "  stop          Stop dosclient"
  echo "  status        Show dosclient status"
  echo "  log           Show dosclient log"
  ;;
esac
