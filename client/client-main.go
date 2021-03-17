package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/kmp091/dynago/messages"
	"google.golang.org/grpc"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	serverAddress = "localhost:4323"
)

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDynagoServiceClient(conn)

	parameters := make(map[string]*anypb.Any)
	detailString := pb.ValueTypeParameter_StringValue{StringValue: "This is the thing you process."}
	parameters["Detail"], err = anypb.New(&pb.ValueTypeParameter{ParameterOneof: &detailString})

	request := pb.DynagoRequest{
		Parameters: parameters,
	}

	ctx := context.Background()

	response, err := client.Process(ctx, &request)
	resultMap := response.GetResults()
	fmt.Println(resultMap["FinalResult"])
}
