package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int16 = 0x0102
	firstByte := *(*byte)(unsafe.Pointer(&i))
	if firstByte == 1 {
		fmt.Println("BigEndian")
	} else {
		fmt.Println("LittleEndian")
	}
}
