#!/bin/bash
install_docker() {
  if [ ! -x "$(command -v docker)" ]; then
    echo "Install docker"
    yes | sudo dpkg --configure -a
    yes | sudo apt-get update
    yes | sudo apt-get install docker.io
    sudo usermod -aG docker $USER
    echo "Rebooting!!!"
    yes | sudo reboot
  fi
}

uninstall_docker() {
  if [ -x "$(command -v docker)" ]; then
    sudo apt-get purge -y docker-engine docker docker.io docker-ce
    sudo apt-get autoremove -y --purge docker-engine docker docker.io docker-ce
    sudo rm -rf /var/lib/docker /etc/docker
    sudo rm /etc/apparmor.d/docker
    sudo groupdel docker
    sudo rm -rf /var/run/docker.sock
  fi
}

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
  install_docker
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
  echo "4) Pull dosnetwork/dosnode:beta"
  docker pull dosnetwork/dosnode:beta
  echo "5) Check if node keystore exists"
  docker run -it --mount type=bind,source=$(pwd)/vault,target=/vault \
    dosnetwork/dosnode:beta /dosclient wallet create
  echo -n Keystore Password [Enter]:
  read -s password
  docker run -it -d -p 7946:7946 -p 8080:8080 -p 9501:9501 \
    --hostname $(hostname) \
    --mount type=bind,source=$(pwd),target=/config \
    --mount type=bind,source=$(pwd)/vault,target=/vault \
    -e CONFIGPATH=config -e PASSPHRASE=$password \
    dosnetwork/dosnode:beta /dosclient start
}

status() {
  curl http://localhost:8080/
}

stop() {
  docker stop $(docker ps -a -q)
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
  echo "Usage: bash dos_docker.sh [OPTION]"
  echo "OPTION:"
  echo "  start         Start dosclient from Docker Hub"
  echo "  stop          Stop dosclient"
  echo "  status        Show dosclient status"
  echo "  log           Show dosclient log"
  ;;
esac
