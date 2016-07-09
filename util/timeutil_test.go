package util

import (
	"testing"
	"time"
)

func TestDateTimeMerger(t *testing.T) {
	actual := MergeDateTime(
		time.Date(1000, 11, 12, 13, 14, 15, 16, time.Local),
		time.Date(2000, 3, 4, 5, 6, 7, 8, time.Local),
	)

	expected := time.Date(1000, 11, 12, 5, 6, 7, 8, time.Local)
	if actual != expected {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}
