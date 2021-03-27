package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/kmp091/dynago/messages"
	"github.com/kmp091/dynago/server/service"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 4323, "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Unable to listen to port %d\n", *port)
		return
	}

	server := service.DynagoService{}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	messages.RegisterDynagoServiceServer(grpcServer, &server)
	grpcServer.Serve(listener)
}
