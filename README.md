# demo

this is a demo used to register services.

## test
```http request
POST http://localhost:8081/services/node?server_name=123&address=456&user_name=123&address=456
Accept: */*
Cache-Control: no-cache

{"server_name":"localhost","address":"127.0.0.1:8000","provide":["calculator","i_o_server"],"protocols":3}

```
