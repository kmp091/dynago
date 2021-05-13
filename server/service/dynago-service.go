package service

import (
	"context"
	"log"
	"os"
	"path"
	"plugin"

	"github.com/go-redis/redis/v8"
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
	RedisClient *redis.Client
}

func (s *DynagoService) ReadCache(key string) ([]byte, error) {
	ctx := context.Background()
	result := s.RedisClient.Get(ctx, key)
	return result.Bytes()
}

func (s *DynagoService) Process(ctx context.Context, request *pb.DynagoRequest) (*pb.DynagoResponse, error) {
	//service level is responsible for unmarshaling and marshaling Protobuf layer

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

	//future enhancement: plugin doesn't have byte array support
	//maybe using go plugin is not the best idea
	pluginData, redisGetError := s.ReadCache("multiplication-plugin")

	if redisGetError != nil {
		log.Fatal("Unable to get plugin from cache", redisGetError)
	}

	pluginParentDir, pluginName := "./plugins", "multiplication-plugin.so"
	pluginPath := path.Join(pluginParentDir, pluginName)

	pluginFile, _ := os.Create(pluginPath)
	pluginFile.Write(pluginData)
	pluginFile.Close()

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
