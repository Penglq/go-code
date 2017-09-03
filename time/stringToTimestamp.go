package main

import (
	"fmt"
	"time"
)

func main() {
	s, _ := time.LoadLocation("Asia/Chongqing") //当地时间
	regTimestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-08-07 00:00:00", s)
	fmt.Print(regTimestamp)

	fmt.Print(GetMillisecondTime(`2016-01-02 15:04:05`))
}

// 将date字符串（2017-01-02 10:02:20）转换成毫秒
func GetMillisecondTime(date string) int64 {
	tm, _ := time.Parse("2006-01-02 15:04:05", date)
	return tm.UnixNano() / 1e6
}
