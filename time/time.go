package utilsTime

import (
	"time"
)

const (
	FormatDate       = "2006-01-02"
	FormatTime       = "15:04:05"
	FormatMinute     = "2006-01-02 15:04"
	FormatSecond     = "2006-01-02 15:04:05"
	FormatMilli      = "2006-01-02 15:04:05.000"
	SecondTimeStamp  = "20060102150405"
	SecondMilliStamp = "20060102150405000"
)

// NowDate 获取当前日期
func NowDate() string {
	return time.Now().Format(FormatDate)
}

// NowTime 获取当前时间 时分秒
func NowTime() string {
	return time.Now().Format(FormatTime)
}

// NowMinute 日期到分钟
func NowMinute() string {
	return time.Now().Format(FormatMinute)
}

// NowSecond 日期时分秒
func NowSecond() string {
	return time.Now().Format(FormatSecond)
}

// NowMilli 日期时分秒毫秒
func NowMilli() string {
	return time.Now().Format(FormatMilli)
}

// NowSecondStamp 格式化秒级时间戳
func NowSecondStamp() string {
	return time.Now().Format(SecondTimeStamp)
}

// NowMilliStamp 格式化毫秒级时间戳
func NowMilliStamp() string {
	return time.Now().Format(SecondMilliStamp)
}
