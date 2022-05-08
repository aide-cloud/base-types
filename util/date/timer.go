package date

import "time"

// MicTimeToFormatStr 毫秒时间戳转字符串
func MicTimeToFormatStr() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}
