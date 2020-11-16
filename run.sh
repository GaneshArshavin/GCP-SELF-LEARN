#!/bin/bash
docker build . -t chope-assignment || exit
docker run  -p 8080:8080 -d --name chope-assignment chope-assignment go run server/main.go
echo ""
echo "Type 'supervisorctl status' to view services status"
echo "Type 'logs' to view app logs"
echo "Type 'code' to go to app code"
echo "Type 'slogs' to wait 20 seconds and then start tailing logs"
docker exec -ti chope-assignment /bin/bash
