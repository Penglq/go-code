package main

import (
	"fmt"
	"time"
	"log"
)

//func main() {
//	m := make(map[int]int, 2)
//
//	m[0] = 1
//	m[2] = 2
//	m[3] = 3
//	m[4] = 4
//	fmt.Println(m)
//	TimeLocation, _ := time.LoadLocation("Asia/Chongqing") //当地时间
//	log.Print(time.Now().In(TimeLocation).Unix())
//}

func sliceReturn(s ...int) {
	for _, v := range s {
		fmt.Println(v)
	}
}
