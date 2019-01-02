package format

import (
	"fmt"
	"time"
)

type RelTime struct {
	Time time.Time
}

func (c RelTime) String() string {
	diff := time.Now().Sub(c.Time)
	sec := diff.Seconds()
	suffix := "ago"
	if sec < 0 {
		suffix = "from now"
		sec = -sec
	}
	sec += 0.5
	switch {
	case sec < 1:
		return "just now"
	case sec < 60:
		return fmt.Sprintf("%s %s", AutoPluralize(int(sec), "second"), suffix)
	case sec < 60*60:
		return fmt.Sprintf("%s %s", AutoPluralize(int(sec/60), "minute"), suffix)
	case sec < 60*60*24:
		return fmt.Sprintf("%s %s", AutoPluralize(int(sec/60/60), "hour"), suffix)
	default:
		return fmt.Sprintf("%s %s", AutoPluralize(int(sec/60/60/24), "day"), suffix)
	}
}
