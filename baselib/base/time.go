package base

import (
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/timer"
)

// TryParseDate func; try parse date to time
// 2006-1
// 2006-01
// 2006-1-2
// 2006-01-02
func TryParseDate(date string) (time.Time, error) {
	return timer.TryParseDate(date)
}

const (
	TimestampFormat = timer.TimestampFormat
)

func GetTimestampData() string {
	return timer.GetTimestampData()
}
