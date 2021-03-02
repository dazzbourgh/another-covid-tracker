package util

import (
	"fmt"
	"time"
)

type MyTime time.Time
type MyMonth time.Month
type TimeString string

func (t MyTime) DateString() string {
	year, month, day := time.Time(t).Date()
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), day)
}

func (ts TimeString) Date() time.Time {
	date, _ := time.Parse("2006-January-02", string(ts))
	return date
}
