package scheduler

import "time"

// BroadcastScheduler 全体通知の日時を管理するスケジューラ
type BroadcastScheduler struct {
	target time.Time
	sched  NotificationSchedule
}

func NewBroadcastScheduler(target time.Time, s NotificationSchedule) BroadcastScheduler {
	return BroadcastScheduler{
		target: target,
		sched:  s,
	}
}

func (s BroadcastScheduler) Target() time.Time {
	return s.target
}

func (s BroadcastScheduler) NotificationSchedule() NotificationSchedule {
	return s.sched
}

// IsTimeToNotify 受け取った時刻が通知を行うべきタイミングかどうかを判定する
func (s BroadcastScheduler) IsTimeToNotify(t time.Time) bool {
	tgt := s.Target()
	sched := s.NotificationSchedule()
	timeToNotify := time.Date(tgt.Year(), tgt.Month(), tgt.Day()-sched.DaysBefore,
		sched.Time.Hour.Value(), sched.Time.Minute.Value(), 0, 0, sched.Loc)

	// 通知タイミングよりも前の時刻の場合は通知しない
	if t.Before(timeToNotify) {
		return false
	}

	// AllowedDelayよりも後の時刻の場合は通知しない
	if t.After(timeToNotify.Add(sched.AllowedDelay)) {
		return false
	}

	return true
}
