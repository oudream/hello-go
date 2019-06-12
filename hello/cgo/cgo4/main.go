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

void PrintPacked(packed p){
	printf("From C\na:%d\nb:%d\nc:%d\nd:%d\ne:%s\n", p.a, p.b, p.c, p.d, p.e);
}

*/
import "C"
import (
	"bytes"
	"encoding/binary"
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

//Pack Produces a packed version of the go struct
func (g *GoPack) Pack(out unsafe.Pointer) {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, g.a)
	binary.Write(buf, binary.LittleEndian, g.b)
	binary.Write(buf, binary.LittleEndian, g.c)
	binary.Write(buf, binary.LittleEndian, g.d)
	binary.Write(buf, binary.LittleEndian, g.e)
	//Getting the lenfth of memory
	l := buf.Len()
	//Cast the point to byte slie to allow for direct memory manipulation
	o := (*[1 << 20]C.uchar)(out)
	//Write to memory
	for i := 0; i < l; i++ {
		b, _ := buf.ReadByte()
		o[i] = C.uchar(b)
	}
}

func main() {
	pack := &GoPack{1, 2, 3, 4, [10]byte{}}
	copy(pack.e[:], "TEST123")
	cpack := C.packed{} //just to allocate the memory, still under GC control
	pack.Pack(unsafe.Pointer(&cpack))
	C.PrintPacked(cpack)
}