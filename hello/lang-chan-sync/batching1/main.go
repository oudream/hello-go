package main

import (
	"fmt"
	"github.com/eapache/channels"
	"time"
)

func testBatches(ch channels.Channel) {
	go func() {
		for i := 0; i < 1000; i++ {
			ch.In() <- "name"+string(i)
		}
		ch.Close()
	}()

	i := 0
	for val := range ch.Out() {
		for _, elem := range val.([]interface{}) {
			if i != elem.(int) {
				fmt.Println("batching channel expected", i, "but got", elem.(int))
			}
			i++
		}
		fmt.Println(i)
	}
	fmt.Println("batching channel In <- 1000 : <- Out : ", i)
}

func testChannelConcurrentAccessors(name string, ch channels.Channel) {
	// no asserts here, this is just for the race detector's benefit
	go ch.Len()
	go ch.Cap()

	go func() {
		ch.In() <- nil
	}()

	go func() {
		<-ch.Out()
	}()
}

func TestBatchingChannel() {
	ch := channels.NewBatchingChannel(channels.Infinity)
	testBatches(ch)

	ch = channels.NewBatchingChannel(2)
	testBatches(ch)

	ch = channels.NewBatchingChannel(1)
	testChannelConcurrentAccessors("batching channel", ch)
}

func TestBatchingChannelCap() {
	ch := channels.NewBatchingChannel(channels.Infinity)
	if ch.Cap() != channels.Infinity {
		fmt.Errorf("incorrect capacity on infinite channel")
	}

	ch = channels.NewBatchingChannel(5)
	if ch.Cap() != 5 {
		fmt.Errorf("incorrect capacity on infinite channel")
	}
}

func main() {
	fmt.Println(`begin: ` + time.Now().String())
	TestBatchingChannel()
	TestBatchingChannelCap()
	fmt.Println(`end. ` + time.Now().String())
	for {
		time.Sleep(5 * time.Second)
		fmt.Println(`--end.`)
		break
	}
}