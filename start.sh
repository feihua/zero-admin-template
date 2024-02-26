#!/bin/bash

#停止服务
docker stop api

#删除容器
docker rm api

#删除镜像
docker rmi api:v1

#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t api:v1 -f api/Dockerfile .

#启动服务
docker run -itd --net=host --name=api api:v1