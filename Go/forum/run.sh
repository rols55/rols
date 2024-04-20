#! /bin/bash

docker build --tag docker-forum:latest .

docker run --publish 8080:8080 docker-forum:latest