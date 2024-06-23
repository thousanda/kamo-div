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

// Month 月を表す構造体
// 1月から12月までを表現できる
type Month struct {
	month   int
	isValid bool
}

func NewMonth(m int) Month {
	if m < 1 || 12 < m {
		return Month{
			isValid: false,
		}
	}
	return Month{
		month:   m,
		isValid: true,
	}
}

func (m Month) Value() int {
	return m.month
}

func (m Month) IsValid() bool {
	return m.isValid
}

func (m Month) String() string {
	if !m.IsValid() {
		return fmt.Sprintf("Invalid Month")
	}
	return fmt.Sprintf("%d", m.Value())
}
