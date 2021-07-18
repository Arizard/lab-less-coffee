package spc

import "time"

func newTime(hour int, minute int, deltaDay int) time.Time {
	now := time.Date(2021, 7, 1, hour, minute, 0, 0, time.Local)
	now.Add(time.Hour * 24 * time.Duration(deltaDay))

	return now
}
