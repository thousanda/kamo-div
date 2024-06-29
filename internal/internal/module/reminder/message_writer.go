package reminder

// MessageWriter 送るリマインダーメッセージを書くライターさんです
type MessageWriter struct {
}

func NewMessageWriter() *MessageWriter {
	return &MessageWriter{}
}
