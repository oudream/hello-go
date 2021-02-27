package main

import (
	"container/ring"
	"fmt"
)

func main() {
	rg := ring.New(3)

	for i := 1; i <= 3; i++ {
		rg.Value = i
		rg = rg.Next()
	}

	// 计算 1+2+3
	s := 0
	rg.Do(func(p interface{}){
		s += p.(int)
	})
	fmt.Println("sum is", s)
}
