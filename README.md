# Distributed Calculator
    A distributed calculator built with Microservices using go-kit, using etcd as service registry

## Go-Kit architecture 

    - Setup
        1. Setup is done in the cmd/service 
        2. Hook all endpoints + middleware in service.go

    - Request flow: client req -> handler -> endpoint -> service (logic) -> endpoint -> handler -> client res
        1. Request is received by the handler (transport specific)
        2. The transport handler will convert / deserialize the data and will call the endpoint. 
        3. Endpoint will validate the input data and will call the service 
        4. The service will contain the business logic. 

## Kit CLI commands 
 
    - Generate new services with HTTP endpoints: 
        1. # kit new service <service-name>
        2. Add functions to the service interface
        3. # kit generate service <service-name> --dmw (dmw: generate default middlewares)

    - Generate new services with GRPC endpoints: 
        1 + 2 - as in HTTP service 
        3. # kit generate service <service-name> -t grpc --dmw (dmw: generate default middlewares)

    - Generate Docerfile and docker compose files
        # kit generate docker - will generate a docker compose file 

## Docker compose commands 

    - Starting / View a service/s
        # docker-compose ps - show all composed images 
        # docker-compose up - will run the docker images 
        # docker-compose up <service name> - will run the docker image of a specific service 
        # docker-compose up -d <service name> - will run the docker image of a specific service in the background

    - Build composed package: 
        # docker-compose build --force-rm --no-cache <image name>

    - Remove composed image 
        # docker-compose rm <image name>

    - Send data to the services: 
        Calc FE: curl -s -d '{"expr": "1 + 1"}' -X POST localhost:8800/calculate
        Add service: # curl -s -d '{"numbers": [1, 2, 3]}' -X POST localhost:8800/add

## gRPC / Protobuf 

    - Install: 
        * curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
        * unzip protoc-3.6.1-linux-x86_64.zip -d protoc3
        * sudo mv protoc3/bin/* /usr/local/bin/
        * sudo mv protoc3/include/* /usr/local/include/

    - Compile 
        1. Kit will create a compile.sh in the service protobuf transport folder. 
        2. Run the compile.sh to generate the .go file from the .proto file
        3. The .go file contains: 
            3.1. The service protobuf implementation in Go. 
            3.2. Client that can speak with the gRPC service that was just created. 


        
## External services / DB's 
    - etcd: 
        * Start only etcd (-d = run in the bg): # docker-compose up -d etcd 
        * Query etcd: 
            version: # curl -L http://127.0.0.1:23791/version
            get all keys: # curl -L http://127.0.0.1:23791/v2/keys/?recursive=true
            add new key: # curl http://127.0.0.1:23791/v2/keys/services/add:8081/ -X PUT -d value="add:8081"
            find by key: # curl -L http://127.0.0.1:23791/v2/keys/services/calc/
            delete a key: #  curl http://127.0.0.1:23791/v2/keys/services/v2/add:8081/ -X DELETE

    - Postgres docker image run:
        docker run -it --rm -e POSTGRES_PASSWORD='' -p 5432:5432 postgres:alpine

    - Redis docker image spinup:
        docker run -p "6379:6379" --rm --name dwarf-redis redis:4-alpine