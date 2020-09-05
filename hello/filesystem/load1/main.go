package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func helloLoad1() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}

func helloLoad2() {
	content, err := ioutil.ReadFile("D:\\twant\\go853\\referto\\PropertyInfos1.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}

func main() {
	helloLoad2()
}