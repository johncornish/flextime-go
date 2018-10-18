package main_test

import (
	"encoding/json"
	"net/http"
	"os/exec"

	"github.com/johncornish/flextime-go/pkg/flextime"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
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

	It("serves tasks", func() {
		Eventually(func() int {
			resp, err := http.Get("http://127.0.0.1:8080/tasks/")
			if err != nil {
				return -1
			}

			return resp.StatusCode
		}).Should(Equal(http.StatusOK))
	})

	It("returns an array of tasks", func() {
		tasks := []flextime.Task{
			{
				Name: "task-1",
			},
		}

		Eventually(func() []flextime.Task {
			// _, err := http.PostForm(
			// 	"http://127.0.0.1:8080/tasks/",
			// 	url.Values{
			// 		"name": {
			// 			"task-1",
			// 		},
			// 		// "estimate": "15m",
			// 	},
			// )
			// if err != nil {
			// 	return []flextime.Task{}
			// }

			return getTasksFromJson("http://127.0.0.1:8080/tasks/")
		}).Should(Equal(tasks))
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
	path, err := gexec.Build("github.com/johncornish/flextime-go/cmd/flextime")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(path)
	// cmd.Env = envs
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	if err != nil {
		panic(err)
	}

	return session
}
