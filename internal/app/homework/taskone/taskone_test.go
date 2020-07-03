package taskone

import (
	"testing"
	"time"
)

func TestCorrectTime(t *testing.T) {
	tm := CorrectTime()
	now := time.Now()

	if tm != now {
		t.Logf("Local Time %v\n", now)
		t.Logf("~True Time %v\n", tm)
		t.Logf("Offset %v\n", tm.Sub(now))
	}
}
