package main

func main() {

	s := make(chan int, 1)
	s <- 1
	L:for {
		select {
		case <-s:
			break L
		}
	}
}
