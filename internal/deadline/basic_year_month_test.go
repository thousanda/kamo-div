package deadline

import "testing"

func TestNewYear(t *testing.T) {
	t.Run("指定した年のYearが作成できる", func(t *testing.T) {
		input := 2021

		y := NewYear(input)

		if !y.isValid {
			t.Errorf("意図せずInvalidなYearが作成されてしまいました。\n"+
				"got : %v\n"+
				"want: true", y.isValid)
		}

		if y.year != input {
			t.Errorf("指定した年のYearが作成できませんでした。\n"+
				"got : %v\n"+
				"want: %v", y, input)
		}
	})

	t.Run("番組開始が2021年なので、それより前の年を持ったYearは作成できない", func(t *testing.T) {
		input := 2020

		y := NewYear(input)

		if y.isValid {
			t.Errorf("番組開始より前の年を持つYearが作成されてしまいました。\n"+
				"got : %v\n"+
				"want: false", y.isValid)
		}
	})
}

func TestYear_Value(t *testing.T) {
	t.Run("YearのValueメソッドが正しい値を返す", func(t *testing.T) {
		input := 2021

		y := NewYear(input)

		if y.Value() != input {
			t.Errorf("YearのValue()メソッドが正しい値を返しませんでした。\n"+
				"got : %v\n"+
				"want: %v", y.Value(), input)
		}
	})
}

func TestYear_IsValid(t *testing.T) {
	t.Run("正常なYearのIsValidメソッド()がtrueを返す", func(t *testing.T) {
		input := 2021

		y := NewYear(input)

		if !y.IsValid() {
			t.Errorf("2021年のYearは正常なのにIsValid()メソッドがfalseを返しました。\n"+
				"got : %v\n"+
				"want: %v", y.IsValid(), true)
		}
	})

	t.Run("異常なYearのIsValidメソッド()がfalseを返す", func(t *testing.T) {
		input := 2020

		y := NewYear(input)

		if y.IsValid() {
			t.Errorf("2020年のYearは異常なのにIsValid()メソッドがtrueを返しました。\n"+
				"got : %v\n"+
				"want: %v", y.IsValid(), false)
		}
	})
}

func TestYear_String(t *testing.T) {
	t.Run("正常なYearのString()メソッドが正しい値を返す", func(t *testing.T) {
		input := 2021

		y := NewYear(input)

		if y.String() != "2021" {
			t.Errorf("正常な年である2021年のYearのString()メソッドが '2021' を返しませんでした。\n"+
				"got : %v\n"+
				"want: %v", y.String(), "2021")
		}
	})

	t.Run("異常なYearのString()メソッドが正しい値を返す", func(t *testing.T) {
		input := 2020

		y := NewYear(input)

		if y.String() != "Invalid Year" {
			t.Errorf("異常な年である2020年のYearのString()メソッドが 'Invalid Year' 以外の文字列を返しました。\n"+
				"got : %v\n"+
				"want: %v", y.String(), "Invalid Year")
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
