package main_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os/exec"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/johncornish/flextime-go/pkg/flextime"
	v1 "github.com/johncornish/flextime-go/rpc/flextime_v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"google.golang.org/grpc"
)

var _ = Describe("Main", func() {
	var (
		session *gexec.Session
	)

	BeforeSuite(func() {
		session = startFlextime()
	})

	AfterSuite(func() {
		session.Kill()
	})

	It("receives tasks and responds", func() {
		taskResps := make(chan *v1.AddTaskResponse, 100)
		var taskResp *v1.AddTaskResponse

		task := v1.Task{
			Name:     "task-1",
			Estimate: "15m",
			Repeat:   "",
			Due: &timestamp.Timestamp{
				Seconds: 1542295067,
			},
		}

		addr := "127.0.0.1:50051"
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		Expect(err).ToNot(HaveOccurred())

		restClient := v1.NewRESTClient(conn)
		resp, err := restClient.AddTask(context.Background(), &task)
		taskResps <- resp

		Expect(taskResps).To(Receive(&taskResp))
	})
})

func getTasksFromJson(url string) []flextime.Task {
	t := []flextime.Task{}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&t)
	return t
}

func startFlextime() *gexec.Session {
	path, err := gexec.Build("github.com/johncornish/flextime-go/cmd/flextime/api")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(path)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	if err != nil {
		panic(err)
	}

	return session
}
