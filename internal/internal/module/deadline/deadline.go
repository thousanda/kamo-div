package deadline

import (
	"fmt"
	"time"
)

// Deadline 締め切り日時を表すもの
type Deadline time.Time

func NewDeadline(t time.Time) Deadline {
	return Deadline(t)
}

func (d Deadline) ToTime() time.Time {
	return time.Time(d)
}

// ToJapaneseMDWdHMString 日本語表記の日時を返す
func (d Deadline) ToJapaneseMDWdHMString() string {
	t := d.ToTime()
	return fmt.Sprintf("%d月%d日 (%s) %d:%02d",
		t.Month(), t.Day(), d.weekdayInShortJapanese(),
		t.Hour(), t.Minute())
}

func (d Deadline) Equal(other Deadline) bool {
	return d == other
}

func (d Deadline) weekdayInShortJapanese() string {
	switch d.ToTime().Weekday() {
	case time.Sunday:
		return "日"
	case time.Monday:
		return "月"
	case time.Tuesday:
		return "火"
	case time.Wednesday:
		return "水"
	case time.Thursday:
		return "木"
	case time.Friday:
		return "金"
	default:
		return "土"
	}
}
