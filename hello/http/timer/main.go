package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	timeSumsMu sync.RWMutex
	timeSums   int64
)

func main() {
	// Start the goroutine that will sum the current time
	// once per second.
	go runDataLoop()
	// Create a handler that will read-lock the mutext and
	// write the summed time to the client
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timeSumsMu.RLock()
		defer timeSumsMu.RUnlock()
		fmt.Fprint(w, timeSums)
	})
	http.ListenAndServe(":8080", nil)
}

func runDataLoop() {
	for {
		// Within an infinite loop, lock the mutex and
		// increment our value, then sleep for 1 second until
		// the next time we need to get a value.
		timeSumsMu.Lock()
		time.Sleep(10 * time.Second)
		timeSums += time.Now().Unix()
		timeSumsMu.Unlock()
		time.Sleep(1 * time.Second)
	}
}