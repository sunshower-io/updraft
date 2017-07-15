package observer

import "fmt"

type EventType string

const (
	SYNTAX_ERROR        EventType = "compiler:syntax-error"
	COMMENT 			EventType = "source::comment"
	TOKEN               EventType = "source::token-event"
	SOURCE_LINE         EventType = "source::line"
	SOURCE_LINE_FORMAT  EventType = "source::line-format"
	PARSER_SUMMARY      EventType = "parser::summary"
	COMPILER_SUMMARY    EventType = "compiler::summary"
	INTERPRETER_SUMMARY EventType = "interpreter::summary"
)

func CreateEvent(t EventType, body interface{}) Message {
	return &BaseEvent{
		Topic: t,
		Body:  body,
	}
}

type BaseEvent struct {
	Message

	Topic EventType

	Body interface{}
}

func (b *BaseEvent) TopicId() EventType {
	return b.Topic
}

func (b *BaseEvent) Format() string {
	return fmt.Sprintf("Message{%T %s}", b, b)
}

func (b *BaseEvent) GetBody() interface{} {
	return b.Body
}
