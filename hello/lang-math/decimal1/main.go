package main

import (
	"fmt"
	"strconv"
)

func maDecimal2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func helloDecimal1() {
	f1 := 4890000.00
	a1 := 420.00
	fmt.Println(maDecimal2(f1 / a1))
}

func helloMod1() {
	i := 7
	j := 8
	fmt.Println(i % 2)
	fmt.Println(i % 2 == 0)
	fmt.Println(j % 2)
}

func main() {
	helloMod1()
}
