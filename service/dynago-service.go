package service

import (
	"context"

	pb "github.com/kmp091/dynago/messages"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

type DynagoService struct {
	pb.UnimplementedDynagoServiceServer
}

func (s *DynagoService) Process(ctx context.Context, request *pb.DynagoRequest) (*pb.DynagoResponse, error) {
	//TODO: run plugins
	response := pb.DynagoResponse{
		Results: make(map[string]*anypb.Any),
	}

	var err error
	doubleParameter := pb.ValueTypeParameter_DoubleValue{DoubleValue: 5.5}
	response.Results["FinalResult"], err = anypb.New(&pb.ValueTypeParameter{ParameterOneof: &doubleParameter})
	return &response, err
}
