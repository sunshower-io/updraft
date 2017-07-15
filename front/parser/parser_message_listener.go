package parser

import "gitlab.com/sunshower.io/updraft/common/observer"

type ParserMessageListener struct {
	observer.EventListener
}

func (l *ParserMessageListener) Id() string {
	return "token-parser-message-listener"
}

func (l *ParserMessageListener) ListensFor(m observer.Message) bool {

	return true
}

func (l *ParserMessageListener) OnMessage(m observer.Message) {
	tid := m.TopicId()
	switch tid {
	case observer.TOKEN,
		observer.SYNTAX_ERROR:
		println(m.Format())
	}
}
