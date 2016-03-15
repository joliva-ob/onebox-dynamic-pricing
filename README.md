# onebox-dynamic-pricing
A golang implementation of onebox-dynamic-pricing-api, an API to share sails and prices data in order to retrieve recommended prices in real time.
Find specifications at:
+ [github](https://github.com/joliva-ob/onebox-dynamic-pricing-api)
+ [onebox-developer](http://developer.oneboxtickets.com/dynamic-pricing-api)
Compiled with runtime with: 
+ GOOS=windows GOARCH=386 go build -o dynamic-pricing.exe dynamic-pricing.go
+ GOOS=linux GOARCH=386 go build -o dynamic-pricing.linux dynamic-pricing.go
+ GOOS=darwin GOARCH=386 go build -o dynamic-pricing dynamic-pricing.go



## TODO list
+ export status and version to /info and /health
+ handle panic errors and recover
+ add endpoint /summaries
+ Link oauth to the server oauth trough eureka register


## Optional TODO list
+ coger informacion de DAL-mysql o de elasticsearch o MS de prices, TTL 1 min.
+ version history
+ API links
+ Track and audit to monitoring api console
+ add unit tests and mocks
+ add log level as a main app argument
+ filter by price_id (priceZoneId)
+ filter by sale_id


## DONE list
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