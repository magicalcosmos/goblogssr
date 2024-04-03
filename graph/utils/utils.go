package utils

import "time"

// 获取当前系统时间
func GetCurrentTime() time.Time {
	currentTime := time.Now()
	date, _ := time.Parse("2006-01-02 15:04:05", currentTime.Format("2006-01-02 15:04:05"))
	return date
}
