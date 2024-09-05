package dates

import (
	"strconv"
	"time"
)

const (
	LOC                  = "Asia/Shanghai"
	FmtDate              = "2006-01-02"
	FmtTime              = "15:04:05"
	FmtDateTime          = "2006-01-02 15:04:05"
	FmtDateTimeNoSeconds = "2006-01-02 15:04"
)

// NowUnix 秒时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// FromUnix 秒时间戳转时间
func FromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// NowTimestamp 当前毫秒时间戳
func NowTimestamp() int64 {
	return Timestamp(time.Now())
}

// Timestamp 毫秒时间戳
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// FromTimestamp 毫秒时间戳转时间
func FromTimestamp(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond))
}

// Format 时间格式化
func Format(time time.Time, layout string) string {
	return time.Format(layout)
}

// Parse 字符串时间转时间类型
func Parse(timeStr, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}

func Time2Int64(timeStr string) int64 {
	//Loc, _ := time.LoadLocation(LOC)
	tims, _ := time.ParseInLocation(FmtDateTime, timeStr, time.Local)
	return tims.Unix()
}

func Int642Time(timeInt int64) string {
	tims := FromUnix(timeInt)
	return Format(tims, FmtDateTime)
}

// GetDay return yyyyMMdd
func GetDay(time time.Time) int {
	ret, _ := strconv.Atoi(time.Format("20060102"))
	return ret
}

// WithTimeAsStartOfDay
// 返回指定时间当天的开始时间
func WithTimeAsStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
