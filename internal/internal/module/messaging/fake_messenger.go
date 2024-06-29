package messaging

import (
	"fmt"
)

type nothingDoer struct {
}

func NewDoNothingMessenger() Messenger {
	return &nothingDoer{}
}

func (n *nothingDoer) Broadcast() error {
	fmt.Println("Just Broadcast!! (うそです)")
	return nil
}

type failure struct {
}

func NewFailMessenger() Messenger {
	return &failure{}
}

func (f *failure) Broadcast() error {
	return fmt.Errorf("sorry, I can't broadcast")
}
