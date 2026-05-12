package queue

import (
	"fmt"
	"net/http"
)

type Job struct {
	URL string
	W   http.ResponseWriter
	R   *http.Request
}

var JobQueue = make(chan Job, 1000)

func StartWorker() {
	go func() {
		for job := range JobQueue {
			resp, err := http.Get(job.URL)
			if err != nil {
				http.Error(job.W, "Service error", 500)
				continue
			}
			defer resp.Body.Close()

			fmt.Fprintf(job.W, "Processed from queue: %s", job.URL)
		}
	}()
}
