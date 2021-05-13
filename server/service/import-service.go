package service

import {
	pb "github.com/kmp091/dynago/messages"
	redis "github.com/go-redis/redis/v8"
}

type ImportService struct {
	pb.UnimplementedImportPluginServiceServer
	RedisHostName   *string
	RedisPortNumber *int
	RedisSecretPath *string
}

func AddSampleRedisKey(host *string, port *int, secretPath *string) {
	secret, secretErr := ioutil.ReadFile(*secretPath)
	if secretErr != nil {
		log.Fatal(secretErr)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", *host, *port),
		Password: string(secret),
		DB:       0, // use default DB
	})

	ctx := context.Background()
	status := rdb.Set(ctx, "key", "value2", 0)
	_, err := status.Result()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *ImportServer) Import(context.Context, *ImportPluginRequest) (*ImportPluginResponse, error) {
	AddSampleRedisKey(s.RedisHostName, s.RedisPortNumber, s.RedisSecretPath)
	return nil, status.Errorf(codes.Unimplemented, "method Import not implemented")
}
