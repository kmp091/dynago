package service

import (
	"context"
	"fmt"
	"log"
	"path"
	"plugin"

	redis "github.com/go-redis/redis/v8"
	pb "github.com/kmp091/dynago/messages"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

type DynagoPlugin interface {
	//parameter list is read only and only accessible via a closure
	//however, the return value is freeform
	Process(parameterAccessor func(string) (interface{}, bool)) *map[string]interface{}
}

type DynagoService struct {
	pb.UnimplementedDynagoServiceServer
	RedisHostName   *string
	RedisPortNumber *int
}

func AddSampleRedisKey(host *string, port *int) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", *host, *port),
		Password: "portugal_CUCUMBER_bleucheese",
		DB:       0, // use default DB
	})

	ctx := context.Background()
	status := rdb.Set(ctx, "key", "value", 0)
	_, err := status.Result()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *DynagoService) Process(ctx context.Context, request *pb.DynagoRequest) (*pb.DynagoResponse, error) {
	//service level is responsible for unmarshaling and marshaling Protobuf layer

	AddSampleRedisKey(s.RedisHostName, s.RedisPortNumber)

	accessor := func(key string) (interface{}, bool) {
		pbValue, ok := request.Parameters[key]
		var returnVal interface{}

		if ok {
			parameterValue, err := pbValue.UnmarshalNew()
			if err == nil {
				switch p := parameterValue.(type) {
				case *pb.ValueTypeParameter:
					oneofValue := p.GetParameterOneof()
					switch oneofType := oneofValue.(type) {
					case *pb.ValueTypeParameter_IntegerValue:
						returnVal = oneofType.IntegerValue
					case *pb.ValueTypeParameter_BoolValue:
						returnVal = oneofType.BoolValue
					case *pb.ValueTypeParameter_DoubleValue:
						returnVal = oneofType.DoubleValue
					case *pb.ValueTypeParameter_StringValue:
						returnVal = oneofType.StringValue
					}
				default:
					returnVal = parameterValue
				}
			}
		}

		return returnVal, ok
	}

	pluginParentDir, pluginName := "./plugins", "multiplication-plugin.so"
	pluginPath := path.Join(pluginParentDir, pluginName)

	activePlugin, pluginErr := plugin.Open(pluginPath)
	if pluginErr != nil {
		log.Fatalf("Plugin load failure: %s", pluginErr)
	}

	sym, symErr := activePlugin.Lookup("Implementation")
	pluginImpl, ok := sym.(DynagoPlugin)

	if symErr != nil {
		log.Fatalf("Plugin could not load symbol Process: %s", symErr)
	} else if !ok {
		log.Fatal("Plugin does not implement Process()")
	}

	response := pb.DynagoResponse{
		Results: make(map[string]*anypb.Any),
	}

	log.Printf("Now running plugin %v\n", pluginName)
	pluginResult := pluginImpl.Process(accessor)

	var err error
	//marshal pluginResult to anypb.Any
	for key, val := range *pluginResult {
		//check if val is compatible with typical value types
		var message protoreflect.ProtoMessage

		switch v := val.(type) {
		case float32:
			floatValue := pb.ValueTypeParameter_DoubleValue{DoubleValue: float64(v)}
			message = &pb.ValueTypeParameter{ParameterOneof: &floatValue}
		case float64:
			floatValue := pb.ValueTypeParameter_DoubleValue{DoubleValue: v}
			message = &pb.ValueTypeParameter{ParameterOneof: &floatValue}
		case string:
			stringValue := pb.ValueTypeParameter_StringValue{StringValue: v}
			message = &pb.ValueTypeParameter{ParameterOneof: &stringValue}
		case int:
			intValue := pb.ValueTypeParameter_IntegerValue{IntegerValue: int32(v)}
			message = &pb.ValueTypeParameter{ParameterOneof: &intValue}
		case int32:
			intValue := pb.ValueTypeParameter_IntegerValue{IntegerValue: v}
			message = &pb.ValueTypeParameter{ParameterOneof: &intValue}
		case int64:
			intValue := pb.ValueTypeParameter_IntegerValue{IntegerValue: int32(v)}
			message = &pb.ValueTypeParameter{ParameterOneof: &intValue}
		case bool:
			boolValue := pb.ValueTypeParameter_BoolValue{BoolValue: v}
			message = &pb.ValueTypeParameter{ParameterOneof: &boolValue}
		default:
			//not a known value type parameter
			//unsupported
			continue
		}

		response.Results[key], err = anypb.New(message)
		if err != nil {
			log.Default().Println("Could not serialize parameter. Skipping...")
		}
	}

	return &response, err
}
