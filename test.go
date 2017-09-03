package main

import (
	"log"
	"time"
	"strconv"
)

func main() {
	var s string

	log.Println(s, strconv.FormatInt(time.Now().Unix(),10))

	log.Println(time.Unix(1504084457, 0).Format("2006-01-02 15:04:05"))
}
