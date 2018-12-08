package impl

import "time"

type byTimestamp []record

type record struct {
	timestamp time.Time
	guardID   int
	eventID   int
	// 0: begins shift
	// 1: falls asleep
	// 2: wakes up
}

func (r byTimestamp) Len() int {
	return len(r)
}
func (r byTimestamp) Swap(i int, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r byTimestamp) Less(i, j int) bool {
	return r[i].timestamp.Before(r[j].timestamp)
}
