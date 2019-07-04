package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a byte
		b bool
		s uint64
	}
	fmt.Println(unsafe.Sizeof(x))
}
