package deadline

import "fmt"

// Year 年を表すもの
// 番組の開始が2021年なので、それ以降のみを表現できる
// 必ずNewYear()を使って作成しましょう
type Year int

func NewYear(y int) (Year, error) {
	if y < 2021 {
		return Year(0), fmt.Errorf("year must be greater than or equal to 2021; got %d", y)
	}
	return Year(y), nil
}

func (y Year) Value() int {
	return int(y)
}

// Month 月を表すもの
// 1月から12月までを表現できる
type Month int

func NewMonth(m int) (Month, error) {
	if m < 1 || 12 < m {
		return Month(0), fmt.Errorf("month must be between 1 and 12; got %d", m)
	}
	return Month(m), nil
}

func (m Month) Value() int {
	return int(m)
}
