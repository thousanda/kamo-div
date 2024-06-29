package broadcast

type UseCase struct {
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) ScheduledBroadcast() error {
	return nil
}
