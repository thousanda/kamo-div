package scheduler

import (
	"testing"
	"time"
)

func TestNotificationSchedule_Equal(t *testing.T) {
	t.Run("同じNotificationScheduleが同じと判定できる", func(t *testing.T) {
		// テスト対象のものを用意
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}

		// 入力を用意
		other := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}

		// テスト対象関数を実行
		isEqual := sched.Equal(other)

		// 出力を検証
		if !isEqual {
			t.Errorf("同じNotificationScheduleが異なると判定されてしまいました。")
		}
	})

	t.Run("DaysBeforeが異なるNotificationScheduleが異なると判定できる", func(t *testing.T) {
		// テスト対象のものを用意
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}

		// 入力を用意
		other := NotificationSchedule{
			DaysBefore:   2,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}

		// テスト対象関数を実行
		isEqual := sched.Equal(other)

		// 出力を検証
		if isEqual {
			t.Errorf("DaysBeforeが異なるNotificationScheduleが同じと判定されてしまいました。")
		}
	})

	t.Run("Timeが異なるNotificationScheduleが異なると判定できる", func(t *testing.T) {
		// テスト対象のものを用意
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}

		// 入力を用意
		other := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 8, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}

		// テスト対象関数を実行
		isEqual := sched.Equal(other)

		// 出力を検証
		if isEqual {
			t.Errorf("Timeが異なるNotificationScheduleが同じと判定されてしまいました。")
		}
	})

	t.Run("Locが異なるNotificationScheduleが異なると判定できる", func(t *testing.T) {
		// テスト対象のものを用意
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}

		// 入力を用意
		other := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.FixedZone("JST", 9*60*60),
			AllowedDelay: 30 * time.Minute,
		}

		// テスト対象関数を実行
		isEqual := sched.Equal(other)

		// 出力を検証
		if isEqual {
			t.Errorf("Locが異なるNotificationScheduleが同じと判定されてしまいました。")
		}
	})

	t.Run("AllowedDelayが異なるNotificationScheduleが異なると判定できる", func(t *testing.T) {
		// テスト対象のものを用意
		sched := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 30 * time.Minute,
		}

		// 入力を用意
		other := NotificationSchedule{
			DaysBefore:   1,
			Time:         NotificationTime{Hour: 7, Minute: 30},
			Loc:          time.UTC,
			AllowedDelay: 40 * time.Minute,
		}

		// テスト対象関数を実行
		isEqual := sched.Equal(other)

		// 出力を検証
		if isEqual {
			t.Errorf("AllowedDelayが異なるNotificationScheduleが同じと判定されてしまいました。")
		}
	})
}
