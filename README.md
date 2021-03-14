# dynago
Just tinkering around with Go plugins (and learning Go itself)

## The plan
### Server 
- processing requests with arbitrary parameters, response has arbitrary data as well
- use Go plugins for dynamically adding processing logic
- accept any message type in the parameters
- eventually open an endpoint to import plugins?

### Client
- Go and .NET 5 versions of gRPC Client
- eventually, a web frontend for modifying, testing and importing plugins