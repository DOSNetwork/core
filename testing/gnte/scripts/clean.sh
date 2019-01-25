#!/bin/bash -x

docker rm -f china10.250.1.2
docker rm -f china10.250.2.2
docker rm -f china10.250.3.2
docker rm -f china10.250.4.2
docker rm -f eu10.250.5.2
docker rm -f eu10.250.6.2
docker rm -f eu10.250.7.2
docker rm -f jpn10.250.8.2
docker rm -f jpn10.250.9.2
docker rm -f jpn10.250.10.2
docker rm -f jpn10.250.11.2
docker network rm CovenantSQL_testnet
