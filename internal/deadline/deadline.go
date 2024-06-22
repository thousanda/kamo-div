package deadline

type Deadline struct {
	streamingMonth *StreamingMonth
}

func NewDeadline(sm StreamingMonth) Deadline {
	return Deadline{
		streamingMonth: &sm,
	}
}
