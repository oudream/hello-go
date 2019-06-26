package main

import (
	"fmt"
)

type name struct {
	name string
	age  int
	height  float64
}

func main() {
	c := make(chan name)

	go func() {
		c <- name{"sfsaf", 1, 1.73}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

	fmt.Println("channel was closed (all done!).")
}