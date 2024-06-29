package reminder

// Reminder リマインダーのメッセージです
type Reminder []string

func NewReminder(s []string) Reminder {
	clone := make([]string, len(s))
	copy(clone, s)
	return clone
}

func (r Reminder) ToStrings() []string {
	clone := make([]string, len(r))
	copy(clone, r)
	return clone
}
