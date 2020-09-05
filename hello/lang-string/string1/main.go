package main

import (
	"bytes"
	"fmt"
	"strings"
)
import "os"

type point struct {
	x, y int
}

func format1() {
//  {1 2}
	p := point{1, 2}
	fmt.Printf("%v\n", p)

//  {x:1 y:2}
	fmt.Printf("%+v\n", p)

//  main.point{x:1, y:2}
	fmt.Printf("%#v\n", p)

//  main.point
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", 123141235634656)
	fmt.Printf("%T\n", 1e16)

//  true
	fmt.Printf("%t\n", true)

//  123
	fmt.Printf("%d\n", 123)

//  1110
	fmt.Printf("%b\n", 14)

//  !
	fmt.Printf("%c\n", 33)

//  1c8
	fmt.Printf("%x\n", 456)

//  78.900000
	fmt.Printf("%f\n", 78.9)

//  1.234000e+08
//  1.234000E+08
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

//  "string"
	fmt.Printf("%s\n", "\"string\"")

//  "\"string\""
	fmt.Printf("%q\n", "\"string\"")

//  6865782074686973
	fmt.Printf("%x\n", "hex this")

//  0xc000082010
	fmt.Printf("%p\n", &p)

//  |    12|   345|
	fmt.Printf("|%6d|%6d|\n", 12, 345)

//  |  1.20|  3.45|
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

//  |1.20  |3.45  |
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

//  |   foo|     b|
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

//  |foo   |b     |
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

//  a string
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

//  an error
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

func helloPrint2() {
	fmt.Print(`line 1
line 2
line 3`)
}

func helloJion1() {
	s := "1231"
	fmt.Print(strings.Join([]string{s,"ewrwer"},"--"))
}

func helloJion2() {
	var s []string
	s = append(s,"123123")
	ss := strings.Join(s,"")
	fmt.Print(ss)
}

func helloJion3()  {
	var buf bytes.Buffer
	buf.WriteString("12312")
	buf.WriteString("werwer")
	buf.String()
}

func main() {
	//format1()
	//helloPrint2()
	helloJion2()
}
