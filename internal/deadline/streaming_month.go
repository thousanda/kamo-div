package deadline

import (
	"fmt"
	"time"
)

// StreamingMonth エピソードが配信される時期を表すもの
// 例) 2021-04
// 番組の開始が2021年4月なので、それ以降のみを表現できる
// 必ずNewStreamingMonth()を使って作成しましょう
type StreamingMonth struct {
	year  Year
	month Month
	loc   *time.Location
}

func NewStreamingMonth(y Year, m Month, loc *time.Location) (StreamingMonth, error) {
	if y == 2021 && m < 4 {
		return StreamingMonth{},
			fmt.Errorf("year and month must be greater than or equal to 2021-04; got %d-%d", y, m)
	}
	if loc == nil {
		return StreamingMonth{},
			fmt.Errorf("time.Location is required")
	}
	return StreamingMonth{
		year:  y,
		month: m,
		loc:   loc,
	}, nil
}

func (sm StreamingMonth) Year() Year {
	return sm.year
}

func (sm StreamingMonth) Month() Month {
	return sm.month
}

func (sm StreamingMonth) Location() *time.Location {
	return sm.loc
}

func (sm StreamingMonth) String() string {
	return fmt.Sprintf("%d-%02d", sm.year, sm.month)
}

// Deadline スポンサーやお便りの締め切り時刻を計算して返す
// 例) 2024-06-19 23:59:59
// 放送月の第一火曜日の13日前にあたる日の23:59:59
// つまり、締め切り日は水曜日
func (sm StreamingMonth) Deadline() time.Time {
	return time.Time{}
}

// firstTuesday 放送月の第一火曜日を計算して返す
// 例) 2024-07-02 00:00:00
// 締め切りの計算に使う
func (sm StreamingMonth) firstTuesday() time.Time {
	// まずは放送月の1日を取得
	firstDay := time.Date(sm.Year().Value(), time.Month(sm.Month()),
		1, 0, 0, 0, 0, sm.Location())
	// 1日の曜日を取得
	// 0: 日曜, 1: 月曜, 2: 火曜, ..., 6: 土曜
	dayOfWeekOfFirstDay := firstDay.Weekday()
	// 何日進めば火曜になるかを計算
	offset := (7 + 2 - int(dayOfWeekOfFirstDay)) % 7

	return time.Date(sm.Year().Value(), time.Month(sm.Month()),
		1+offset, 0, 0, 0, 0, sm.Location())
}
