# onebox-dynamic-pricing
A golang implementation of onebox-dynamic-pricing-api, an API to share sails and prices data in order to retrieve recommended prices in real time.
Find specifications at:
+ [github](https://github.com/joliva-ob/onebox-dynamic-pricing-api)
+ [onebox-developer](http://developer.oneboxtickets.com/dynamic-pricing-api)

Compiled with runtime with: 
+ GOOS=windows GOARCH=386 go build -o dynamic-pricing.exe dynamic-pricing.go
+ GOOS=linux GOARCH=386 go build -o dynamic-pricing.linux dynamic-pricing.go
+ GOOS=darwin GOARCH=386 go build -o dynamic-pricing dynamic-pricing.go

Build Docker image with
+ cp /source_cfg_files/*env* .
+ docker build -f docker/Dockerfile . -tag dynamic-pricing
+ docker run --publish 8000:8000 --name dynamic-pricing --rm dynamic-pricing --restart=always dynamic-pricing



## TODO list
+ set forced list of events per client authenticated from configuration
+ filter forced events list per oauth client

+ Monitorize API
+ Get environment and general configurations from environment vars
+ Versioning policy
+ Handle requests by gorutines pool and control them by channels
+ unit testing
+ add endpoint /summaries
+ Perfomance: Handle requests by gorutines pool and control them by channels (Marcio.io)
+ handle panic errors and recover it
+ dockerize app


## Optional TODO list
+ extend /info and /health with discovery service status, resources statuses, and version from file + git branch
+ coger informacion de DAL-mysql o de elasticsearch o MS de prices, TTL 1 min.
+ version history
+ API links
+ Track and audit to monitoring api console
+ add unit tests and mocks
+ add log level as a main app argument


## DONE list

# v0.0.1
+ initial version from db to json server endpoint
+ rehuse mysql sql.DB connections pool
+ load configurations from external file
+ connect to data store to retrieve prices information
+ marshall prices to json throw endpoint prices
+ OAUTH2 authentication reading bear header
+ redirect logs to a file and standard output
+ added log levels
+ adapt json responses to API specifications
+ API filters (dates between, page num. (page size with default limit = 10), default values (last 10 items as of now)
+ cache sql pagination / optimize select
+ Add missing information to datasources. link event with session and entity type venue (DB, Solr, couch,...)
+ add endpoint /sales
+ add cache to sales-dataservice.getSales
+ add usage instructions
+ adjust to specifications
+ register to eureka
+ eureka details from config
+ log requested url + origin ip address
+ Añadir identificador único a las transacciones (traceID)
+ export status and version to /info and /health like onebox microservices
+ added price_zone_name to any product from sales

# v1.0.0
+ sql query + event_id
+ elk query + sale_id (doc.code)
+ get the params event_id and sale_id
+ set query and calls
+ api change event_id for id and sale_id for id
+ update documentation and versioning
+ authorize via oauth server and cache it