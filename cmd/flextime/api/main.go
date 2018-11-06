package main

import (
	"encoding/json"
	"net/http"

	"github.com/johncornish/flextime-go/pkg/flextime"
)

type taskResponse struct {
	Name string `json:"name"`
}

func main() {
	tasks := []flextime.Task{
		{
			Name: "task-1",
		},
	}

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(tasks)
		w.Write(b)
	})

	println("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
