package tool

import (
	"time"
)

func CTimeYMDHmS() string {
	currentTime := time.Now().Local().Format("2006-01-02 15:04:05 Z07")
	return currentTime
}
