package main

import (
	"fmt"
	"time"
)

func HelloChanTimeOut11() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "result"
	}()
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}

func testVariable1(v int) (r int) {
	HelloChanTimeOut11()
	r = 0
	return
}
//var iTemp = utils.RegCallBack(testVariable1)
