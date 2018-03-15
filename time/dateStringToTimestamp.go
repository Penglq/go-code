package main

import (
	"go-code/utils"
	"time"
	"fmt"
)

func main() {

	//t := time.Now()
	t, _ := time.ParseInLocation("2006-01-02", "2017-12-02", utils.TimeLocation)

	tAdd := t.AddDate(0,1,0)

	fmt.Println(tAdd.Format("2006-01"))


	//startDate := "2006-01-02"
	//tmpTime, err := time.ParseInLocation("2006-01-02", startDate, utils.TimeLocation)
	//if err != nil {
	//	utils.MyLog("loanLogPrefix time.ParseInLocation ", err)
	//}
	//endDate := time.Unix(tmpTime.Unix()+60*5, 0).Format("2006-01-02 15:04:05")
	//
	//fmt.Print(endDate)
}
