package main

import (
	"time"
	"log"
)

//57 65 20 63 6c 6f 73 65 64 20 6f 75 72 20 24 35 30 4d 20 73 65 72 69 65 73 20 43 20 66 75 6e 64 69 6e 67 2e 20 54 68 61 6e 6b 20 79 6f 75 20 61 6c 6c 2e 20
func main() {

	//bs := []byte()
	//fmt.Println(bs)

	TimeLocation, _ := time.LoadLocation("Asia/Chongqing") //当地时间
	layout := "2006-01-02 15:04:05"

	log.Print(time.Now().In(TimeLocation).Format(layout))

	log.Print(time.Now().Unix())

	test("a", "b")
}

func test(a ... string) {
	log.Print(a[0])
}
