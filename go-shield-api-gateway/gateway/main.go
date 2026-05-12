package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"api-gateway-project/queue"
)

var services = []string{
	"http://localhost:8001",
	"http://localhost:8002",
	"http://localhost:8003",
}

var counter uint64

// 🔹 Load Balancer (Round Robin)
func getNextService() string {
	idx := atomic.AddUint64(&counter, 1)
	return services[idx%uint64(len(services))]
}

// 🔹 Rate Limiter (Token Bucket)
var tokens = make(chan struct{}, 100)

func init() {
	for i := 0; i < 100; i++ {
		tokens <- struct{}{}
	}

	// refill tokens every second
	go func() {
		for {
			time.Sleep(time.Second)
			for len(tokens) < 100 {
				tokens <- struct{}{}
			}
		}
	}()
}

// 🔹 Request Handler
func handler(w http.ResponseWriter, r *http.Request) {
	select {
	case <-tokens:
		service := getNextService()

		// simulate overload
		if len(tokens) < 10 {
			fmt.Println("Queueing request...")
			queue.JobQueue <- queue.Job{
				URL: service,
				W:   w,
				R:   r,
			}
			return
		}

		resp, err := http.Get(service)
		if err != nil {
			http.Error(w, "Service unavailable", 500)
			return
		}
		defer resp.Body.Close()

		fmt.Fprintf(w, "Direct response from %s", service)

	default:
		http.Error(w, "Rate limit exceeded", 429)
	}
}

func main() {
	queue.StartWorker()

	http.HandleFunc("/", handler)
	fmt.Println("API Gateway running on :8080")
	http.ListenAndServe(":8080", nil)
}
