package utils

import (
	"time"

	"gitlab.com/fastogt/gofastogt/gofastogt"
)

func GetMinSecTimestamp(timestamp gofastogt.DurationMsec) (int, int) {
	tm := time.UnixMilli(int64(timestamp))

	return tm.Minute(), tm.Second()
}
