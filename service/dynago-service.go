package service

import (
	"context"
	"log"
	"plugin"

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
					returnVal = p.GetParameterOneof()
				default:
					returnVal = parameterValue
				}
			}
		}

		return returnVal, ok
	}

	//TODO: run plugins
	activePlugin, pluginErr := plugin.Open("plugins/multiplication-plugin.so")
	if pluginErr != nil {
		log.Fatal("Plugin load failure")
	}

	sym, symErr := activePlugin.Lookup("Process")
	pluginImpl, ok := sym.(DynagoPlugin)

	if symErr != nil || !ok {
		log.Fatal("Plugin does not implement Process()")
	}

	response := pb.DynagoResponse{
		Results: make(map[string]*anypb.Any),
	}

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
