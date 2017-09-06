package main

import "fmt"

type human struct {
	name string
}

func (h *human) say()  {
	fmt.Println("human")
}

type student struct {
	human
}

func (s *student) say()  {
	fmt.Println("student")
}

func main() {
	human := human{name:"jack"}
	human.say()

	student := student{}
	student.say()
}
