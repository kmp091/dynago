package service

import (
	"context"

	anypb "github.com/golang/protobuf/types/known/anypb"
	pb "github.com/kmp091/dynago/messages"
)

type server struct {
	pb.UnimplementedDynagoServiceServer
}

func (s *server) Process(ctx context.Context, request *pb.DynagoRequest) (*pb.DynagoResponse, error) {
	//TODO: run plugins
	var response pb.DynagoResponse
	response.Results["FinalResult"] = anypb.New(pb.ValueTypeParameter_DoubleValue{5.5})
	return &response, nil
}
