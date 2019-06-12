package main

/*
#include "stdio.h"
#pragma pack(1)
typedef struct{
	unsigned char a;
	char b;
	int c;
	unsigned int d;
	char e[10];
}packed;

packed PackedInit(){
	packed p;
	p.a = 1;
	p.b = 2;
	p.c = 3;
	p.d = 4;
	p.e[0] = 'T';
	p.e[1] = 'E';
	p.e[2] = 'S';
	p.e[3] = 'T';
	p.e[4] = '1';
	p.e[5] = '2';
	p.e[6] = '3';
	p.e[7] = '\0';
	p.e[8] = '\0';
	p.e[9] = '\0';
	return p;
}

*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

//GoPack is the go version of the c packed structure
type GoPack struct {
	a uint8
	b int8
	c int32
	d uint32
	e [10]uint8
}

func (g *GoPack) Unpack(i *C.packed) {
	cdata := C.GoBytes(unsafe.Pointer(i), C.sizeof_packed)
	buf := bytes.NewBuffer(cdata)
	binary.Read(buf, binary.LittleEndian, &g.a)
	binary.Read(buf, binary.LittleEndian, &g.b)
	binary.Read(buf, binary.LittleEndian, &g.c)
	binary.Read(buf, binary.LittleEndian, &g.d)
	binary.Read(buf, binary.LittleEndian, &g.e)
}

func main() {
	packed := C.PackedInit()
	gopacked := GoPack{}
	gopacked.Unpack(&packed)
	fmt.Printf("a: %d\nb: %d\nc: %d\nd: %d\ne: %s\n", gopacked.a, gopacked.b, gopacked.c, gopacked.d, string(gopacked.e[:]))
}