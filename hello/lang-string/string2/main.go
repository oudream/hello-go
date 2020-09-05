package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a long string with many repeated characters it wor myit.."
	numberOfa := strings.Count(str, "it")

	fmt.Printf("[%v] string has %d of characters of [a] ", str, numberOfa)
}
