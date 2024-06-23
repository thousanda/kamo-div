package deadline

import "fmt"

// Year 年を表す構造体
// 番組の開始が2021年なので、それ以降のみを表現できる
type Year struct {
	year    int
	isValid bool
}

func NewYear(y int) Year {
	if y < 2021 {
		return Year{
			isValid: false,
		}
	}
	return Year{
		year:    y,
		isValid: true,
	}
}

func (y Year) Value() int {
	return y.year
}

func (y Year) IsValid() bool {
	return y.isValid
}

func (y Year) String() string {
	if !y.IsValid() {
		return fmt.Sprintf("Invalid Year")
	}
	return fmt.Sprintf("%d", y.Value())
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
