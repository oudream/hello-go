package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	str := "hello, let's go"
	upper := flag.Int("u", 0, "the upper bound index")
	lower := flag.Int("l", 0, "the lower bound index")

	flag.Parse()

	if *upper <= *lower {
		fmt.Printf("%v",
			fmt.Errorf("invalid arguments for upper %d and lower %d", *upper, *lower))
		os.Exit(100)
	}

	os.Exit(1)

	fmt.Printf("%s\n", str[*(lower):(*upper)])
}