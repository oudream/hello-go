package main

import (
	"container/list"
	"fmt"
)

func main() {
	ls := list.New()
	ls.PushBack(1)
	ls.PushBack(2)

	fmt.Printf("len: %v\n", ls.Len())
	fmt.Printf("first: %#v\n", ls.Front())
	fmt.Printf("second: %#v\n", ls.Front().Next())
}