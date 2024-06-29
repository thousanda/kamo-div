package messaging

import (
	"fmt"

	"github.com/thousanda/kamo-div/internal/internal/module/reminder"
)

type nothingDoer struct {
}

func NewDoNothingMessenger() Messenger {
	return &nothingDoer{}
}

func (n *nothingDoer) Broadcast(r reminder.Reminder) error {
	fmt.Println("Just Broadcast!! (うそです)")
	return nil
}

type failure struct {
}

func NewFailMessenger() Messenger {
	return &failure{}
}

func (f *failure) Broadcast(r reminder.Reminder) error {
	return fmt.Errorf("sorry, I can't broadcast")
}
