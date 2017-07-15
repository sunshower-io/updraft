package observer

type EventListener interface {
	Id() string

	OnMessage(Message)

	ListensFor(Message) bool
}
