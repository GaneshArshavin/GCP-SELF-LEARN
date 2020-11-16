#!/bin/bash
docker build . -t chope-assignment || exit
docker run  -p 8080:8080 -d --name chope-assignment chope-assignment go run server/main.go
docker exec -ti chope-assignment /bin/bash
