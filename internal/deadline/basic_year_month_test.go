package deadline

import "testing"

func TestNewYear(t *testing.T) {
	t.Run("指定した年のYearが作成できる", func(t *testing.T) {
		// 入力を用意
		input := 2021

		// テスト対象を実行
		y, err := NewYear(input)

		// 結果を検証
		if err != nil {
			t.Errorf("意図せずInvalidなYearが作成されてしまいました。\n"+
				"error: %v", err)
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

		// テスト対象を実行
		_, e := NewYear(input)

		// 結果を検証
		if e == nil {
			t.Errorf("番組開始より前の年を持つYearを作成しようとしたのにerrが返りませんでした。")
		}
	})
}

func TestYear_Value(t *testing.T) {
	t.Run("YearのValueメソッドが正しい値を返す", func(t *testing.T) {
		// 入力を用意
		input := 2021

		// テスト対象を実行
		y := Year(input)

		// 結果を検証
		if y.Value() != input {
			t.Errorf("YearのValue()メソッドが正しい値を返しませんでした。\n"+
				"got : %v\n"+
				"want: %v", y.Value(), input)
		}
	})
}

func TestNewMonth(t *testing.T) {
	t.Run("指定した月のMonthが作成できる (1月)", func(t *testing.T) {
		input := 1

		m := NewMonth(input)

		if !m.isValid {
			t.Errorf("意図せずInvalidなMonthが作成されてしまいました。\n"+
				"got : %v\n"+
				"want: true", m.isValid)
		}

		if m.month != input {
			t.Errorf("指定した月のMonthが作成できませんでした。\n"+
				"got : %v\n"+
				"want: %v", m, input)
		}
	})

	t.Run("指定した月のMonthが作成できる (12月)", func(t *testing.T) {
		input := 12

		m := NewMonth(input)

		if !m.isValid {
			t.Errorf("意図せずInvalidなMonthが作成されてしまいました。\n"+
				"got : %v\n"+
				"want: true", m.isValid)
		}

		if m.month != input {
			t.Errorf("指定した月のMonthが作成できませんでした。\n"+
				"got : %v\n"+
				"want: %v", m, input)
		}
	})

	t.Run("0月は存在しない", func(t *testing.T) {
		input := 0

		m := NewMonth(input)

		if m.isValid {
			t.Errorf("月の値として0を入力したのにValidなMonthが作成されてしまいました。\n"+
				"got : %v\n"+
				"want: true", m.isValid)
		}
	})

	t.Run("負の月を持つMonthは存在しない", func(t *testing.T) {
		input := 0

		m := NewMonth(input)

		if m.isValid {
			t.Errorf("月の値として負の値を入力したのにValidなMonthが作成されてしまいました。\n"+
				"got : %v\n"+
				"want: true", m.isValid)
		}
	})

	t.Run("13月以上は存在しない", func(t *testing.T) {
		input := 13

		m := NewMonth(input)

		if m.isValid {
			t.Errorf("月の値として13を入力したのにValidなMonthが作成されてしまいました。\n"+
				"got : %v\n"+
				"want: true", m.isValid)
		}
	})
}

func TestMonth_Value(t *testing.T) {
	t.Run("MonthのValueメソッドが正しい値を返す", func(t *testing.T) {
		input := 4

		m := NewMonth(input)

		if m.Value() != input {
			t.Errorf("MonthのValue()メソッドが正しい値を返しませんでした。\n"+
				"got : %v\n"+
				"want: %v", m.Value(), input)
		}
	})
}

func TestMonth_IsValid(t *testing.T) {
	t.Run("正常なMonthのIsValidメソッド()がtrueを返す", func(t *testing.T) {
		input := 12

		m := NewMonth(input)

		if !m.IsValid() {
			t.Errorf("12月のMonthは正常なのにIsValid()メソッドがfalseを返しました。\n"+
				"got : %v\n"+
				"want: %v", m.IsValid(), true)
		}
	})

	t.Run("異常なMonthのIsValidメソッド()がfalseを返す", func(t *testing.T) {
		input := 0

		m := NewMonth(input)

		if m.IsValid() {
			t.Errorf("0月のMonthは異常なのにIsValid()メソッドがtrueを返しました。\n"+
				"got : %v\n"+
				"want: %v", m.IsValid(), false)
		}
	})
}

func TestMonth_String(t *testing.T) {
	t.Run("正常なMonthのString()メソッドが正しい値を返す", func(t *testing.T) {
		input := 12

		m := NewMonth(input)

		if m.String() != "12" {
			t.Errorf("正常な月である12月のMonthのString()メソッドが '12' を返しませんでした。\n"+
				"got : %v\n"+
				"want: %v", m.String(), "12")
		}
	})

	t.Run("異常なMonthのString()メソッドが正しい値を返す", func(t *testing.T) {
		input := 0

		m := NewMonth(input)

		if m.String() != "Invalid Month" {
			t.Errorf("異常な月である0月のMonthのString()メソッドが 'Invalid Month' 以外の文字列を返しました。\n"+
				"got : %v\n"+
				"want: %v", m.String(), "Invalid Month")
		}
	})
}
