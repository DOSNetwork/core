install_docker() {
  if [ ! -x "$(command -v docker)" ]; then
    echo "Install docker"
    sudo dpkg --configure -a
    yes | sudo apt-get update
    yes | sudo apt-get install docker.io
    sudo usermod -aG docker $USER
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

uninstall_LightClient() {
  if [ -f /lib/systemd/system/geth.service ]; then
    yes | sudo systemctl stop geth
    yes | sudo systemctl disable geth
    yes | sudo rm /lib/systemd/system/geth.service
  fi
}

install_LightClient() {
  if [ ! -f $(pwd)/geth-linux-amd64-1.9.2-e76047e9/geth ]; then
    wget https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.9.2-e76047e9.tar.gz
    tar -C . -xzf geth-linux-amd64-1.9.2-e76047e9.tar.gz
    yes | sudo cp $(pwd)/geth-linux-amd64-1.9.2-e76047e9/geth /usr/bin/geth
    rm $(pwd)/geth-linux-amd64-1.9.2-e76047e9.tar.gz
  fi
  yes | sudo cp $(pwd)/geth-linux-amd64-1.9.2-e76047e9/geth /usr/bin/geth
  if [ ! -f /lib/systemd/system/geth.service ]; then
    yes | sudo chmod 666 geth.service.tmpl
    sed -i -e 's?#REPLACE-WITH-PATH?'$(pwd)'?g' geth.service.tmpl
    sed -i -e 's?#REPLACE-WITH-IP?'$nodeIP'?g' geth.service.tmpl
    yes | sudo chmod 322 geth.service.tmpl
    yes | sudo cp geth.service.tmpl /lib/systemd/system/geth.service
    yes | sudo systemctl enable geth
    yes | sudo systemctl start geth
  fi
}

findNodeIP() {
  if [ -z "$nodeIP" ]; then
    nodeIP=$(host myip.opendns.com resolver1.opendns.com | grep "myip.opendns.com has address" | awk -F " " '{print $NF}')
    echo "Is your node public IP :"$nodeIP
    echo -n "Enter no if it is not correct:"
    read reply
    if [ "$reply" == "no" ]; then
      echo "Please provide node public IP. ex:dos.sh install IP"
      exit 1
    fi
    echo $reply
  fi
  echo "node public IP : " $nodeIP
}

checkConfigs() {
  mkdir -p $(pwd)/vault
  if [ ! -f $(pwd)/config.json ]; then
    echo "config.json not found!"
    exit 1
  fi
  if [ ! -f $(pwd)/config.toml ]; then
    echo "config.toml not found!"
    exit 1
  fi
  if [ ! -f $(pwd)/geth.service.tmpl ]; then
    echo "geth.service not found!"
    exit 1
  fi
}

installAll() {
  uninstall_docker
  uninstall_LightClient
  checkConfigs
  findNodeIP
  install_docker
  install_LightClient
}

startClient() {
  if grep -q "#REPLACE-WITH-IP" $(pwd)/config.json; then
    findNodeIP
    echo "Is your node public IP :"$nodeIP
    echo -n "Enter no if it is not correct:"
    if [ "$reply" == "no" ]; then
      echo "Please provide node public IP in config.json ."
      exit 1
    fi
    sed -i -e 's?#REPLACE-WITH-IP?'$nodeIP'?g' $(pwd)/config.json
  fi
  docker pull dosnetwork/dosnode:beta
  docker run -it --mount type=bind,source=$(pwd)/vault,target=/vault \
    dosnetwork/dosnode:beta /client wallet create
  echo -n Enter password :
  read -s password
  docker run -it -d -p 7946:7946 -p 8080:8080 -p 9501:9501 \
    --mount type=bind,source=$(pwd),target=/config \
    --mount type=bind,source=$(pwd)/vault,target=/vault \
    -e CONFIGPATH=config -e PASSPHRASE=$password \
    dosnetwork/dosnode:beta /client start
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
"install")
  nodeIP=$2
  installAll
  ;;
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
  echo "  install       install docker,geth and client"
  echo "  start         Start client from Docker Hub"
  echo "  stop          Stop client"
  echo "  status        Show client status"
  echo "  log           Show client log"
  ;;
esac
