#/bin/bash
USER=""
VPSIP=""
VPSKEY=""
KEYPATH=""

install_lightnode(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo apt-get update'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo add-apt-repository -y ppa:ethereum/ethereum'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo apt-get update'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo apt-get install ethereum'
  scp -i $VPSKEY geth.service  $USER@$VPSIP:~/
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo chmod 322 geth.service'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo mv geth.service /lib/systemd/system/geth.service'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo systemctl enable geth'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo systemctl start geth'
}

install_docker(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo apt-get install docker.io'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'sudo usermod -a -G docker $USER'
}

install_client(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'mkdir -p dos'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'mkdir -p credential'
	scp -i $VPSKEY $KEYPATH $USER@$VPSIP:~/credential/
}

run_client(){
    DIR="/home/"$USER
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'sudo rm dos/doslog'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'docker pull dosnetwork/dosnode:latest'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'docker run -it -d \
-p 7946:7946 \
-p 8080:8080 \
-p 9501:9501 \
-p 80:80 \
--mount type=bind,source='$DIR'/credential,target=/credential  \
--mount type=bind,source='$DIR'/dos,target=/dos  \
-e LOGIP=163.172.36.173:9500  \
-e CHAINNODE=rinkebyPrivateNode  \
-e PASSPHRASE=123  \
-e APPSESSION=Beta  \
-e APPNAME=DosNode  \
dosnetwork/dosnode:latest'
}

stop_client(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP  'docker stop $(docker ps -a -q)'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP  'docker rm $(docker ps -a -q)'
}

show_proxy(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'curl http://localhost:8080/proxy'
}

trigger_guardian(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'curl http://localhost:8080/guardian'
}

check_client(){
	result=$(ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP docker container ls | awk '(index($2, "dos") != 0) {print $2}')
	newlog="doslog_$VPSIP"
	if [ -z "$result" ]; 
	then scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY $USER@$VPSIP:~/dos/doslog $newlog;
	else echo "client is running";
	fi
}

case "$1" in
  "install_lightnode")
    install_lightnode
    ;;
  "install_docker")
    install_docker
    ;;
  "install_client")
    install_client
    ;;
  "update_config")
    update_config
    ;;
  "run_client")
    run_client
    ;;
  "stop_client")
    stop_client
    ;;
  "show_proxy")
    show_proxy
    ;;
  "trigger_guardian")
    trigger_guardian
    ;;
  "check_client")
    check_client
    ;;
  *)
    echo "Usage: bash vps_docker.sh [OPTION]"
	echo "OPTION:"
	echo "  install_lightnode"
	echo "  install_docker"
	echo "  install_client"
	echo "  run_client"
	echo "  stop_client"
	echo "  show_proxy"
	echo "  trigger_guardian"
	echo "  check_client"
    exit 1
    ;;
esac
