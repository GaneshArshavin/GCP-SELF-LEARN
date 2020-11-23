#!/bin/bash
docker build . -t gcp-self-study || exit
docker run  -p 8080:8080 -d --name gcp-self-study gcp-self-study go run server/main.go
docker exec -ti gcp-self-study /bin/bash
