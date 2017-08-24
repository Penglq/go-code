// label可以定义在函数体的任何地方，并且不可重复
package main

import (
	"fmt"
	"time"
)

func main() {
End:
	fmt.Println(1)
	time.Sleep(10 * time.Second)
	goto End
	fmt.Println(2)
	fmt.Println(3)
}
