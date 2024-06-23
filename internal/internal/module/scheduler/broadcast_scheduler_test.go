package scheduler

import (
	"testing"
	"time"
)

func TestNewBroadcastScheduler(t *testing.T) {
	t.Run("指定したtargetとschedを持ったBroadcastSchedulerが生成される", func(t *testing.T) {
		// 入力を用意
		target := time.Date(2021, 4, 6, 23, 1, 2, 3, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}

		// テスト対象関数を実行
		scheduler := NewBroadcastScheduler(target, sched)

		// 出力を検証
		if scheduler.target != target {
			t.Errorf("BroadcastSchedulerのtargetが指定したtime.Timeと異なります。\n"+
				"got : %v\n"+
				"want: %v", scheduler.target, target)
		}
		if scheduler.sched != sched {
			t.Errorf("BroadcastSchedulerのschedが指定したNotificationScheduleと異なります。\n"+
				"got : %v\n"+
				"want: %v", scheduler.sched, sched)
		}
	})
}

func TestBroadcastScheduler_Target(t *testing.T) {
	t.Run("BroadcastSchedulerのTarget()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象を用意
		target := time.Date(2021, 4, 6, 23, 1, 2, 3, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}
		scheduler := NewBroadcastScheduler(target, sched)

		// テスト対象のメソッドを実行
		result := scheduler.Target()

		// 結果を検証
		if result != target {
			t.Errorf("BroadcastSchedulerのTarget()メソッドが正しい値を返しませんでした。\n"+
				"got : %v\n"+
				"want: %v", result, target)
		}
	})
}

func TestBroadcastScheduler_NotificationSchedule(t *testing.T) {
	t.Run("BroadcastSchedulerのNotificationSchedule()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象を用意
		target := time.Date(2021, 4, 6, 23, 1, 2, 3, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}
		scheduler := NewBroadcastScheduler(target, sched)

		// テスト対象のメソッドを実行
		result := scheduler.NotificationSchedule()

		// 結果を検証
		if result != sched {
			t.Errorf("BroadcastSchedulerのNotificationSchedule()メソッドが正しい値を返しませんでした。\n"+
				"got : %v\n"+
				"want: %v", result, sched)
		}
	})
}

func TestBroadcastScheduler_IsTimeToNotify(t *testing.T) {
	t.Run("通知時刻ぴったりのTimeが渡されたときにtrueを返す", func(t *testing.T) {
		// テスト対象を用意
		target := time.Date(2024, 6, 19, 23, 59, 59, 0, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}
		scheduler := NewBroadcastScheduler(target, sched)

		// 入力を用意
		input := time.Date(2024, 6, 18, 7, 30, 0, 0, time.UTC)

		// テスト対象のメソッドを実行
		isTheTime := scheduler.IsTimeToNotify(input)

		// 結果を検証
		if !isTheTime {
			t.Errorf("通知対象時刻ぴったりなのにfalseが返りました。")
		}
	})

	t.Run("通知時刻以降かつAllowedDelayぎりぎりのTimeが渡されたときにtrueを返す", func(t *testing.T) {
		// テスト対象を用意
		target := time.Date(2024, 6, 19, 23, 59, 59, 0, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}
		scheduler := NewBroadcastScheduler(target, sched)

		// 入力を用意
		input := time.Date(2024, 6, 18, 8, 0, 0, 0, time.UTC)

		// テスト対象のメソッドを実行
		isTheTime := scheduler.IsTimeToNotify(input)

		// 結果を検証
		if !isTheTime {
			t.Errorf("ぎりぎり通知対象時刻内なのにfalseが返りました。")
		}
	})

	t.Run("Timezoneが異なっていても通知対象時間内のTimeが渡されたときにtrueを返す", func(t *testing.T) {
		// テスト対象を用意
		target := time.Date(2024, 6, 19, 23, 59, 59, 0, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}
		scheduler := NewBroadcastScheduler(target, sched)

		// 入力を用意
		input := time.Date(2024, 6, 18, 16, 30, 0, 0, time.FixedZone("JST", 9*60*60))

		// テスト対象のメソッドを実行
		isTheTime := scheduler.IsTimeToNotify(input)

		// 結果を検証
		if !isTheTime {
			t.Errorf("Timezoneは異なっているが通知対象時間内なのにfalseが返りました。")
		}
	})

	t.Run("通知時刻より前のTimeが渡されたときにfalseを返す", func(t *testing.T) {
		// テスト対象を用意
		target := time.Date(2024, 6, 19, 23, 59, 59, 0, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}
		scheduler := NewBroadcastScheduler(target, sched)

		// 入力を用意
		input := time.Date(2024, 6, 18, 7, 29, 59, 0, time.UTC)

		// テスト対象のメソッドを実行
		isTheTime := scheduler.IsTimeToNotify(input)

		// 結果を検証
		if isTheTime {
			t.Errorf("通知対象時刻より前なのにtrueが返りました。")
		}
	})

	t.Run("通知時刻以降かつAllowedDelayより後のTimeが渡されたときにfalseを返す", func(t *testing.T) {
		// テスト対象を用意
		target := time.Date(2024, 6, 19, 23, 59, 59, 0, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}
		scheduler := NewBroadcastScheduler(target, sched)

		// 入力を用意
		input := time.Date(2024, 6, 18, 8, 0, 1, 0, time.UTC)

		// テスト対象のメソッドを実行
		isTheTime := scheduler.IsTimeToNotify(input)

		// 結果を検証
		if isTheTime {
			t.Errorf("通知対象時刻より後なのにtrueが返りました。")
		}
	})

	t.Run("時間は合っているが日付が違うTimeが渡されたときにfalseを返す", func(t *testing.T) {
		// テスト対象を用意
		target := time.Date(2024, 6, 19, 23, 59, 59, 0, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}
		scheduler := NewBroadcastScheduler(target, sched)

		// 入力を用意
		input := time.Date(2024, 6, 17, 7, 30, 0, 0, time.UTC)

		// テスト対象のメソッドを実行
		isTheTime := scheduler.IsTimeToNotify(input)

		// 結果を検証
		if isTheTime {
			t.Errorf("日付が違うのにtrueが返りました。")
		}
	})

	t.Run("Timezoneが異なっていて通知対象時間外のTimeが渡されたときにfalseを返す", func(t *testing.T) {
		// テスト対象を用意
		target := time.Date(2024, 6, 19, 23, 59, 59, 0, time.UTC)
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}
		scheduler := NewBroadcastScheduler(target, sched)

		// 入力を用意
		input := time.Date(2024, 6, 18, 16, 29, 59, 0, time.FixedZone("JST", 9*60*60))

		// テスト対象のメソッドを実行
		isTheTime := scheduler.IsTimeToNotify(input)

		// 結果を検証
		if isTheTime {
			t.Errorf("Timezoneは異なっているが通知対象時間外なのにtrueが返りました。")
		}
	})
}
