package main

import (
	"time"
	"log"
)

func main() {

	TimeLocation, _ := time.LoadLocation("Asia/Chongqing") //当地时间
	layout := "2006-01-02 15:04:05"

	log.Print(time.Now().In(TimeLocation).Format(layout))

	log.Print(time.Now().Unix())

	test("a", "b")
}

func test(a... string)  {
	log.Print(a[0])
}