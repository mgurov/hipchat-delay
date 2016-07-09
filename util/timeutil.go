package util

import (
	"time"
)

func MergeDateTime(datePart, timePart time.Time) time.Time {

	year, month, day := datePart.Date()
	hour, min, sec := timePart.Clock()

	return time.Date(year, month, day, hour, min, sec, timePart.Nanosecond(), datePart.Location())
}
