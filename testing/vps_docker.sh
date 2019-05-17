#/bin/bash
IFS=$'\r\n' GLOBIGNORE='*' command eval  'ips=($(cat nodeIP.config))'
ipslength=${#ips[@]}

install_docker(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'yes | sudo apt-get install docker.io'
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'sudo usermod -a -G docker $USER'
  scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i "" dos.txt  ubuntu@${ips[$i]}:~/
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'cat ~/dos.txt | docker login --username dosnetwork --password-stdin'
}

install_lightnode(){
  echo ubuntu@[${ips[$i]}]
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'ps cax | grep geth > /dev/null'
  if [ $? -eq 0 ]
  then
      echo "Geth Process is running."
  else
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'sudo apt-get update'
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'sudo add-apt-repository -y ppa:ethereum/ethereum'
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'sudo apt-get update'
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'yes | sudo apt-get install ethereum'
    scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i "" geth.service  ubuntu@${ips[$i]}:~/
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'sudo chmod 322 geth.service'
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'sudo mv geth.service /lib/systemd/system/geth.service'
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'sudo systemctl enable geth'
    ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'sudo systemctl start geth'
  fi	
}

check_geth(){
  ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'ps cax | grep geth > /dev/null'
  if [ $? -eq 0 ]
  then
      echo "Geth Process is running."
  else
      echo "Geth Process is not running."
      exit 1
  fi	
}

update_dos_config(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'mkdir -p credential'
	scp -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i "" testAccounts/$(($i + 1))/credential/usrKey ubuntu@${ips[$i]}:~/credential/
}

run_dos(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'docker pull dosnetwork/dosnode:beta'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'docker run -it -d \
--mount type=bind,source=/home/ubuntu/credential,target=/credential  \
--mount type=bind,source=/home/ubuntu/credential,target=/app-logs  \
-e LOGIP=  \
-e CHAINNODE=rinkeby  \
-e PASSPHRASE=  \
-e APPSESSION=BETA  \
-e APPNAME=DosNode  \
dosnetwork/dosnode:beta'
}

dos_log(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'cat credential/doslog'
}

stop_dos(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'docker stop $(docker ps -a -q)'
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'docker rm $(docker ps -a -q)'
}

guardian(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'curl http://localhost:8080/guardian'
}

proxy(){
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -tt -i "" ubuntu@${ips[$i]} 'curl http://localhost:8080/proxy'
}

case "$1" in
  "init_ssh")
    for (( i=0; i<${ipslength}; i++ ));
    do
      check_geth $i 
    done
    ;;
  "install_docker")
    for (( i=0; i<${ipslength}; i++ ));
    do
      install_docker $i &
    done
    wait
    ;;
  "install_lightnode")
    for (( i=0; i<${ipslength}; i++ ));
    do
        install_lightnode $i & # Put a function in the background
    done
    wait
    echo "All done"
    ;;
  "install_dosuser")
    for (( i=1; i<=${AMACOUNT}; i++ ));
    do
        cp dos_user dosUser${i}/
        cp config.json dosUser${i}/
        cp ama.json dosUser${i}/
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
    for (( i=0; i<21; i++ ));
    do
        run_dos $i & # Put a function in the background
    done
    wait
    for (( i=21; i<200; i++ ));
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
  "guardian")
    guardian 0
    echo "All done"
    ;;
  "proxy")
    proxy 0
    echo "All done"
    ;;
  *)
    echo "You have failed to specify what to do correctly."
    exit 1
    ;;
esac

