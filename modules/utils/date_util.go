package utils

import "time"

func NowInt64() int64 {
	return time.Now().UnixMilli()
}
