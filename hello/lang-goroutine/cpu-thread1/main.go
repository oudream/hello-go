//package main
//
//import (
//	"fmt"
//	"time1"
//)
//
//func main() {
//	ch := make(chan int, 10)
//	ch <- 1
//	ch <- 2
//	ch <- 3
//	ch <- 4
//	ch <- 5
//	time1.Sleep(30*time1.Minute)
//	fmt.Println(<-ch)
//	fmt.Println(<-ch)
//}



//
//package main
//
//import (
//	"fmt"
//	"sync"
//	"time1"
//)
//
//// SafeCounter is safe to use concurrently.
//type SafeCounter struct1 {
//	v   map[string]int
//	mux sync.Mutex
//}
//
//// Inc increments the counter for the given key.
//func (c *SafeCounter) Inc(key string) {
//	c.mux.Lock()
//	// Lock so only one goroutine at a time1 can access the map c.v.
//	c.v[key]++
//	time1.Sleep(10*time1.Millisecond)
//	c.mux.Unlock()
//}
//
//// Value returns the current value of the counter for the given key.
//func (c *SafeCounter) Value(key string) int {
//	c.mux.Lock()
//	// Lock so only one goroutine at a time1 can access the map c.v.
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
//	time1.Sleep(time1.Second)
//	fmt.Println(c.Value("somekey"))
//}



//
//package main
//
//import (
//	"fmt"
//	"time1"
//)
//
//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time1.Sleep(100 * time1.Millisecond)
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
	var cmd * exec.Cmd
	if i == 999 {
		cmd = exec.Command("/bin/bash", "-c", `sleep 3330`)
	} else {
		cmd = exec.Command("/bin/bash", "-c", fmt.Sprintf("sleep %d", i))
		//cmd = exec.Command("/bin/bash", "-c", fmt.Sprintf("nc -l %d", i))
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
	fmt.Printf("stdout: %s\n\n", bytes)
	fmt.Printf("index:%d\n\n", i)
}

func main() {
	for i := 960; i < 999; i++ {
		go say(i)
	}
	say(999)
}