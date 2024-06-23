package scheduler

import "fmt"

// Hour 時間を表すもの
// 0時から23時までを表現できる
type Hour int

func NewHour(h int) (Hour, error) {
	if h < 0 || 23 < h {
		return Hour(0), fmt.Errorf("hour must be between 0 and 23; got %d", h)
	}
	return Hour(h), nil
}

func (h Hour) Value() int {
	return int(h)
}

// Minute 分を表すもの
// 0分から59分までを表現できる
type Minute int

func NewMinute(m int) (Minute, error) {
	if m < 0 || 59 < m {
		return Minute(0), fmt.Errorf("minute must be between 0 and 59; got %d", m)
	}
	return Minute(m), nil
}

func (m Minute) Value() int {
	return int(m)
}
