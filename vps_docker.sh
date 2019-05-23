#!/bin/bash

setting="./dos.setting"
# If file exists
if [[ -f "$setting" ]]; then
  source $setting
  if [ -z "$DOSIMAGE" ]; then
    echo "Please assign a value to DOSIMAGE field in dos.setting"
    exit
  fi
  if [ -z "$USER" ]; then
    echo "Please assign a value to USER field in dos.setting"
    exit
  fi
  if [ -z "$IP" ]; then
    echo "Please assign a value to IP field in dos.setting"
    exit
  fi
  if [ -z "$SSHKEY" ]; then
    echo "Please assign a value to SSHKEY field in dos.setting"
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

GETHPOOL="$GETHPOOL;ws://$IP:8546"
if [ "$USER" == "root" ]; then
  DIR="/root"
else
  DIR="/home/$USER"
fi

install_lightnode(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo apt-get update'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo add-apt-repository -y ppa:ethereum/ethereum'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo apt-get update'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo apt-get install ethereum'
  wget https://www.rinkeby.io/rinkeby.json -q -O rinkeby.json
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $SSHKEY rinkeby.json $USER@$IP:~/
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'geth --datadir='$DIR'/.rinkeby init rinkeby.json'
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $SSHKEY static-nodes.json $USER@$IP:$DIR/.rinkeby/geth/static-nodes.json
  rm -f rinkeby.json
  sed -i 's|xxx|'$DIR'|g' geth.service
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $SSHKEY geth.service  $USER@$IP:~/
  git checkout -- geth.service
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo chmod 322 geth.service'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo mv geth.service /lib/systemd/system/geth.service'
#  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo systemctl enable geth'
#  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo systemctl start geth'
}

start_lightnode(){
  echo "start_lightnode"
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'ps cax | grep geth > /dev/null'
  if [ $? -eq 0 ]; then
    echo "Geth Process is running."
  else
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo systemctl start geth'
  fi
}

install_docker(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo apt-get update'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo apt-get install docker.io'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'sudo usermod -a -G docker $USER'
}

install_dos(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'mkdir -p '$DIR'/dos'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'mkdir -p '$DIR'/credential'
  echo $KEYSTORE
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $SSHKEY $KEYSTORE $USER@$IP:$DIR'/credential/'
}

run(){
  result=$(ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP docker container ls | awk '(index($2, "dos") != 0) {print $2}')
  newlog="doslog_$IP"
  if [ -z "$result" ]; then
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'sudo rm dos/doslog'
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'docker pull '$DOSIMAGE
    echo -n Password:;read -s password ;
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'docker run -it -d \
      -p 7946:7946 \
      -p 8080:8080 \
      -p 9501:9501 \
      --mount type=bind,source='$DIR'/credential,target=/credential  \
      --mount type=bind,source='$DIR'/dos,target=/dos  \
      -e PUBLICIP="'$IP'" \
      -e GETHPOOL="'$GETHPOOL'" \
      -e PASSPHRASE='$password'  \
      -e CHAINNODE=rinkeby  \
      -e APPSESSION="'$DOSVERSION'" \
      -e APPNAME=DosClient  \
      '$DOSIMAGE
  else
    echo "client is running";
  fi
}

stop(){
  result=$(ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP docker container ls | awk '(index($2, "dos") != 0) {print $2}')
  newlog="doslog_$IP"
  if [ -z "$result" ]; then echo "client is not running"
  else
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP  'docker stop $(docker ps -a -q)'
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP  'docker rm $(docker ps -a -q)'
  fi
}

check(){
  result=$(ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP docker container ls | awk '(index($2, "dos") != 0) {print $2}')
  newlog="doslog_$IP"
  if [ -z "$result" ]; then
    scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $SSHKEY $USER@$IP:~/dos/doslog $newlog;
  else
    echo "client is running";
  fi
}

proxyInfo(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'curl http://localhost:8080/proxy'
}

clientInfo(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'curl http://localhost:8080/status'
}

guardian(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'curl http://localhost:8080/guardian'
}

case "$1" in
  "install")
    install_lightnode
    install_docker
    install_dos
    ;;
  "run")
    run
    ;;
  "stop")
    stop
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
