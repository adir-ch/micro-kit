#! /bin/bash

echo "--- building code --- "
echo 

ls -d */ | cut -d '/' -f1 | while read service 
    do 
        echo "building service $service ---> start"
        rm -f ./$service/${service}.service
        go build -o ${service}.service $service/cmd/main.go && mv ${service}.service ./$service/
        echo "building service $service ---> finished"
    done

echo
echo "--- building docker images ---"
docker-compose build --force-rm --no-cache
