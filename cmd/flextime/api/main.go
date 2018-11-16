package main

import (
	"context"
	"log"
	"net"

	v1 "github.com/johncornish/flextime-go/rpc/flextime_v1"
	"google.golang.org/grpc"
)

type taskResponse struct {
	Name string `json:"name"`
}

type taskServer struct{}

func newTaskServer() v1.RESTServer {
	return new(taskServer)
}

func (s *taskServer) AddTask(ctx context.Context, task *v1.Task) (*v1.AddTaskResponse, error) {
	return new(v1.AddTaskResponse), nil
}

func Run() error {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterRESTServer(server, newTaskServer())
	server.Serve(listen)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
