package scheduler

import "time"

// NotificationTime 1日のうち、どのタイミングで通知すればいいかを示すもの
// 時と分の粒度で通知タイミングを指定する
type NotificationTime struct {
	Hour
	Minute
}

// NotificationSchedule 通知のスケジュールを示すもの
type NotificationSchedule struct {
	// DaysBefore 何日前に通知するか
	DaysBefore int
	// Time その日の何時何分に通知するか
	Time NotificationTime
	Loc  *time.Location
	// AllowedDelay 指定時刻からどれだけの遅れまで許容するか
	AllowedDelay time.Duration
}

func (s NotificationSchedule) Equal(other NotificationSchedule) bool {
	return s.DaysBefore == other.DaysBefore &&
		s.Time == other.Time &&
		s.Loc.String() == other.Loc.String() &&
		s.AllowedDelay == other.AllowedDelay
}
