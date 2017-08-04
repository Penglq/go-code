package main

import (
	"fmt"
	"time"
)

func main() {
	s, _ := time.LoadLocation("Asia/Chongqing") //当地时间
	regTimestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-04 10:38:20", s)
	fmt.Print(regTimestamp.Unix())
}
