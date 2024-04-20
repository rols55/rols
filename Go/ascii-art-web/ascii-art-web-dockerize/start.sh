#!/bin/sh
#to build Docker image from Dockerfile
DOCKER_SCAN_SUGGEST=false docker build -t ascii .
#to run a Docker container based on a given Docker image
docker run --name ascii-container -p 8080:8080 ascii &
echo "Visit in your browser http://localhost:8080"

echo "Press any key to remove image/container"
read -n 1 -s

#to find container id
containerId=$(docker ps -aqf "name=ascii")
#to stop running container
docker container stop $containerId
#to remove container
docker container rm $containerId
#to remove image
docker rmi -f ascii
echo "Files are removed! Adios!"