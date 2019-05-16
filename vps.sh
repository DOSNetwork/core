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
GETHPOOL=$GETHPOOL";ws://"$VPSIP":8546"
DIR="/home/"$USER

install_lightnode(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'ps cax | grep geth > /dev/null'
  if [ $? -eq 0 ]
  then
	echo "Geth Process is running."
  else
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo apt-get update'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo add-apt-repository -y ppa:ethereum/ethereum'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo apt-get update'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo apt-get install ethereum'
	scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY rinkeby.json $USER@$VPSIP:~/
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'geth --datadir='$DIR'/.rinkeby init rinkeby.json'
	scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY static-nodes.json $USER@$VPSIP:$DIR/.rinkeby/geth/static-nodes.json
	scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY geth.service  $USER@$VPSIP:~/
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo chmod 322 geth.service'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo mv geth.service /lib/systemd/system/geth.service'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo systemctl enable geth'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo systemctl start geth'
  fi
}

start_lightnode(){
  echo "start_lightnode"
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'ps cax | grep geth > /dev/null'
  if [ $? -eq 0 ]
  then
	echo "Geth Process is running."
  else
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | sudo systemctl start geth'
  fi
}

install_client(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | [ -f "client" ] && rm client'
	scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY client  $USER@$VPSIP:$DIR/
	update_config
}

update_config(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'mkdir -p '$DIR'/dos'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'mkdir -p '$DIR'/credential'
	scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY config.json $USER@$VPSIP:~/
	scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY $KEYSTORE $USER@$VPSIP:$DIR'/credential/'
}

run_client(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'export GETHPOOL="'$GETHPOOL'";export PUBLICIP='$VPSIP'; echo -n Password:;read -s password ;export PASSPHRASE=$password ;setsid ./client start >dos/doslog 2>&1'
}

stop_client(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP './client stop'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP '[ -f "client.pid" ] && rm client.pid'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'sudo pkill -f client'
}

show_proxy(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP './client cmd showProxyStatus'
}

clientInfo(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP './client cmd showStatus'
}

trigger_guardian(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP './client cmd triggerGuardian'
}

check_client(){
    result=$(ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'pgrep -x client')
    newlog="doslog_$VPSIP"
    if [ -z "$result" ];
    then scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i $VPSKEY $USER@$VPSIP:~/dos/doslog $newlog;
    else echo "client is running";
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
    exit 1
    ;;
esac

