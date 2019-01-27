package utils

import "time"

func NowTimef() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NowN() int {
	return time.Now().Nanosecond()
}
