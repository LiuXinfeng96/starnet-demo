#!/bin/bash
path=`pwd`

i=$(docker images | grep "starnet-demo" | awk '{print $1}')
if test -z $i; then
echo "not found the docker image, start build image..."
docker build -f ./DockerFile -t starnet-demo:v0.9.0 ../starnet-demo
fi

i=$(docker images | grep "starnet-demo" | awk '{print $1}')
if test -z $i; then
echo "build image error, exit shell!"
exit
fi

c=$(docker ps -a | grep "starnet-mysql" | awk '{print $1}')
if test -z $c; then
echo "not found the mysql server, start mysql server..."
docker run -d \
    -p 13306:3306 \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=chainmaker \
    --name starnet-mysql \
    --restart always \
    mysql:8.0
echo "waiting for database initialization..."
sleep 20s
docker logs --tail=10 starnet-mysql
fi

i=$(docker ps -a | grep "starnet-demo" | awk '{print $1}')
if test ! -z $i; then
echo "the container already exists, delete..."
docker rm -f starnet-demo
fi

echo "start starnet demo server..."
docker run -d \
-p 8096:8096 \
-w /starnet-demo \
-v $path/conf:/starnet-demo/conf \
-v $path/log:/starnet-demo/log \
-v $path/crypto-config:/starnet-demo/crypto-config \
--name starnet-demo \
--restart always \
starnet-demo:v0.9.0 \
bash -c "cd src&&./starnet-demo -config ../conf/system_config.yaml"
sleep 2s
docker logs starnet-demo
echo "the starnet demo server has been started!"
