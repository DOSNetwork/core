#/bin/bash

setting="./dos.setting"
# If file exists 
if [[ -f "$setting" ]]
then
    source $setting
	if [ -z "$DOSIMAGE" ]; 
	then 
		echo "Please assign a value to DOSIMAGE in dos.setting"
		exit
	fi
	if [ -z "$USER" ]; 
	then 
		echo "Please assign a value to USER in dos.setting"
		exit
	fi
	if [ -z "$VPSIP" ]; 
	then 
		echo "Please assign a value to VPSIP in dos.setting"
		exit
	fi
	if [ -z "$VPSKEY" ]; 
	then 
		echo "Please assign a value to VPSKEY in dos.setting"
		exit
	fi
	if [ -z "$KEYSTORE" ]; 
	then 
		echo "Please assign a value to KEYSTORE in dos.setting"
		exit
	fi
	if [ -z "$GETHPOOL" ]; 
	then 
		echo "Please assign a value to GETHPOOL in dos.setting"
		exit
	fi
else
    echo "Can't find dos.setting"
	exit
fi

installAll(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo apt-get update'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo apt-get install docker.io'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'sudo usermod -a -G docker $USER'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'mkdir -p dos'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'mkdir -p credential'
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY $KEYSTORE $USER@$VPSIP:~/credential/
}

run(){
	result=$(ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP docker container ls | awk '(index($2, "dos") != 0) {print $2}')
	newlog="doslog_$VPSIP"
	if [ -z "$result" ]; 
	then
    DIR="/home/"$USER
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'sudo rm dos/doslog'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'docker pull '$DOSIMAGE
	echo -n Password:;read -s password ;
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'docker run -it -d \
-p 7946:7946 \
-p 8080:8080 \
-p 9501:9501 \
--mount type=bind,source='$DIR'/credential,target=/credential  \
--mount type=bind,source='$DIR'/dos,target=/dos  \
-e PUBLICIP="'$VPSIP'" \
-e GETHPOOL="'$GETHPOOL'" \
-e PASSPHRASE='$password'  \
-e CHAINNODE=rinkeby  \
-e APPSESSION="'$DOSVERSION'" \
-e APPNAME=DosClient  \
'$DOSIMAGE
	else echo "client is running";
	fi
}

stop(){
	result=$(ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP docker container ls | awk '(index($2, "dos") != 0) {print $2}')
	newlog="doslog_$VPSIP"
	if [ -z "$result" ];
	then echo "client is not running"
	else
	    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP  'docker stop $(docker ps -a -q)'
		ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP  'docker rm $(docker ps -a -q)'
	fi
}

check(){
	result=$(ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP docker container ls | awk '(index($2, "dos") != 0) {print $2}')
	newlog="doslog_$VPSIP"
	if [ -z "$result" ];
	then scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY $USER@$VPSIP:~/dos/doslog $newlog;
	else echo "client is running";
	fi
}

proxyInfo(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'curl http://localhost:8080/proxy'
}

clientInfo(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'curl http://localhost:8080/status'
}

guardian(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'curl http://localhost:8080/guardian'
}

case "$1" in
  "install")
    installAll
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
    exit 1
    ;;
esac
