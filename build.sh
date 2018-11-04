#! /bin/bash

set -e # exit when something fails

echo "--- building code --- "
echo 

target=$1

if [ "$target" == "" ] 
then    
    echo ""
    echo "--- building all services ---" 
    echo ""
else 
    echo ""
    echo "--- building target $target ---" 
    echo ""
fi

ls -d */ | xargs -e basename -a | while read service 
    do 
        #echo "check service $service"
        if [ "$target" != "" ] && [ "$target" != "$service" ]
        then 
            continue
        fi
        echo "building service $service ---> start"
        rm -f ./$service/${service}.service
        go build -o ${service}.service $service/cmd/main.go 
        mv ${service}.service ./$service/
        echo "building service $service ---> finished"
        echo "--- building docker images ---"
        docker-compose rm --force $service
        docker-compose build --force-rm --no-cache $service
    done