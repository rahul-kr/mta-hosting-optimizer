# mta-hosting-optimizer
mta-hosting-optimizer

CONTENTS OF THIS FILE
---------------------

 * Introductionn
 * Requirements
 * Installation
 * Maintainers

INTRODUCTION
------------

 * This will expose the rest API on docker based environemt.
 * The rest api details are -
 * 
 * Get Rest API will return the message reated to health check else return error
 * Run the command - curl http://localhost:8080/ping
 * Outcome - "Health check is OK"
 * 
 * Post rest API for create the sample data
 * curl -H 'Content-Type: application/json' -d '{"ip":"127.0.0.1","hostname":"mta-prod-1","active":true}' -X POST http://localhost:8080/create_sample_data 
 * curl -H 'Content-Type: application/json' -d '{"ip":"127.0.0.2","hostname":"mta-prod-1","active":true}' -X POST http://localhost:8080/create_sample_data 
 * curl -H 'Content-Type: application/json' -d '{"ip":"127.0.0.3","hostname":"mta-prod-2","active":false}' -X POST http://localhost:8080/create_sample_data 
 * curl -H 'Content-Type: application/json' -d '{"ip":"127.0.0.4","hostname":"mta-prod-2","active":true}' -X POST http://localhost:8080/create_sample_data 
 * curl -H 'Content-Type: application/json' -d '{"ip":"127.0.0.5","hostname":"mta-prod-3","active":true}' -X POST http://localhost:8080/create_sample_data 
 * curl -H 'Content-Type: application/json' -d '{"ip":"127.0.0.6","hostname":"mta-prod-4","active":true}' -X POST http://localhost:8080/create_sample_data 
 * curl -H 'Content-Type: application/json' -d '{"ip":"127.0.0.7","hostname":"mta-prod-4","active":true}' -X POST http://localhost:8080/create_sample_data 
 * curl -H 'Content-Type: application/json' -d '{"ip":"127.0.0.8","hostname":"mta-prod-4","active":true}' -X POST http://localhost:8080/create_sample_data 
 * 
 *  
 * Get Rest API will return all the data related to host mapping
 * Run the command - curl http://localhost:8080/mappings
 * 
 * Like  - [{"id":1,"ip":"127.0.0.1","hostname":"mta-prod-1","active":true},{"id":2,"ip":"127.0.0.2","hostname":"mta-prod-1","active":true},{"id":3,"ip":"127.0.0.3","hostname":"mta-prod-2"},{"id":4,"ip":"127.0.0.4","hostname":"mta-prod-2","active":true},{"id":5,"ip":"127.0.0.5","hostname":"mta-prod-3","active":true},{"id":6,"ip":"127.0.0.6","hostname":"mta-prod-4","active":true},{"id":7,"ip":"127.0.0.7","hostname":"mta-prod-4","active":true},{"id":8,"ip":"127.0.0.8","hostname":"mta-prod-4","active":true}]
 * 
 * 
 * Get Rest API will return host name the related active IP count of host mapping
 * curl http://localhost:8080/threshold/2                                                                 
 * Outcome - [{"hostname":"mta-prod-1"}]

 



REQUIREMENTS
-------------------

 * Docker and permissions to execute the required command

INSTALLATION
------------

 * Go to the root folder i.e mta-hosting-optimizer and run the
 * docker-compose up -d


MAINTAINERS
-----------

Current maintainers:

 * Rahul kumar (rahul2u88@gmail.com)

