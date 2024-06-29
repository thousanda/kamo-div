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
