package deadline

import "fmt"

// StreamingMonth エピソードが配信される時期を表すもの
// 例) 2021-04
// 番組の開始が2021年4月なので、それ以降のみを表現できる
// 必ずNewStreamingMonth()を使って作成しましょう
type StreamingMonth struct {
	year  Year
	month Month
}

func NewStreamingMonth(y Year, m Month) (StreamingMonth, error) {
	if y == 2021 && m < 4 {
		return StreamingMonth{},
			fmt.Errorf("year and month must be greater than or equal to 2021-04; got %d-%d", y, m)
	}
	return StreamingMonth{
		year:  y,
		month: m,
	}, nil
}

func (sm StreamingMonth) Year() Year {
	return sm.year
}

func (sm StreamingMonth) Month() Month {
	return sm.month
}

func (sm StreamingMonth) String() string {
	return fmt.Sprintf("%d-%02d", sm.year, sm.month)
}
