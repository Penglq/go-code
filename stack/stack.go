package main

func main() {
	a, b := 1, 2
	_ = add1(a, b)
	_ = add2(a, b)
}
func add1(x, y int) int {
	return x + y
}
func add2(x, y int) int {
	_ = make([]byte, 200)
	return x + y
}
