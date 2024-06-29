package reminder

import (
	"fmt"

	"github.com/thousanda/kamo-div/internal/internal/module/deadline"
)

const sponsorUrl = "https://higuchi.world/gichiland-spot-sponsor"

// MessageWriter 送るリマインダーメッセージを書くライターさんです
type MessageWriter struct {
}

func NewMessageWriter() *MessageWriter {
	return &MessageWriter{}
}

func (w *MessageWriter) WriteReminderOneWeekAhead(d deadline.Deadline) Reminder {
	msgs := make([]string, 3)
	msgs[0] = "1週間後がスポットスポンサーの締切だよ！🦆"
	msgs[1] = fmt.Sprintf("締め切り: %s", d.ToJapaneseMDWdHMString())
	msgs[2] = sponsorUrl
	return NewReminder(msgs)
}

func (w *MessageWriter) WriteReminderOneDayAhead(d deadline.Deadline) Reminder {
	msgs := make([]string, 3)
	msgs[0] = "明日がスポットスポンサーの締切だよ！🦆"
	msgs[1] = fmt.Sprintf("締め切り: %s", d.ToJapaneseMDWdHMString())
	msgs[2] = sponsorUrl
	return NewReminder(msgs)
}

func (w *MessageWriter) WriteReminderOnTheDay(d deadline.Deadline) Reminder {
	msgs := make([]string, 3)
	msgs[0] = "今日がスポットスポンサーの締切だよ！🦆"
	msgs[1] = fmt.Sprintf("締め切り: %s", d.ToJapaneseMDWdHMString())
	msgs[2] = sponsorUrl
	return NewReminder(msgs)
}
