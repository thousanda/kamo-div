package scheduler

import "testing"

func TestNewHour(t *testing.T) {
	t.Run("0時のHourが生成できる", func(t *testing.T) {
		// 入力を用意
		input := 0

		// テスト対象関数を実行
		h, err := NewHour(input)

		// 出力を検証
		if err != nil {
			t.Errorf("正常にHourが作成されるべきなのにerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		if int(h) != input {
			t.Errorf("指定した時間のHourが作成できませんでした。\n"+
				"got : %d\n"+
				"want: %d", h, input)
		}
	})

	t.Run("23時のHourが生成できる", func(t *testing.T) {
		// 入力を用意
		input := 23

		// テスト対象関数を実行
		h, err := NewHour(input)

		// 出力を検証
		if err != nil {
			t.Errorf("正常にHourが作成されるべきなのにerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		if int(h) != input {
			t.Errorf("指定した時間のHourが作成できませんでした。\n"+
				"got : %d\n"+
				"want: %d", h, input)
		}
	})

	t.Run("-1時は存在しない", func(t *testing.T) {
		// 入力を用意
		input := -1

		// テスト対象関数を実行
		_, err := NewHour(input)

		// 出力を検証
		if err == nil {
			t.Errorf("-1時を作ろうとしたのにerrが返ってきませんでした。")
		}
	})

	t.Run("24時は存在しない", func(t *testing.T) {
		// 入力を用意
		input := 24

		// テスト対象関数を実行
		_, err := NewHour(input)

		// 出力を検証
		if err == nil {
			t.Errorf("24時を作ろうとしたのにerrが返ってきませんでした。")
		}
	})
}

func TestHour_Value(t *testing.T) {
	t.Run("HourのValue()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象を用意
		h := Hour(12)

		// テスト対象のメソッドを実行
		result := h.Value()

		// 結果を検証
		if result != 12 {
			t.Errorf("HourのValueメソッドが正しい値を返しませんでした。\n"+
				"got : %d\n"+
				"want: %d", result, 12)
		}
	})
}

func TestNewMinute(t *testing.T) {
	t.Run("0分のMinuteが生成される", func(t *testing.T) {
		// 入力を用意
		input := 0

		// テスト対象関数を実行
		m, err := NewMinute(input)

		// 出力を検証
		if err != nil {
			t.Errorf("正常にMinuteが作成されるべきなのにerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		if int(m) != input {
			t.Errorf("指定した分のMinuteが作成できませんでした。\n"+
				"got : %d\n"+
				"want: %d", m, input)
		}
	})

	t.Run("59分のMinuteが生成される", func(t *testing.T) {
		// 入力を用意
		input := 59

		// テスト対象関数を実行
		m, err := NewMinute(input)

		// 出力を検証
		if err != nil {
			t.Errorf("正常にMinuteが作成されるべきなのにerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		if int(m) != input {
			t.Errorf("指定した分のMinuteが作成できませんでした。\n"+
				"got : %d\n"+
				"want: %d", m, input)
		}
	})

	t.Run("-1分は存在しない", func(t *testing.T) {
		// 入力を用意
		input := -1

		// テスト対象関数を実行
		_, err := NewMinute(input)

		// 出力を検証
		if err == nil {
			t.Errorf("-1分を作ろうとしたのにerrが返ってきませんでした。")
		}
	})

	t.Run("60分は存在しない", func(t *testing.T) {
		// 入力を用意
		input := 60

		// テスト対象関数を実行
		_, err := NewMinute(input)

		// 出力を検証
		if err == nil {
			t.Errorf("60分を作ろうとしたのにerrが返ってきませんでした。")
		}
	})
}

func TestMinute_Value(t *testing.T) {
	t.Run("MinuteのValue()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象を用意
		m := Minute(30)

		// テスト対象のメソッドを実行
		result := m.Value()

		// 結果を検証
		if result != 30 {
			t.Errorf("MinuteのValueメソッドが正しい値を返しませんでした。\n"+
				"got : %d\n"+
				"want: %d", result, 30)
		}
	})
}
