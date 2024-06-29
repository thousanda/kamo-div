package deadline

import "time"

// Deadline 締め切り日時を表すもの
type Deadline time.Time

func NewDeadline(t time.Time) Deadline {
	return Deadline(t)
}

func (d Deadline) ToTime() time.Time {
	return time.Time(d)
}

func (d Deadline) Equal(other Deadline) bool {
	return d == other
}
