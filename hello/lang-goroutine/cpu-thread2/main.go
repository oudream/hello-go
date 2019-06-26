package main

import (
	"fmt"
	"math"
	"time"
)

func say(i int) {
	dtNow := time.Now().UnixNano()
	for k := 0; k < i; k++ {
		var j int64 = 0
		for ; j<math.MaxInt64; j++ {
			if time.Now().UnixNano() - dtNow > int64(5 * time.Second) {
				dtNow = time.Now().UnixNano()
				fmt.Printf(" -- %d -- %s\n", i, time.Now().String())
			}
		}
	}
}

func main() {
	for i := 960; i < 999; i++ {
		go say(i)
	}
	say(999)
}