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

func TestStreamingMonth_Year(t *testing.T) {
	t.Run("StreamingMonthのYear()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		y, err := NewYear(2021)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m, err := NewMonth(4)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm, err := NewStreamingMonth(y, m)
		if err != nil {
			t.Errorf("StreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		got := sm.Year()

		// 結果を検証
		if got != Year(2021) {
			t.Errorf("StreamingMonthのYear()メソッドが正しい値を返しませんでした。\n"+
				"got : %d\n"+
				"want: %d", got, y)
		}
	})
}

func TestStreamingMonth_Month(t *testing.T) {
	t.Run("StreamingMonthのMonth()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		y, err := NewYear(2021)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m, err := NewMonth(4)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm, err := NewStreamingMonth(y, m)
		if err != nil {
			t.Errorf("StreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		got := sm.Month()

		// 結果を検証
		if got != Month(4) {
			t.Errorf("StreamingMonthのMonth()メソッドが正しい値を返しませんでした。\n"+
				"got : %d\n"+
				"want: %d", got, m)
		}
	})
}

func TestStreamingMonth_String(t *testing.T) {
	t.Run("StreamingMonthのString()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		y, err := NewYear(2021)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m, err := NewMonth(4)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm, err := NewStreamingMonth(y, m)
		if err != nil {
			t.Errorf("StreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		got := sm.String()

		// 結果を検証
		want := "2021-04"
		if got != want {
			t.Errorf("StreamingMonthのString()メソッドが正しい値を返しませんでした。\n"+
				"got : %s\n"+
				"want: %s", got, want)
		}
	})
}

func TestStreamingMonth_IsEqualTo(t *testing.T) {
	t.Run("StreamingMonthのIsEqualTo()メソッドで同じ月を同じ月と判定できる", func(t *testing.T) {
		// テスト対象のものを用意
		y, err := NewYear(2021)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m, err := NewMonth(4)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm1, err := NewStreamingMonth(y, m)
		if err != nil {
			t.Errorf("ひとつ目のStreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm2, err := NewStreamingMonth(y, m)
		if err != nil {
			t.Errorf("ふたつ目のStreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		isEqual := sm1.IsEqualTo(sm2)

		// 結果を検証
		if !isEqual {
			t.Errorf("StreamingMonthのIsEqualTo()メソッドが同じ月を同じ月と判定しませんでした。")
		}
	})

	t.Run("StreamingMonthのIsEqualTo()メソッドで異なる月を異なる月と判定できる (年は同じパターン)", func(t *testing.T) {
		// テスト対象のものを用意
		y, err := NewYear(2021)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m1, err := NewMonth(4)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m2, err := NewMonth(5)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm1, err := NewStreamingMonth(y, m1)
		if err != nil {
			t.Errorf("ひとつ目のStreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm2, err := NewStreamingMonth(y, m2)
		if err != nil {
			t.Errorf("ふたつ目のStreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		isEqual := sm1.IsEqualTo(sm2)

		// 結果を検証
		if isEqual {
			t.Errorf("StreamingMonthのIsEqualTo()メソッドが異なる月を同じ月と判定してしまいました。")
		}
	})

	t.Run("StreamingMonthのIsEqualTo()メソッドで異なる月を異なる月と判定できる (月は同じパターン)", func(t *testing.T) {
		// テスト対象のものを用意
		y1, err := NewYear(2021)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		y2, err := NewYear(2022)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m, err := NewMonth(4)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm1, err := NewStreamingMonth(y1, m)
		if err != nil {
			t.Errorf("ひとつ目のStreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm2, err := NewStreamingMonth(y2, m)
		if err != nil {
			t.Errorf("ふたつ目のStreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		isEqual := sm1.IsEqualTo(sm2)

		// 結果を検証
		if isEqual {
			t.Errorf("StreamingMonthのIsEqualTo()メソッドが異なる月を同じ月と判定してしまいました。")
		}
	})

	t.Run("StreamingMonthのIsEqualTo()メソッドで異なる月を異なる月と判定できる (年月どちらも異なるパターン)", func(t *testing.T) {
		// テスト対象のものを用意
		y1, err := NewYear(2021)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		y2, err := NewYear(2022)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m1, err := NewMonth(4)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		m2, err := NewMonth(5)
		if err != nil {
			t.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm1, err := NewStreamingMonth(y1, m1)
		if err != nil {
			t.Errorf("ひとつ目のStreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		sm2, err := NewStreamingMonth(y2, m2)
		if err != nil {
			t.Errorf("ふたつ目のStreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		isEqual := sm1.IsEqualTo(sm2)

		// 結果を検証
		if isEqual {
			t.Errorf("StreamingMonthのIsEqualTo()メソッドが異なる月を同じ月と判定してしまいました。")
		}
	})
}
