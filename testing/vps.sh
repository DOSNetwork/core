#/bin/bash
IFS=$'\r\n' GLOBIGNORE='*' command eval  'ips=($(cat nodeIP.config))'
ipslength=${#ips[@]}

install_lightnode(){
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

install_dos(){
  ssh -tt -i "" ubuntu@${ips[$i]} '[ -f "dosclient" ] && rm dosclient'
  scp -i "" dosclient  ubuntu@${ips[$i]}:~/
  update_dos_config
}

update_dos_config(){
	scp -i "" .env_profile ubuntu@${ips[$i]}:~/
	ssh -tt -i "" ubuntu@${ips[$i]} 'source ~/.env_profile'
	scp -i "" config.json  ubuntu@${ips[$i]}:~/
	ssh -tt -i "" ubuntu@${ips[$i]} 'mkdir -p credential'
	scp -i "" testAccounts/$(($i + 1))/credential/usrKey ubuntu@${ips[$i]}:~/credential/
}

run_dos(){
	#ssh tt -i "" ubuntu@${ips[$i]} 'source ~/.env_profile; echo -n Password:;read -s password ;export PASSPHRASE=$password ;setsid ./dosclient start >dos.log 2>&1'
	ssh -tt -i "" ubuntu@${ips[$i]} 'source ~/.env_profile; export PASSPHRASE= ;setsid ./dosclient start >dos.log 2>&1'
}

dos_log(){
	ssh -tt -i "" ubuntu@${ips[$i]} 'cat dos.log'
}

stop_dos(){
	ssh -tt -i "" ubuntu@${ips[$i]} './dosclient stop'
	ssh -tt -i "" ubuntu@${ips[$i]} '[ -f "dosclient.pid" ] && rm dosclient.pid'
	ssh -tt -i "" ubuntu@${ips[$i]} 'pkill -f dosclient'
}

test_p2p(){
	ssh -tt -i "" ubuntu@${ips[$i]} './dosclient cmd testP2P'
}

case "$1" in
  "init_ssh")
    for (( i=0; i<${ipslength}; i++ ));
    do
      ssh -tt -i "" ubuntu@${ips[$i]} 'yes | sudo ls'
    done
    ;;
  "install_lightnode")
    for (( i=0; i<${ipslength}; i++ ));
    do
        install_lightnode $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "install_dos")
    for (( i=0; i<${ipslength}; i++ ));
    do
        install_dos $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "update_dos_config")
    for (( i=0; i<${ipslength}; i++ ));
    do
        update_dos_config $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "run_dos")
    for (( i=0; i<9; i++ ));
    do
        run_dos $i & # Put a function in the background
    done
    wait
    echo "All done"
    #sleep 600
    for (( i=9; i<${ipslength}; i++ ));
    do
        run_dos $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "stop_dos")
    for (( i=0; i<${ipslength}; i++ ));
    do
        stop_dos $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "test_p2p")
    for (( i=0; i<${ipslength}; i++ ));
    do
        test_p2p $i
    done
    echo "All done"
    ;;
  *)
    echo "You have failed to specify what to do correctly."
    exit 1
    ;;
esac

