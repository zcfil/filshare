package utils

import (
	"time"
)

const (
	GOTIME    = "2006-01-02 15:04:05"
	GOTIMEDAY = "2006-01-02"
	MONTH     = "2006-01"

	GOTIMEDAY1 = "20060102"
	GOTIME1    = "20060102150405"
	MONTH1     = "200601"

	FILTIME  = "2020-08-25 06:00:00"
	FILCHECK = 30
	UTC_8    = int64(28800)
	DayTUNX  = 86400
	DAYTIME  = 15552000 // 180天
)

func Timedatetimes(height int64) string {
	times, _ := time.Parse(GOTIME, FILTIME) //2020-08-25 14:00:00 东八区时区  28800
	tsum := height*FILCHECK + times.Unix() - UTC_8
	timeStr := time.Unix(tsum, 0).Format(GOTIME)
	return timeStr
}
func TimedateMonth(height int64) string {
	times, _ := time.Parse(GOTIME, FILTIME) //2020-08-25 14:00:00 东八区时区  28800
	tsum := height*FILCHECK + times.Unix() - UTC_8
	timeStr := time.Unix(tsum, 0).Format(MONTH)
	return timeStr
}

func TimedateTimes(height int64) time.Time {
	times, _ := time.Parse(GOTIME, FILTIME) //2020-08-25 14:00:00 东八区时区  28800
	tsum := height*FILCHECK + times.Unix() - UTC_8
	return time.Unix(tsum, 0)
}

func Timeunix(height int64) int64 {
	times, _ := time.Parse(GOTIME, FILTIME) //2020-08-25 14:00:00 东八区时区  28800
	tsum := height*FILCHECK + times.Unix() - UTC_8
	return tsum
}

// 业务系统专用

func Timetimes() string {
	//2020-08-25 14:00:00 东八区时区  28800
	tsum := TimeUNix()
	timeStr := time.Unix(tsum, 0).Format(GOTIME1)
	return timeStr
}

func TimeMonth() string {
	//2020-08 东八区时区  28800
	tsum := TimeUNix()
	timeStr := time.Unix(tsum, 0).Format(MONTH1)
	return timeStr
}

func TimeDay() string {
	//2020-08-25 东八区时区  28800
	tsum := TimeUNix()
	timeStr := time.Unix(tsum, 0).Format(GOTIMEDAY1)
	return timeStr
}

func TimeDay1() string {
	//2020-08-25 东八区时区  28800
	tsum := TimeUNix()
	timeStr := time.Unix(tsum, 0).Format(GOTIMEDAY)
	return timeStr
}

func TimeHMS() string {
	//2020-08-25 东八区时区  28800
	tsum := TimeUNix()
	timeStr := time.Unix(tsum, 0).Format(GOTIME)
	return timeStr
}

func TimeHMSStr(tsum int64) string {
	//2020-08-25 东八区时区  28800
	timeStr := time.Unix(tsum, 0).Format(GOTIME)
	return timeStr
}

func TimeHMS1() time.Time {
	//2020-08-25 东八区时区  28800
	tsum := TimeUNix()
	return time.Unix(tsum, 0)
}

func TimeUNix() int64 {
	// 时间戳
	times := NewNtp()
	//return time.Now().Unix()
	return times.Unix()
}

func TimesStr(date string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation(GOTIME, date, loc)
	return tt.Unix()
}

func TImeInt64(date string) int64 {
	d, _ := time.Parse(date, GOTIMEDAY)
	return d.Unix()
}

func TimeDayInt64() int64 {
	//2020-08-25 东八区时区  28800
	tsum := TimeUNix()
	timeStr := time.Unix(tsum, 0).Format(GOTIMEDAY)
	d, _ := time.Parse(timeStr, GOTIMEDAY)
	return d.Unix()
}

func TimesStr1(date string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation(GOTIME, date, loc)
	return tt.Local()
}
