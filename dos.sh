#!/bin/bash
RELEASE="https://github.com/DOSNetwork/core/releases/download/Beta1.2/client"
findNodeIP() {
  if [ -z "$nodeIP" ]; then
    nodeIP=$(host myip.opendns.com resolver1.opendns.com | grep "myip.opendns.com has address" | awk -F " " '{print $NF}')
    echo "node public IP :"$nodeIP
    echo -n "Press [Enter] to continue or no if IP is not correct:"
    read reply
    if [ "$reply" == "no" ]; then
      echo -n "Enter your public IP [Enter]: "
      read nodeIP
    fi
  fi
}

startClient() {
  echo "1) Check if client exists"
  if [ ! -f $(pwd)/dosnode ]; then
    echo "-->Downloading dosnode from "$RELEASE
    wget $RELEASE
    chmod 111 dosnode
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
  $(pwd)/dosnode wallet create
  echo -n Keystore Password [Enter]:
  read -s password
  echo ""
  export PASSPHRASE=$password
  echo "5) Run dosnode"
  nohup ./dosnode start > dos.out 2>&1 &
  sleep 3
  pgrep -x dosnode >/dev/null && echo "dosnode is running" || echo "dosnode is not running"
}

status() {
  ./dosnode status
}

stop() {
  ./dosnode stop
}

log() {
  cat $(pwd)/vault/doslog.txt
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
  echo "  start         Start dosnode from Docker Hub"
  echo "  stop          Stop dosnode"
  echo "  status        Show dosnode status"
  echo "  log           Show dosnode log"
  ;;
esac
