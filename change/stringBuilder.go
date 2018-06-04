package change

import (
	"unsafe"
)

func Builder() {
	b := []byte{'a', 'b'}
	_ = *(*string)(unsafe.Pointer(&b))
}

func SBuilder() {
	b := []byte{'a', 'b'}
	_ = string(b)
}
