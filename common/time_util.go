package common

import "time"

func NsToMs(ns int64) int64 {
	return int64(time.Nanosecond) * ns / int64(time.Millisecond)
}
