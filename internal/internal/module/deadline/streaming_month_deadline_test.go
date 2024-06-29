package deadline

import (
	"testing"
	"time"
)

func TestStreamingMonth_Deadline(t *testing.T) {
	t.Run("締め切りは放送月の第一火曜日の13日前にあたる日の23:59:59を返す", func(t *testing.T) {
		// テスト対象のものを用意
		sm, err := createTestStreamingMonth(2024, 7, time.UTC)
		if err != nil {
			t.Errorf("StreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		d := sm.Deadline()

		// 結果を検証
		want := NewDeadline(time.Date(2024, 6, 19, 23, 59, 59, 0, time.UTC))
		if !d.Equal(want) {
			t.Errorf("2024年7月放送分の締め切りは6月19日の23:59:59のはずですが、異なる時刻が計算されました。\n"+
				"got : %v\n"+
				"want: %v", d, want)
		}
	})
}

func TestStreamingMonth_FirstTuesday(t *testing.T) {
	t.Run("放送月の第一火曜日を計算して返す", func(t *testing.T) {
		// テスト対象のものを用意
		sm, err := createTestStreamingMonth(2024, 7, time.UTC)
		if err != nil {
			t.Errorf("StreamingMonthを作成しようとした時点でerrが返ってきてしまいました。\n"+
				"err: %v", err)
		}

		// テスト対象のメソッドを実行
		d := sm.firstTuesday()

		// 結果を検証
		want := time.Date(2024, 7, 2, 0, 0, 0, 0, time.UTC)
		if !d.Equal(want) {
			t.Errorf("2024年7月の第一火曜日は7月2日の0:00:00のはずですが、異なる時刻が計算されました。\n"+
				"got : %v\n"+
				"want: %v", d, want)
		}
	})
}
