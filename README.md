# dynago
This project uses Go plugins to dynamically execute code on a freeform request object.
The request and response objects are Protobuf-serialized (as you would) and transmitted via gRPC.

## Server 
- processing requests with arbitrary parameters, response has arbitrary data as well
- use Go plugins for dynamically adding processing logic
- accept any message type in the parameters (this comes with some repercussions for serialization)
- add dynamic server logic with go plugins
- Dockerized so you can do:
```bash
docker-compose up --build
```

## Client
- Go lang based sample gRPC Client

## Future Plans
- server: open an API to import and build plugins
- .NET 5 based gRPC Client
- eventually, a web frontend for modifying, testing and importing plugins
- port numbers are not hardcoded