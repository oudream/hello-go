package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// A struct with a mix of fields, used for the GOB example.
type complexData struct {
	N int
	S string
	M map[string]int
	P []byte
	C *complexData
	E Addr
}

type Addr struct {
	Comment string
}

func main() {

	testStruct := complexData{
		N: 23,
		S: "string data",
		M: map[string]int{"one": 1, "two": 2, "three": 3},
		P: []byte("abc"),
		C: &complexData{
			N: 256,
			S: "Recursive structs? Piece of cake!",
			M: map[string]int{"01": 1, "10": 2, "11": 3},
			E: Addr{
				Comment: "InnerTest123123123123",
			},
		},
		E: Addr{
			Comment: "Test123123123",
		},
	}

	fmt.Println("Outer complexData struct: ", testStruct)
	fmt.Println("Inner complexData struct: ", testStruct.C)
	fmt.Println("Inner complexData struct: ", testStruct.E)
	fmt.Println("===========================")

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(testStruct)
	if err != nil {
		fmt.Println(err)
	}

	dec := gob.NewDecoder(&b)
	var data complexData
	err = dec.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding GOB data:", err)
		return
	}

	fmt.Println("Outer complexData struct: ", data)
	fmt.Println("Inner complexData struct: ", data.C)
	fmt.Println("Inner complexData struct: ", testStruct.E)

}
