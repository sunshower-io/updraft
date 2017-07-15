package observer

type Message interface {
	GetBody() interface{}

	TopicId() EventType

	Format() string
}
