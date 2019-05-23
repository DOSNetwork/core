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

if [ "$USER" == "root" ]; then
  DIR="/root"
else
  DIR="/home/$USER"
fi

install_lightnode(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'ps cax | grep geth > /dev/null'
  if [ $? -eq 0 ]; then
    echo "Geth Process is running."
  else
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
#    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo systemctl enable geth'
#    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | sudo systemctl start geth'
  fi
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

install_client(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'yes | [ -f "client" ] && rm client'
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $SSHKEY client  $USER@$IP:$DIR/
  update_config
}

update_config(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'mkdir -p '$DIR'/dos'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'mkdir -p '$DIR'/credential'
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $SSHKEY config.json $USER@$IP:~/
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $SSHKEY $KEYSTORE $USER@$IP:$DIR'/credential/'
}

run_client(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'export GETHPOOL="'$GETHPOOL'";export PUBLICIP='$IP'; echo -n Password:;read -s password ;export PASSPHRASE=$password ;setsid ./client start >dos/doslog 2>&1'
}

stop_client(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP './client stop'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP '[ -f "client.pid" ] && rm client.pid'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'sudo pkill -f client'
}

show_proxy(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP './client cmd showProxyStatus'
}

clientInfo(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP './client cmd showStatus'
}

trigger_guardian(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP './client cmd triggerGuardian'
}

check_client(){
  result=$(ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $SSHKEY $USER@$IP 'pgrep -x client')
  newlog="doslog_$IP"
  if [ -z "$result" ]; then
    scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $SSHKEY $USER@$IP:~/dos/doslog $newlog;
  else
    echo "client is running";
  fi
}

case "$1" in
  "install")
    install_lightnode
    install_client
    ;;
  "run")
    run_client
    ;;
  "stop")
    stop_client
    ;;
  "check")
    check_client
    ;;
  "proxyInfo")
    show_proxy
    ;;
  "clientInfo")
    clientInfo
    ;;
  "guardian")
    guardian
    ;;
  *)
    echo "Usage: bash vps.sh [OPTION]"
    echo "OPTION:"
    echo "  install       Install lightnode and dos client"
    echo "  run           Run the client"
    echo "  stop          Stop the client"
    echo "  check         Check to see if client is running; Download log if client is not running"
    echo "  proxyInfo     Print proxy information"
    echo "  clientInfo    Print client information"
    echo "  guardian      Trigger guardian"
    ;;
esac
