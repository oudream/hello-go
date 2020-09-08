package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	for i := 1; i <= 5; i++ {
		go func (i int, co chan<- string) {
			for j := 1; j <= 5; j++ {
				co <- fmt.Sprintf("hi from %d.%d", i, j)
				time.Sleep(time.Second)
			}
		}(i, c)
	}

	for i := 1; i <= 25; i++ {
		fmt.Println(<-c)
	}
}


//
//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func routine1(ch chan string, index int64) {
//	//time1.Sleep(time1.Second)
//	ch <- "he: "
//	ch <- strconv.FormatInt(index, 10)
//}
//
//
//func hello1() {
//	ch := make(chan string, 110)
//	var i int64 = 0
//	for i = 0; i<10; i++ {
//		go routine1(ch, i)
//	}
//	for i = 0; i<10; i++ {
//		res := <-ch
//		fmt.Println(res)
//	}
//	fmt.Println("end.")
//}
//
//func main() {
//	hello1()
//}