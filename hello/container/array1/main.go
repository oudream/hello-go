package main

import (
	"fmt"
	"sort"
	"time"
)

func getName() (r string) {
	println("i am getName[array1]")
	r = "a"
	return
}

var fsName = getName()

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func helloArrayLoop1()  {
	intArray := [5]int{10, 20, 30, 40, 50}
	fmt.Println("\n---------------Example 1--------------------\n")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}
	fmt.Println("\n---------------Example 2--------------------\n")
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}
	fmt.Println("\n---------------Example 3--------------------\n")
	for _, value := range intArray {
		fmt.Println(value)
	}
	j := 0
	fmt.Println("\n---------------Example 4--------------------\n")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}

func HelloArray11() {
	arr1 := make([]int, 3, 10)
	printSlice(arr1)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 10
	}
	printSlice(arr1)
	for i := len(arr1); i < cap(arr1); i++ {
		arr1 = append(arr1, i*11)
	}
	printSlice(arr1)
	ints := append(arr1, 1111)
	printSlice(ints)
	printSlice(arr1)
}

func HelloArray12() {
	arr1 := make([]int, 3, 10)
	printSlice(arr1)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 10
	}
	printSlice(arr1)
	for i := len(arr1); i < cap(arr1); i++ {
		arr1 = append(arr1, i*11)
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
	i    int
	j, k float64
	d    int64
}

func helloSort1() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)

}

type UserPwd struct {
	Username string
	Password string
}

func helloArray1() {
	userPwds := []UserPwd{
		{"ua", "p1"},
		{"ub", "p2"},
		{"uc", "p3"},
	}
	fmt.Println(userPwds)
}

func getDayIndex() int {
	now := time.Now()
	return (now.Year() + int(now.Month()) + now.Day()) % 7
}

func pringArray1(a []int)  {
	for i,x:= range a {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
}

func main() {
	helloSort1()
}
