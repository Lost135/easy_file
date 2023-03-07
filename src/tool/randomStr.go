package tool

import (
	"math/rand"
	"time"
)

const (
	All     = 0
	EnLower = 1
	EnUpper = 2
	Num     = 4
	Pun     = 8
)

var SignDetail = map[int]string{
	EnLower: "abcdefghijklmnopqrstuvwxyz",
	EnUpper: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	Num:     "1234567890",
	//字符串隐患: "'\*?
	Pun: "`~!@#$%^&*()-_=+[|\\{+]};:,<.>/?\"",
}

// TODO 逻辑优化，字符串类型优化
func RandomStr(length int, sign int) string {
	StrRange := ""
	if sign > 15 || sign < All {
		return StrRange
	}
	if sign == All {
		for _, v := range SignDetail {
			StrRange += v
		}
	} else {
		for {
			if sign == All {
				break
			}
			if sign >= Pun {
				sign -= Pun
				StrRange += SignDetail[Pun]
			}
			if sign >= Num {
				sign -= Num
				StrRange += SignDetail[Num]
			}
			if sign >= EnUpper {
				sign -= EnUpper
				StrRange += SignDetail[EnUpper]
			}
			if sign == EnLower {
				sign -= EnLower
				StrRange += SignDetail[EnLower]
				break
			}
		}
	}
	bytes := []byte(StrRange)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
