#/bin/bash
USER="ubuntu"
VPSIP="52.12.7.248"
VPSKEY="LightsailDefaultKey-us-west-2.pem"
KEYPATH="credential/usrKey"

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

install_client(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'yes | [ -f "client" ] && rm client'
  scp -i $VPSKEY client  $USER@$VPSIP:~/
  update_config
}

update_config(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'mkdir -p dos'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'mkdir -p credential'
	scp -i $VPSKEY .dosenv $USER@$VPSIP:~/
	scp -i $VPSKEY config.json $USER@$VPSIP:~/
	scp -i $VPSKEY $KEYPATH $USER@$VPSIP:~/credential/
}

run_client(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'source ~/.dosenv; echo -n Password:;read -s password ;export PASSPHRASE=$password ;setsid ./client start >dos/dos.log 2>&1'
}

stop_client(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP './client stop'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP '[ -f "client.pid" ] && rm client.pid'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP 'sudo pkill -f client'
}

show_proxy(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i $VPSKEY $USER@$VPSIP './client cmd showProxyStatus'
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
  "install_lightnode")
    install_lightnode
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
    echo "Usage: bash vps.sh [OPTION]"
	echo "OPTION:"
	echo "  install_lightnode"
	echo "  install_client"
	echo "  run_client"
	echo "  stop_client"
	echo "  show_proxy"
	echo "  trigger_guardian"
	echo "  check_client"
	echo "  update_config"
    exit 1
    ;;
esac
