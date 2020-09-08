
package main

import (
	"fmt"
	"strconv"
)

var chan2 = make(chan interface{})

func routineSend1(ch chan<- string, index int64) {
	//time1.Sleep(time1.Second)
	ch <- "he: " + strconv.FormatInt(index, 10)
	if index < 5 {
		chan2 <- index
	} else {
		chan2 <- nil
	}
}

func routineRecv1(ch <-chan string, index int64) {
	//time1.Sleep(time1.Second)
	res := <-ch
	fmt.Println(res)
	chan2 <- nil
}


func hello1() {
	chan1 := make(chan string)
	var i int64 = 0
	for i = 0; i<10; i++ {
		go routineSend1(chan1, i)
	}
	for i = 0; i<10; i++ {
		go routineRecv1(chan1, i)
	}
	for i := 0; i < 20; i++ {
		fmt.Println(<-chan2)
	}
	fmt.Println("end.")
}

func main() {
	hello1()
}