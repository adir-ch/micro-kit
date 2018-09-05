# Distributed Calculator
    A distributed calculator built with Microservices using go-kit, using etcd as service registry

## Kit CLI Commands 
 
    - Generate new services: 
        1. # kit new service <service-name>
        2. Add functions to the service interface
        3. # kit generate service add --dmw
            --dmw = generate default middlewares 
            
    - Generate Docerfile and docker compose files
        # kit generate docker - will generate a docker compose file 

    - docker-compose up - will run the docker images 

    - Send data to the services: 
        Calc FE: curl -s -d '{"expr": "1 + 1"}' -X POST localhost:8800/calculate
        Add service: # curl -s -d '{"numbers": [1, 2, 3]}' -X POST localhost:8800/add

    - etcd: 
        * Start only etcd (-d = run in the bg): # docker-compose up -d etcd 
        * Query etcd: 
            version: # curl -L http://127.0.0.1:23791/version
            get all keys: # curl -L http://127.0.0.1:23791/v2/keys/?recursive=true
            add new key: # curl http://127.0.0.1:23791/v2/keys/services/v2/add:8081/ -X PUT -d value="add:8081/calc"
            find by key: # curl -L http://127.0.0.1:23791/v2/keys/services/calc/
            delete a key: #  curl http://127.0.0.1:23791/v2/keys/services/v2/add:8081/ -X DELETE