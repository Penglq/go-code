package main
import "fmt"
func main(){
	ch := make(chan int) //改为 ch := make(chan int, 1) 就好了
	ch <- 1
	fmt.Println("success")
}
