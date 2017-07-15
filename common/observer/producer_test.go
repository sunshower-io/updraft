package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockMessage struct {
	Message
	id   EventType
	body interface{}
}

func (m *MockMessage) GetBody() interface{} {
	return m.body
}

func (m *MockMessage) TopicId() EventType {
	return m.id
}

func mockMessage(t EventType) Message {
	return &MockMessage{
		id:   t,
		body: t,
	}
}

type RecordingListener struct {
	EventListener
	id       string
	topicId  EventType
	messages []Message
}

func (f *RecordingListener) Id() string {
	return f.id
}

func (f *RecordingListener) ListensFor(m Message) bool {
	return m.TopicId() == f.topicId
}

func (f *RecordingListener) OnMessage(m Message) {
	if f.messages == nil {
		f.messages = make([]Message, 0)
	}
	f.messages = append(f.messages, m)
}

func TestRemovingListenerWorks(t *testing.T) {

	producer := NewEventProducer()
	listener := &RecordingListener{
		id:       "coolbeans",
		topicId:  "whatever",
		messages: make([]Message, 0),
	}

	producer.AddEventListener(listener)

	message := mockMessage("whatever")

	producer.RemoveMessageListener(listener)

	producer.SendMessage(message)

	assert.Equal(t, 0, len(listener.messages))
}

func TestAddingListenerWorks(t *testing.T) {
	producer := NewEventProducer()
	listener := &RecordingListener{
		id:       "coolbeans",
		topicId:  "whatever",
		messages: make([]Message, 0),
	}

	producer.AddEventListener(listener)

	message := mockMessage("whatever")

	producer.SendMessage(message)

	assert.Equal(t, 1, len(listener.messages))
}
