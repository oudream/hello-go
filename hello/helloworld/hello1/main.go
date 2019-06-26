package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println( i)
	fmt.Println("Hello World")
	fmt.Println( "\u65e5")
	//noteFrequency := map[string]float32 {
	//	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	//	"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
	//fmt.Print(noteFrequency)
	var num float64 = 1.2345

	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))

	const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
	const b = 15 / 4           // b == 3     (untyped integer constant)
	const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
	const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
	const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
	const d = 1 << 3.0         // d == 8     (untyped integer constant)
	const e = 1.0 << 3         // e == 8     (untyped integer constant)
	//const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
	//const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
	const h = "foo" > "bar"    // h == true  (untyped boolean constant)
	const j = true             // j == true  (untyped boolean constant)
	const k = 'w' + 1          // k == 'x'   (untyped rune constant)
	const l = "hi"             // l == "hi"  (untyped string constant)
	const m = string(k)        // m == "x"   (type string)
	const Σ = 1 - 0.707i       //            (untyped complex constant)
	const Δ = Σ + 2.0e-4       //            (untyped complex constant)
	const Φ = iota*1i - 1/1i   //            (untyped complex constant)

	fmt.Println("type: ", reflect.TypeOf(a))
	fmt.Println("type: ", reflect.TypeOf(d))
	fmt.Println("type: ", reflect.TypeOf(e))
	fmt.Println("type: ", reflect.TypeOf(h))
	fmt.Println("type: ", reflect.TypeOf(k))
	fmt.Println("type: ", reflect.TypeOf(Σ))
}
