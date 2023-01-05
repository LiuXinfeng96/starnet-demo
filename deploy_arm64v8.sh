#!/bin/bash
path=`pwd`

i=$(docker images | grep "arm64v8/starnet-demo" | awk '{print $1}')
if test -z $i; then
echo "not found the docker image, start build image..."
docker build -f ./DockerFile.arm64v8 -t arm64v8/starnet-demo:v0.9.0 ../starnet-demo
fi

i=$(docker images | grep "arm64v8/starnet-demo" | awk '{print $1}')
if test -z $i; then
echo "build image error, exit shell!"
exit
fi

c=$(docker ps -a | grep "arm64v8/starnet-mysql" | awk '{print $1}')
if test -z $c; then
echo "not found the mysql server, start mysql server..."
docker run -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=chainmaker \
    --name arm64v8/starnet-mysql \
    --restart always \
    arm64v8/mysql:8.0
echo "waiting for database initialization..."
sleep 20s
docker logs --tail=10 arm64v8/starnet-mysql
fi

i=$(docker ps -a | grep "arm64v8/starnet-demo" | awk '{print $1}')
if test ! -z $i; then
echo "the server container already exists, delete..."
docker rm -f arm64v8/starnet-demo
fi

echo "start starnet demo server..."
docker run -d \
-p 8086:8086 \
-w /starnet-demo \
-v $path/conf:/starnet-demo/conf \
-v $path/log:/starnet-demo/log \
-v $path/crypto-config:/starnet-demo/crypto-config \
--name starnet-demo \
--restart always \
arm64v8/starnet-demo:v0.9.0 \
bash -c "cd src&&./starnet-demo -config ../conf/system_config.yaml"
sleep 2s
docker logs arm64v8/starnet-demo
echo "the starnet demo server has been started!"


i=$(docker ps -a | grep "arm64v8/starnet-web" | awk '{print $1}')
if test ! -z $i; then
echo "the web container already exists, delete..."
docker rm -f arm64v8/starnet-web
fi

echo "start starnet demo web server..."
chmod -R 777 $path/web/
docker run -d \
-p 80:80 \
-v $path/web/conf.d:/etc/nginx/conf.d \
-v $path/web/resources:/usr/share/nginx/resources \
--name arm64v8/starnet-web \
--restart always \
arm64v8/nginx:1.23.3
echo "the starnet demo web server has been started!"