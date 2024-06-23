package deadline

type Year struct {
	year    int
	isValid bool
}

type Month struct {
	month   int
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
