package main

import (
	"fmt"
	"time"
)

func main() {
	s, _ := time.LoadLocation("Asia/Chongqing") //当地时间
	regTimestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-07 00:00:00", s)
	fmt.Print(regTimestamp.Unix())
}
