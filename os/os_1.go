package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Print(runtime.NumCPU())
	fmt.Print(runtime.GOOS)
}
