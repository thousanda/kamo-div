package reminder

import (
	"testing"
	"time"

	"github.com/thousanda/kamo-div/internal/internal/module/deadline"
)

func TestNewMessageWriter(t *testing.T) {
	t.Run("MessageWriterを作成できる", func(t *testing.T) {
		// テスト対象のメソッドを実行
		w := NewMessageWriter()

		// 結果を検証
		if w == nil {
			t.Errorf("MessageWriterがnilです")
		}
	})
}

func TestMessageWriter_WriteReminderOneWeekAhead(t *testing.T) {
	t.Run("1週間前のリマインダーメッセージを書ける", func(t *testing.T) {
		// テスト対象のものを用意
		w := NewMessageWriter()

		// テスト対象のメソッドを実行
		r := w.WriteReminderOneWeekAhead(deadline.NewDeadline(
			time.Date(2020, 12, 31, 23, 59, 0, 0, time.Local),
		))

		// 結果を検証
		want := Reminder{"1週間後がスポットスポンサーの締切だよ！🦆",
			"締め切り: 12月31日 (木) 23:59",
			"https://higuchi.world/gichiland-spot-sponsor",
		}
		if len(r) != len(want) {
			t.Errorf("lengthが違います。\n"+
				"got : %v\n"+
				"want: %v", len(r), len(want))
		}
		for i := range r {
			if r[i] != want[i] {
				t.Errorf("%d番目の要素が違います。\n"+
					"got : %v\n"+
					"want: %v", i, r[i], want[i])
			}
		}
	})
}

func TestMessageWriter_WriteReminderOneDayAhead(t *testing.T) {
	t.Run("1日前のリマインダーメッセージを書ける", func(t *testing.T) {
		// テスト対象のものを用意
		w := NewMessageWriter()

		// テスト対象のメソッドを実行
		r := w.WriteReminderOneDayAhead(deadline.NewDeadline(
			time.Date(2020, 12, 31, 23, 59, 0, 0, time.Local),
		))

		// 結果を検証
		want := Reminder{"明日がスポットスポンサーの締切だよ！🦆",
			"締め切り: 12月31日 (木) 23:59",
			"https://higuchi.world/gichiland-spot-sponsor",
		}
		if len(r) != len(want) {
			t.Errorf("lengthが違います。\n"+
				"got : %v\n"+
				"want: %v", len(r), len(want))
		}
		for i := range r {
			if r[i] != want[i] {
				t.Errorf("%d番目の要素が違います。\n"+
					"got : %v\n"+
					"want: %v", i, r[i], want[i])
			}
		}
	})
}

func TestMessageWriter_WriteReminderOnTheDay(t *testing.T) {
	t.Run("当日のリマインダーメッセージを書ける", func(t *testing.T) {
		// テスト対象のものを用意
		w := NewMessageWriter()

		// テスト対象のメソッドを実行
		r := w.WriteReminderOnTheDay(deadline.NewDeadline(
			time.Date(2020, 12, 31, 23, 59, 0, 0, time.Local),
		))

		// 結果を検証
		want := Reminder{"今日がスポットスポンサーの締切だよ！🦆",
			"締め切り: 12月31日 (木) 23:59",
			"https://higuchi.world/gichiland-spot-sponsor",
		}
		if len(r) != len(want) {
			t.Errorf("lengthが違います。\n"+
				"got : %v\n"+
				"want: %v", len(r), len(want))
		}
		for i := range r {
			if r[i] != want[i] {
				t.Errorf("%d番目の要素が違います。\n"+
					"got : %v\n"+
					"want: %v", i, r[i], want[i])
			}
		}
	})
}
