package deadline

import "testing"

func TestNewStreamingMonth(t *testing.T) {
	t.Run("指定した年月のStreamingMonthが作成できる", func(t *testing.T) {
		// 入力を用意
		y, err := NewYear(2021)
		if err != nil {
			t.Errorf("入力のYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m, err := NewMonth(4)
		if err != nil {
			t.Errorf("入力のMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象の関数を実行
		sm, err := NewStreamingMonth(y, m)

		// 結果を検証
		if err != nil {
			t.Errorf("正常にStreamingMonthが作成されるべきなのにerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		if sm.year != y {
			t.Errorf("StreamingMonthのyearが指定したYearと異なります。\n"+
				"got : %d\n"+
				"want: %d", sm.year, y)
		}
		if sm.month != m {
			t.Errorf("StreamingMonthのmonthが指定したMonthと異なります。\n"+
				"got : %d\n"+
				"want: %d", sm.month, m)
		}
	})

	t.Run("2021年4月より前の年月を持つStreamingMonthは作成できない", func(t *testing.T) {
		// YearやMonthはNewYearやNewMonthで生成される前提なので、これらが正常な範囲の値であることはを信じることにする

		// 入力を用意
		y, err := NewYear(2021)
		if err != nil {
			t.Errorf("入力のYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m, err := NewMonth(3)
		if err != nil {
			t.Errorf("入力のMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象の関数を実行
		_, e := NewStreamingMonth(y, m)

		// 結果を検証
		if e == nil {
			t.Errorf("2021年4月より前の年月を持つStreamingMonthを作成しようとしたのにerrが返りませんでした。")
		}
	})
}
