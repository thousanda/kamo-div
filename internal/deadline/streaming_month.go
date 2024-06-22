package deadline

type Year int
type Month int

type StreamingMonth struct{
	year Year
	month Month
}

func NewStreamingMonth(y Year, m Month) StreamingMonth {
	return StreamingMonth{
		year: y,
		month: m,
	}
}
