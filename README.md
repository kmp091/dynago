# dynago
This project uses Go plugins to dynamically execute code on a freeform request object.
The request and response objects are Protobuf-serialized (as you would) and transmitted via gRPC.

## Server 
- processing requests with arbitrary parameters, response has arbitrary data as well
- use Go plugins for dynamically adding processing logic
- accept any message type in the parameters
- Dockerized so you can do:
```bash
docker build . -t dynago
docker run -p 4323:4323 dynago
```
(Port numbers are currently hardcoded. I plan to do a docker compose file soon enough)

## Client
- Go lang based sample gRPC Client

## Future Plans
- server: open an API to import and build plugins
- .NET 5 based gRPC Client
- eventually, a web frontend for modifying, testing and importing plugins
- port numbers are not hardcoded