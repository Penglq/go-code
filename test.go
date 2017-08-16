package main

import (
	"log"
	"unsafe"
)

func main() {
	s := new([]int64)

	log.Println(unsafe.Sizeof(s))
}
