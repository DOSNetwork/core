#!/bin/bash

setting="./dos.setting"
LABEL="dosnode"
# If file exists
if [[ -f "$setting" ]]; then
  source $setting
  if [ -z "$DOSIMAGE" ]; then
    echo "Please assign a value to DOSIMAGE field in dos.setting"
    exit
  fi
  if [ -z "$IP" ]; then
    echo "Please assign a value to IP field in dos.setting"
    exit
  fi
  if [ -z "$KEYSTORE" ]; then
    echo "Please assign a value to KEYSTORE in dos.setting"
    exit
  fi
  if [ -z "$GETHPOOL" ]; then
    echo "Please assign a value to GETHPOOL in dos.setting"
    exit
  fi
else
  echo "Can't find dos.setting"
  exit
fi

if [ "$USER" == "root" ]; then
  DIR="/root"
else
  DIR="/home/$USER"
fi

install_lightnode(){
  yes | sudo apt-get update
  yes | sudo add-apt-repository -y ppa:ethereum/ethereum
  yes | sudo apt-get update
  yes | sudo apt-get install ethereum
  wget https://www.rinkeby.io/rinkeby.json -q -O rinkeby.json
  sudo geth --datadir='$DIR'/.rinkeby init rinkeby.json
  sudo cp static-nodes.json $DIR/.rinkeby/geth/static-nodes.json
#  rm -f rinkeby.json
  sudo sed -i 's|xxx|'$DIR'|g' geth.service
  yes | sudo chmod 322 geth.service
  yes | sudo cp geth.service /lib/systemd/system/geth.service
  yes | sudo systemctl enable geth
  yes | sudo systemctl start geth
}

start_lightnode(){
  echo "start_lightnode"
  ps cax | grep geth > /dev/null
  if [ $? -eq 0 ]; then
    echo "Geth Process is running."
  else
    yes | sudo systemctl start geth
  fi
}

install_docker(){
  yes | sudo apt-get update
  yes | sudo apt-get install docker.io
  sudo groupadd docker
  sudo usermod -aG docker $USER
}

install_dos(){
  mkdir -p $DIR/dos
  mkdir -p $DIR/credential
  echo $KEYSTORE
  cp $KEYSTORE $DIR/credential/
}

run(){
  result=$(docker container ls | awk '(index($2, "dos") != 0) {print $2}')
  newlog="doslog_$IP"
  if [ -z "$result" ]; then
    sudo rm dos/doslog
    docker pull $DOSIMAGE
    echo -n Password:;read -s password ;
    docker run -l $LABEL -it -d \
      -p 7946:7946 \
      -p 8080:8080 \
      -p 9501:9501 \
      --mount type=bind,source=$DIR/credential,target=/credential  \
      --mount type=bind,source=$DIR/dos,target=/dos  \
      -e PUBLICIP=$IP \
      -e GETHPOOL=$GETHPOOL \
      -e PASSPHRASE=$password  \
      -e CHAINNODE=rinkeby  \
      -e APPSESSION=$DOSVERSION \
      -e APPNAME=DosClient  \
      $DOSIMAGE
  else
    echo "client is running";
  fi
}

stopDos(){
  result=$(docker container ls | awk '(index($2, "dos") != 0) {print $2}')
  newlog="doslog_$IP"
  if [ -z "$result" ]; then echo "client is not running"
  else
    docker stop $(docker ps -aqf "label=$LABEL")
    docker rm $(docker ps -aqf "label=$LABEL")
  fi
}

check(){
  result=$(docker container ls | awk '(index($2, "dos") != 0) {print $2}')
  newlog="doslog_$IP"
  if [ -z "$result" ]; then
    echo "client is not running";
  else
    echo "client is running";
  fi
}

proxyInfo(){
  curl http://localhost:8080/proxy
}

clientInfo(){
  curl http://localhost:8080/status
}

guardian(){
  curl http://localhost:8080/guardian
}

case "$1" in
  "install")
    install_lightnode
    install_docker
    install_dos
    echo "Please log out and log back to finish installation"
    ;;
  "startLightnode")
    start_lightnode
    ;;
  "run")
    run
    ;;
  "stop")
    stopDos
    ;;
  "check")
    check
    ;;
  "proxyInfo")
    proxyInfo
    ;;
  "clientInfo")
    clientInfo
    ;;
  "guardian")
    guardian
    ;;
  *)
    echo "Usage: bash vps_docker.sh [OPTION]"
    echo "OPTION:"
    echo "  install       Install Docker and setup directory for client"
    echo "  run           Run the client from Docker Hub"
    echo "  stop          Stop the client"
    echo "  check         Check to see if client is running; Download log if client is not running"
    echo "  proxyInfo     Print proxy information"
    echo "  clientInfo    Print client information"
    echo "  guardian      Trigger guardian"
    ;;
esac
