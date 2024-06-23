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
