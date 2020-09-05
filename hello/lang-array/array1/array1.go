package main

import (
	"fmt"
	"sort"
)


func getName() (r string) {
	println("i am getName[array1]")
	r = "a"
	return
}

var fsName = getName()

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func HelloArray11() {
	arr1 := make([]int, 3, 10)
	printSlice(arr1)
	for i := 0; i < len(arr1); i++  {
		arr1[i] = i * 10
	}
	printSlice(arr1)
	for i := len(arr1); i < cap(arr1); i++  {
		arr1 = append(arr1, i * 11)
	}
	printSlice(arr1)
	ints := append(arr1, 1111)
	printSlice(ints)
	printSlice(arr1)
}

func HelloArray12() {
	arr1 := make([]int, 3, 10)
	printSlice(arr1)
	for i := 0; i < len(arr1); i++  {
		arr1[i] = i * 10
	}
	printSlice(arr1)
	for i := len(arr1); i < cap(arr1); i++  {
		arr1 = append(arr1, i * 11)
	}
	printSlice(arr1)
}

func testVariable1(v int) (r int) {
	HelloArray11()
	r = 0
	return
}
//var iTemp = utils.RegCallBack(testVariable1)

type StructA struct {
	i int
	j,k float64
	d int64
}

func helloSort1() {
	intList := [] int {2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := [] float64 {4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := [] string {"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)

}

func main() {
	HelloArray11()
}