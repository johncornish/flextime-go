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
	// println("****am listen")
	return nil
}

func main() {
	// tasks := []flextime.Task{
	// 	{
	// 		Name: "task-1",
	// 	},
	// }

	// http.HandleFunc("/v1/tasks/", func(w http.ResponseWriter, r *http.Request) {
	// 	b, _ := json.Marshal(tasks)
	// 	w.Write(b)
	// })

	// println("listening on port 8080")
	// http.ListenAndServe(":8080", nil)
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

// import (
// 	"flag"

// 	"net"

// 	"github.com/golang/glog"
// 	"golang.org/x/net/context"
// 	"google.golang.org/grpc"
// )

// type echoServer struct{}

// func newEchoServer() pb.EchoServiceServer {
// 	return new(echoServer)
// }

// func (s *echoServer) Echo(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
// 	glog.Info(msg)
// 	return msg, nil
// }

// func Run() error {
// 	listen, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		return err
// 	}
// 	server := grpc.NewServer()
// 	pb.RegisterEchoServiceServer(server, newEchoServer())
// 	server.Serve(listen)
// 	return nil
// }

// func main() {
// 	flag.Parse()
// 	defer glog.Flush()

// 	if err := Run(); err != nil {
// 		glog.Fatal(err)
// 	}
// }
