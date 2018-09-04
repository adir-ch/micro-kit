# Distributed Calculator
    A distributed calculator built with Microservices using go-kit

## Kit CLI Commands 
 
    - Generate new services: 
        1. # kit new service <service-name>
        2. Add functions to the service interface
        3. # kit generate service add --dmw
            --dmw = generate default middlewares 
            
    - Generate Docerfile and docker compose files
        # kit generate docker - will generate a docker compose file 

    - docker compose up - will run the docker images 

    - Send data to the services: 
        Calc FE: curl -s -d '{"expr": "1 + 1"}' -X POST localhost:8800/calculate
        Add service: # curl -s -d '{"numbers": [1, 2, 3]}' -X POST localhost:8800/add