// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish,
// you may prefer to use a [WaitGroup](waitgroups).

package main

import "fmt"
import "time"

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool, index int) {
	for {
		exited := false
		select {
		case msg := <-done:
			fmt.Println("received message - ", index, ", ", msg)
			exited = true
			break
		default:
			fmt.Println("working-", index)
			time.Sleep(1 * time.Second)
			fmt.Println("done-", index)
			//fmt.Println("no message received - ", index)
		}
		if exited {
			break
		}
	}
	fmt.Println("worker end.", index)
}

func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool, 0)
	go worker(done, 1)
	go worker(done, 2)
	go worker(done, 3)

	time.Sleep(2 * time.Second)

	// Block until we receive a notification from the
	// worker on the channel.
	done <- true
	done <- true
	done <- true

	fmt.Println("end")
}
