package main

import "fmt"

func fn1(x int) int {
	return x+1
}

// level 1
func adder() func(int) int {

	// level 2
	fn3 := func () func(int) int {

		// level 3
		return func(x int) int {

			// level 4
			fn4 := func () int {
				var i2 int = 1
				return i2
			}

			return fn4() + x
		}
		//return fn1
	}

	//sum := 0
	return fn3()
	//return func(x int) int {
	//	sum += x
	//	return sum
	//}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
