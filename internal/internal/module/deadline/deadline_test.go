package deadline

import (
	"testing"
	"time"
)

func TestNewDeadline(t *testing.T) {
	t.Run("指定した時刻のDeadlineが作成できる", func(t *testing.T) {
		// 入力を用意
		input := time.Date(2021, 4, 6, 1, 2, 3, 0, time.UTC)

		// テスト対象の関数を実行
		d := NewDeadline(input)

		// 結果を検証
		if d.ToTime() != input {
			t.Errorf("指定した時刻のDeadlineが作成できませんでした。\n"+
				"got : %v\n"+
				"want: %v", d.ToTime(), input)
		}
	})
}

func TestDeadline_ToTime(t *testing.T) {
	t.Run("DeadlineのToTime()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		d := NewDeadline(time.Date(2021, 4, 6, 1, 2, 3, 0, time.UTC))

		// テスト対象のメソッドを実行
		v := d.ToTime()

		// 結果を検証
		want := time.Date(2021, 4, 6, 1, 2, 3, 0, time.UTC)
		if v != want {
			t.Errorf("DeadlineのToTime()メソッドが正しい値を返しませんでした。\n"+
				"got : %v\n"+
				"want: %v", d.ToTime(), want)
		}
	})
}

func TestDeadline_Equal(t *testing.T) {
	t.Run("DeadlineがEqual()メソッドで同値かどうかを判定できる", func(t *testing.T) {
		// テスト対象のものを用意
		d1 := NewDeadline(time.Date(2021, 4, 6, 1, 2, 3, 0, time.UTC))
		d2 := NewDeadline(time.Date(2021, 4, 6, 1, 2, 3, 0, time.UTC))

		// テスト対象のメソッドを実行
		v := d1.Equal(d2)

		// 結果を検証
		want := true
		if v != want {
			t.Errorf("DeadlineがEqual()メソッドで同値かどうかを判定できませんでした。\n"+
				"got : %v\n"+
				"want: %v", v, want)
		}
	})
}
