//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	ch := make(chan int, 10)
//	ch <- 1
//	ch <- 2
//	ch <- 3
//	ch <- 4
//	ch <- 5
//	time.Sleep(30*time.Minute)
//	fmt.Println(<-ch)
//	fmt.Println(<-ch)
//}

//
//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//// SafeCounter is safe to use concurrently.
//type SafeCounter struct {
//	v   map[string]int
//	mux sync.Mutex
//}
//
//// Inc increments the counter for the given key.
//func (c *SafeCounter) Inc(key string) {
//	c.mux.Lock()
//	// Lock so only one goroutine at a time can access the map c.v.
//	c.v[key]++
//	time.Sleep(10*time.Millisecond)
//	c.mux.Unlock()
//}
//
//// Value returns the current value of the counter for the given key.
//func (c *SafeCounter) Value(key string) int {
//	c.mux.Lock()
//	// Lock so only one goroutine at a time can access the map c.v.
//	defer c.mux.Unlock()
//	return c.v[key]
//}
//
//func main() {
//	c := SafeCounter{v: make(map[string]int)}
//	for i := 0; i < 1000; i++ {
//		go c.Inc("somekey")
//	}
//
//	time.Sleep(time.Second)
//	fmt.Println(c.Value("somekey"))
//}

//
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}
//
//func main() {
//	for i := 0; i < 1000; i++ {
//		go say("world")
//	}
//	say("hello")
//}






package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func say(i int) {
	//cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("sleep %d", i))
	var cmd * exec.Cmd
	if i == 999 {
		cmd = exec.Command("/bin/bash", "-c", `sleep 77`)
	} else {
		cmd = exec.Command("/bin/bash", "-c", `sleep 35`)
	}
	//cmd := exec.Command("/bin/bash", "-c", `df -lh`)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	fmt.Printf("stdout:\n\n %s", bytes)
	fmt.Printf("index:\n\n %d", i)
}

func main() {
	for i := 0; i < 20; i++ {
		go say(i)
	}
	say(999)
}