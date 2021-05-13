package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/kmp091/dynago/messages"
	"github.com/kmp091/dynago/server/service"
	"google.golang.org/grpc"
)

var (
	port            = flag.Int("port", 4323, "The server address in the format of host:port")
	redisHost       = flag.String("redis-host", "redis", "The host name for the Redis server")
	redisPort       = flag.Int("redis-port", 6379, "The port number for the Redis server")
	redisSecretPath = flag.String("redis-secret-path", "", "The file path for the password configured for the Redis server")
)

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Unable to listen to port %d\n", *port)
		return
	}

	secret, secretErr := ioutil.ReadFile(*redisSecretPath)
	if secretErr != nil {
		log.Fatal(secretErr)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", *redisHost, *redisPort),
		Password: string(secret),
		DB:       0, // use default DB
	})

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	computationService := service.DynagoService{
		RedisClient: redisClient,
	}
	messages.RegisterDynagoServiceServer(grpcServer, &computationService)

	importService := service.ImportService{
		RedisClient: redisClient,
	}
	messages.RegisterImportPluginServiceServer(grpcServer, &importService)

	grpcServer.Serve(listener)
}
