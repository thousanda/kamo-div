package deadline

import (
	"fmt"
	"testing"
	"time"
)

// createTestStreamingMonth StreamingMonthのメソッドをテストしたいときのために、StreamingMonthの作成をすこし簡略化するヘルパー関数
func createTestStreamingMonth(y int, m int) (StreamingMonth, error) {
	year, err := NewYear(y)
	if err != nil {
		return StreamingMonth{},
			fmt.Errorf("StreamingMonthを作るためのYearを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
	}
	month, err := NewMonth(m)
	if err != nil {
		return StreamingMonth{},
			fmt.Errorf("StreamingMonthを作るためのMonthを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
	}
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return StreamingMonth{},
			fmt.Errorf("StreamingMonthを作るためのtime.Locationを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
	}

	return NewStreamingMonth(year, month, loc)
}

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
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			t.Errorf("入力のtime.Locationを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象の関数を実行
		sm, err := NewStreamingMonth(y, m, loc)

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
		if sm.loc != loc {
			t.Errorf("StreamingMonthのlocが指定したtime.Locationと異なります。\n"+
				"got : %v\n"+
				"want: %v", sm.loc, loc)

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
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			t.Errorf("入力のtime.Locationを用意しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象の関数を実行
		_, e := NewStreamingMonth(y, m, loc)

		// 結果を検証
		if e == nil {
			t.Errorf("2021年4月より前の年月を持つStreamingMonthを作成しようとしたのにerrが返りませんでした。")
		}
	})
}

func TestStreamingMonth_Year(t *testing.T) {
	t.Run("StreamingMonthのYear()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		sm, err := createTestStreamingMonth(2021, 4)
		if err != nil {
			t.Errorf("StreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		got := sm.Year()

		// 結果を検証
		want := Year(2021)
		if got != want {
			t.Errorf("StreamingMonthのYear()メソッドが正しい値を返しませんでした。\n"+
				"got : %d\n"+
				"want: %d", got, want)
		}
	})
}

func TestStreamingMonth_Month(t *testing.T) {
	t.Run("StreamingMonthのMonth()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		sm, err := createTestStreamingMonth(2021, 4)
		if err != nil {
			t.Errorf("StreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		got := sm.Month()

		// 結果を検証
		want := Month(4)
		if got != want {
			t.Errorf("StreamingMonthのMonth()メソッドが正しい値を返しませんでした。\n"+
				"got : %d\n"+
				"want: %d", got, want)
		}
	})
}

func TestStreamingMonth_Location(t *testing.T) {
	t.Run("StreamingMonthのLocation()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		sm, err := createTestStreamingMonth(2021, 4)
		if err != nil {
			t.Errorf("StreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		got := sm.Location()

		// 結果を検証
		want, _ := time.LoadLocation("Asia/Tokyo")
		if got.String() != want.String() {
			t.Errorf("StreamingMonthのLocation()メソッドが正しい値を返しませんでした。\n"+
				"got : %v\n"+
				"want: %v", got, want)
		}
	})
}

func TestStreamingMonth_String(t *testing.T) {
	t.Run("StreamingMonthのString()メソッドが正しい値を返す", func(t *testing.T) {
		// テスト対象のものを用意
		sm, err := createTestStreamingMonth(2021, 4)
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
