package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println("结果", path.Base("/a/b/c"))
}
