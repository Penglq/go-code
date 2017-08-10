package main

import "fmt"

func intSeq ()func()int{
	i:=0
	return func()int{
		i+=1
		return i
	}
}

//func main() {
//	initNext := intSeq()
//	fmt.Println(initNext())
//	fmt.Println(initNext())
//	fmt.Println(intSeq())
//}
