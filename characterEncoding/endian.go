package main

import (
	"encoding/binary"
	"fmt"
)

func BigEndian() {

	// 0000 0000 0000 0000 0001 0002 0003 0004
	var testInt int32 = 0x01020304
	fmt.Printf("%d use big endian: \n", testInt)
	var testBytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func LittleEndian() {

	// 0000 0000 0000 0000 0001 0002 0003 0004
	var testInt int32 = 0x01020304
	fmt.Printf("%d use little endian: \n", testInt)
	var testBytes []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.LittleEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func main() {
	BigEndian()
	LittleEndian()
}
