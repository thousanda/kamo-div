package deadline

import "fmt"

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

func (sm StreamingMonth) IsEqualTo(other StreamingMonth) bool {
	return sm.year == other.year && sm.month == other.month
}
