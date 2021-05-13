package service

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path"

	redis "github.com/go-redis/redis/v8"
	pb "github.com/kmp091/dynago/messages"
)

type ImportService struct {
	pb.UnimplementedImportPluginServiceServer
	RedisClient *redis.Client
}

func (s *ImportService) Save(pluginName string, pluginBinary []byte) error {
	ctx := context.Background()
	status := s.RedisClient.Set(ctx, pluginName, pluginBinary, 0)
	_, err := status.Result()
	return err
}

func (s *ImportService) Import(ctx context.Context, request *pb.ImportPluginRequest) (*pb.ImportPluginResponse, error) {
	//temporarily save in filesystem - plugins folder
	tempDir, mkdirErr := os.MkdirTemp(".", "plugins-*")
	if mkdirErr != nil {
		log.Fatal("Error occurred creating a temporary directory", mkdirErr)
	}

	goFileName := path.Join(tempDir, request.Name+".go")
	goFile, fileErr := os.Create(goFileName)

	if fileErr != nil {
		log.Fatal("Error occurred creating a script file", fileErr)
	}

	goFile.Write(request.Contents)
	goFile.Close()
	defer os.Remove(goFileName)

	//use go build
	buildOutputName := path.Join(tempDir, request.Name+".so")
	exec.Command("go", "build", "-buildmode=plugin", "-o", buildOutputName, goFileName)
	defer os.Remove(buildOutputName)
	defer os.Remove(tempDir)

	//read .so file
	buildOutputBytes, readErr := os.ReadFile(buildOutputName)
	if readErr != nil {
		log.Fatal("Error occurred reading build output", readErr)
	}

	err := s.Save(request.Name, buildOutputBytes)
	return &pb.ImportPluginResponse{}, err
}
