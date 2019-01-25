#!/bin/bash -x

docker network create --subnet=10.250.0.1/16 CovenantSQL_testnet
echo starting china10.250.1.2
docker run -d --net CovenantSQL_testnet --ip 10.250.1.2 -h china10.250.1.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name china10.250.1.2 gnte /scripts/china10.250.1.2.sh
echo starting china10.250.2.2
docker run -d --net CovenantSQL_testnet --ip 10.250.2.2 -h china10.250.2.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name china10.250.2.2 gnte /scripts/china10.250.2.2.sh
echo starting china10.250.3.2
docker run -d --net CovenantSQL_testnet --ip 10.250.3.2 -h china10.250.3.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name china10.250.3.2 gnte /scripts/china10.250.3.2.sh
echo starting china10.250.4.2
docker run -d --net CovenantSQL_testnet --ip 10.250.4.2 -h china10.250.4.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name china10.250.4.2 gnte /scripts/china10.250.4.2.sh
echo starting eu10.250.5.2
docker run -d --net CovenantSQL_testnet --ip 10.250.5.2 -h eu10.250.5.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name eu10.250.5.2 gnte /scripts/eu10.250.5.2.sh
echo starting eu10.250.6.2
docker run -d --net CovenantSQL_testnet --ip 10.250.6.2 -h eu10.250.6.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name eu10.250.6.2 gnte /scripts/eu10.250.6.2.sh
echo starting eu10.250.7.2
docker run -d --net CovenantSQL_testnet --ip 10.250.7.2 -h eu10.250.7.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name eu10.250.7.2 gnte /scripts/eu10.250.7.2.sh
echo starting jpn10.250.8.2
docker run -d --net CovenantSQL_testnet --ip 10.250.8.2 -h jpn10.250.8.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name jpn10.250.8.2 gnte /scripts/jpn10.250.8.2.sh
echo starting jpn10.250.9.2
docker run -d --net CovenantSQL_testnet --ip 10.250.9.2 -h jpn10.250.9.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name jpn10.250.9.2 gnte /scripts/jpn10.250.9.2.sh
echo starting jpn10.250.10.2
docker run -d --net CovenantSQL_testnet --ip 10.250.10.2 -h jpn10.250.10.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name jpn10.250.10.2 gnte /scripts/jpn10.250.10.2.sh
echo starting jpn10.250.11.2
docker run -d --net CovenantSQL_testnet --ip 10.250.11.2 -h jpn10.250.11.2 -v $DIR/scripts:/scripts --cap-add=NET_ADMIN --name jpn10.250.11.2 gnte /scripts/jpn10.250.11.2.sh
docker run --rm -v $DIR/scripts:/scripts gnte dot -Tpng scripts/graph.gv -o scripts/graph.png
mv -f $DIR/scripts/graph.png $DIR/graph.png
