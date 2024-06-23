package messaging

type Messenger interface {
	Broadcast() error
}
