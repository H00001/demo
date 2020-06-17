# demo

this is a demo used to register services.

## test
```http request
POST http://localhost:8081/services/node?server_name=123&address=456&user_name=123&address=456
Accept: */*
Cache-Control: no-cache

{"server_name":"hello","address":"456","provide":["calc","io"],"protocols":[1,2]}

```
