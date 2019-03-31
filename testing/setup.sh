#/bin/bash
IFS=$'\r\n' GLOBIGNORE='*' command eval  'ips=($(cat betips))'
ipslength=${#ips[@]}
setuplightnode_func(){
  echo ubuntu@[${ips[$i]}]
  ssh -tt -i "" ubuntu@${ips[$i]} 'sudo apt-get update'
  ssh -tt -i "" ubuntu@${ips[$i]} 'sudo add-apt-repository -y ppa:ethereum/ethereum'
  ssh -tt -i "" ubuntu@${ips[$i]} 'sudo apt-get update'
  ssh -tt -i "" ubuntu@${ips[$i]} 'yes | sudo apt-get install ethereum'
  scp -i "" geth.service  ubuntu@${ips[$i]}:~/
  ssh -tt -i "" ubuntu@${ips[$i]} 'sudo chmod 322 geth.service'
  ssh -tt -i "" ubuntu@${ips[$i]} 'sudo mv geth.service /lib/systemd/system/geth.service'
  ssh -tt -i "" ubuntu@${ips[$i]} 'sudo systemctl enable geth'
  ssh -tt -i "" ubuntu@${ips[$i]} 'sudo systemctl start geth'
}

installClient_func(){
  ssh -tt -i "" ubuntu@${ips[$i]} '[ -f "dosclient" ] && rm dosclient'
  scp -i "" dosclient  ubuntu@${ips[$i]}:~/
  updateClient_func
}

updateClient_func(){
  scp -i "" .env_profile ubuntu@${ips[$i]}:~/
  ssh -tt -i "" ubuntu@${ips[$i]} 'source ~/.env_profile'
  scp -i "" config.json  ubuntu@${ips[$i]}:~/
  ssh -tt -i "" ubuntu@${ips[$i]} 'mkdir -p credential'
  scp -i "" testAccounts/$(($i + 1))/credential/usrKey ubuntu@${ips[$i]}:~/credential/
}

runClient_func(){
  #ssh tt -i "" ubuntu@${ips[$i]} 'source ~/.env_profile; echo -n Password:;read -s password ;export PASSPHRASE=$password ;setsid ./dosclient start'
  ssh -tt -i "" ubuntu@${ips[$i]} 'source ~/.env_profile; export PASSPHRASE= ;setsid ./dosclient start'
}

stopClient_func(){
  ssh -tt -i "" ubuntu@${ips[$i]} './dosclient stop'
  ssh -tt -i "" ubuntu@${ips[$i]} '[ -f "dosclient.pid" ] && rm dosclient.pid'
}
case "$1" in
  "InitSSH")
    for (( i=0; i<${ipslength}; i++ ));
    do
      ssh -tt -i "" ubuntu@${ips[$i]} 'yes | sudo ls'
    done
    ;;
  "-lightnode")
    for (( i=0; i<${ipslength}; i++ ));
    do
        setuplightnode_func $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "-installclient")
    for (( i=0; i<${ipslength}; i++ ));
    do
        installClient_func $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "-updateclient")
    for (( i=0; i<${ipslength}; i++ ));
    do
        updateClient_func $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "-runclient")
    for (( i=0; i<${ipslength}; i++ ));
    do
        runClient_func $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "-stopclient")
    for (( i=0; i<${ipslength}; i++ ));
    do
        stopClient_func $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  *)
    echo "You have failed to specify what to do correctly."
    exit 1
    ;;
esac

