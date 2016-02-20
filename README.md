# onebox-dynamic-pricing
A golang implementation of onebox-dynamic-pricing-api, an API to share sails and prices data in order to retrieve recommended prices in real time.
Find specifications at:
+ [github](https://github.com/joliva-ob/onebox-dynamic-pricing-api)
+ [onebox-developer](http://developer.oneboxtickets.com/dynamic-pricing-api)

## TODO list
+ Link oauth to the server oauth
+ adapt json responses to API specifications
+ API filters (dates between, page num. (page size with default limit = 100), default values (last 100 items as of now)
+ handle general errors
+ catch mysql responses, coger informacion de DAL-mysql o de elasticsearch o MS de prices, TTL 1 min.

## DONE list
+ rehuse mysql sql.DB connections pool
+ load configurations from external file
+ connect to data store to retrieve prices information
+ marshall prices to json throw endpoint prices
+ OAUTH2 authentication reading bear header
