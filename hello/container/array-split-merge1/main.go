package main

import (
	"fmt"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func helloSplit1() {
	arr := []int{1, 2, 3, 5, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9}
	limit := 5

	for i := 0; i < len(arr); i += limit {
		batch := arr[i:min(i+limit, len(arr))]
		fmt.Println(batch)
	}

}

func helloMerge1()  {
	var arr1 []int
	arr1 = append(arr1, []int{5,6}...)
	arr1 = append(arr1, []int{1,2}...)
	arr1 = append(arr1, []int{3,4}...)
	fmt.Println(arr1)
}

func main() {
	//helloSplit1()
	helloMerge1()
}