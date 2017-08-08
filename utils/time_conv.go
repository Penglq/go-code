package utils

import (
	"log"
	"time"
)

var (
	TimeLocation, _ = time.LoadLocation("Asia/Chongqing") //当地时间
)

func GetDateFromTime(t int64) string {
	if t > 0 {
		if t/1e9 >= 1000 {
			//毫秒
			t = int64(t / 1000)
		}
		return time.Unix(t, 0).In(TimeLocation).Format("2006-01-02")
	} else {
		return ""
	}
	//loc, _ := time.LoadLocation("UTC")

}
func GetDateTimeFromTime(t int64) string {
	if t > 0 {
		if t/1e9 >= 1000 {
			//毫秒
			t = int64(t / 1000)
		}
		return time.Unix(t, 0).In(TimeLocation).Format("2006-01-02 15:04:05")
	} else {
		return ""
	}
	//loc, _ := time.LoadLocation("UTC")

}

// GetNowDateTime default layout 2006-01-02 15:04:05
func GetNowDateTime(layout string) string {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return time.Now().In(TimeLocation).Format(layout)
}

func GetLocalUnixTimestamp() int64 {
	return time.Now().In(TimeLocation).Unix()
}

//GetHourOfDay 计算所给时间是一天中第几小时,t为秒或者毫秒
func GetHourOfDay(t int64) int {
	//loc, _ := time.LoadLocation("UTC")
	var hour int = 0
	if t > 0 {
		if t/1e9 >= 1000 {
			//毫秒
			t = int64(t / 1000)
		}
		//TimeLocation, _ = time.LoadLocation("Asia/Chongqing")
		unixtime := time.Unix(t, 0).In(TimeLocation)
		hour = unixtime.Hour()
	}

	return hour
}

//计算所给时间是一天中第几个5分钟,t为秒或者毫秒
func GetFiveMinuteSequenceOfDay(t int64) int {
	var value int = 0
	if t > 0 {
		//loc, _:= time.LoadLocation("UTC")
		if t/1e9 >= 1000 {
			//毫秒
			t = int64(t / 1000)
		}

		//TimeLocation, _ = time.LoadLocation("Asia/Chongqing")
		unixTime := time.Unix(t, 0).In(TimeLocation)
		d, _ := time.ParseInLocation("2006-01-02", unixTime.Format("2006-01-02"), TimeLocation)
		//当天0时的时间戳
		t1 := d.Unix()
		value = int((t - t1) / 5 / 60)
	}

	return value
}

func MyLog(args ...interface{}) {
	log.Printf("%+v", args)
}
