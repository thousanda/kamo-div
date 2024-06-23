package deadline

import "testing"

func TestNewYear(t *testing.T) {
	t.Run("指定した年のYearが作成できる", func(t *testing.T) {
		// 入力を用意
		input := 2021

		// テスト対象の関数を実行
		y, err := NewYear(input)

		// 結果を検証
		if err != nil {
			t.Errorf("正常にYearが作成されるべきなのにerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}
		if int(y) != input {
			t.Errorf("指定した年のYearが作成できませんでした。\n"+
				"got : %d\n"+
				"want: %d", y, input)
		}
	})

	t.Run("番組開始が2021年なので、それより前の年を持ったYearは作成できない", func(t *testing.T) {
		// 入力を用意
		input := 2020

		// テスト対象の関数を実行
		_, e := NewYear(input)

		// 結果を検証
		if e == nil {
			t.Errorf("番組開始より前の年を持つYearを作成しようとしたのにerrが返りませんでした。")
		}
	})
}

func TestYear_Value(t *testing.T) {
	t.Run("YearのValueメソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		y := Year(2021)

		// テスト対象のメソッドを実行
		v := y.Value()

		// 結果を検証
		if v != 2021 {
			t.Errorf("YearのValue()メソッドが正しい値を返しませんでした。\n"+
				"got : %v\n"+
				"want: 2021", y.Value())
		}
	})
}

func TestNewMonth(t *testing.T) {
	t.Run("指定した月のMonthが作成できる (1月)", func(t *testing.T) {
		// 入力を用意
		input := 1

		// テスト対象の関数を実行
		m, err := NewMonth(input)

		// 結果を検証
		if err != nil {
			t.Errorf("正常にMonthが作成されるべきなのにerrが返ってきてしまいました。\n"+
				"error: %v", err)
		}
		if int(m) != input {
			t.Errorf("指定した月のMonthが作成できませんでした。\n"+
				"got : %v\n"+
				"want: %v", m, input)
		}
	})

	t.Run("指定した月のMonthが作成できる (12月)", func(t *testing.T) {
		// 入力を用意
		input := 12

		// テスト対象の関数を実行
		m, err := NewMonth(input)

		// 結果を検証
		if err != nil {
			t.Errorf("正常にMonthが作成されるべきなのにerrが返ってきてしまいました。\n"+
				"error: %v", err)
		}
		if int(m) != input {
			t.Errorf("指定した月のMonthが作成できませんでした。\n"+
				"got : %v\n"+
				"want: %v", m, input)
		}
	})

	t.Run("0月は存在しない", func(t *testing.T) {
		// 入力を用意
		input := 0

		// テスト対象の関数を実行
		_, err := NewMonth(input)

		// 結果を検証
		if err == nil {
			t.Errorf("0月を作ろうとしたのにerrが返ってきませんでした。\n"+
				"error: %v", err)
		}
	})

	t.Run("-1月は存在しない", func(t *testing.T) {
		// 入力を用意
		input := -1

		// テスト対象の関数を実行
		_, err := NewMonth(input)

		// 結果を検証
		if err == nil {
			t.Errorf("-1月を作ろうとしたのにerrが返ってきませんでした。\n"+
				"error: %v", err)
		}
	})

	t.Run("13月は存在しない", func(t *testing.T) {
		// 入力を用意
		input := 13

		// テスト対象の関数を実行
		_, err := NewMonth(input)

		// 結果を検証
		if err == nil {
			t.Errorf("13月を作ろうとしたのにerrが返ってきませんでした。\n"+
				"error: %v", err)
		}
	})
}

func TestMonth_Value(t *testing.T) {
	t.Run("MonthのValueメソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		m := Month(12)

		if m.Value() != 12 {
			t.Errorf("MonthのValue()メソッドが正しい値を返しませんでした。\n"+
				"got : %v\n"+
				"want: 12", m.Value())
		}
	})
}
