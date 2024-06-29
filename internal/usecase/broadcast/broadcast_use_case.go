package broadcast

import (
	"time"

	"github.com/thousanda/kamo-div/internal/internal/module/messaging"
	"github.com/thousanda/kamo-div/internal/internal/module/reminder"
)

type UseCase struct {
	now                   time.Time
	reminderMessageWriter reminder.MessageWriter
	messenger             messaging.Messenger
}

func NewUseCase(
	now time.Time,
	w reminder.MessageWriter,
	m messaging.Messenger,
) *UseCase {
	return &UseCase{
		now:                   now,
		reminderMessageWriter: w,
		messenger:             m,
	}
}

func (u *UseCase) ScheduledBroadcast() error {
	return nil
}
