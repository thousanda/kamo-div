package messaging

import "github.com/thousanda/kamo-div/internal/internal/module/reminder"

type Messenger interface {
	Broadcast(reminder reminder.Reminder) error
}
