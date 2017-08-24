package label

import "fmt"

func main() {
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case tmp := <-ch:
			fmt.Print(tmp)
			goto End
		case ch <- i:
		}
	}
End:fmt.Print("hello")
}
