package frontend

import (
	"gitlab.com/sunshower.io/updraft/common/io"
	"gitlab.com/sunshower.io/updraft/common/observer"
	"gitlab.com/sunshower.io/updraft/middle/core"
)

type Extractor interface {
	Extract(Scanner) (core.Token, error)
}

type Scanner interface {
	observer.EventProducer

	Source() io.Source

	/**

	 */
	SetExtractor(Extractor)

	/**

	 */
	CurrentToken() core.Token

	/**

	 */
	NextToken() (core.Token, error)

	CurrentCharacter() (rune, error)

	NextCharacter() (rune, error)

	Peek() (rune, error)
}

type BaseScanner struct {
	Scanner
	observer.EventProducer

	extractor    Extractor
	source       io.Source
	currentToken core.Token
}

func (s *BaseScanner) Peek() (rune, error) {
	return s.source.Peek()
}

func (s *BaseScanner) Source() io.Source {
	return s.source
}

func (s *BaseScanner) SetExtractor(e Extractor) {
	s.extractor = e
}

func (s *BaseScanner) CurrentToken() core.Token {
	return s.currentToken
}

func (s *BaseScanner) CurrentCharacter() (rune, error) {
	return s.source.CurrentCharacter()
}

func (s *BaseScanner) NextCharacter() (rune, error) {
	return s.source.NextCharacter()
}

func (s *BaseScanner) NextToken() (core.Token, error) {
	if s.extractor == nil {
		panic("No way to extract the next token!  (extractor must not be nil)")
	}
	token, e := s.extractor.Extract(s)
	s.currentToken = token
	return token, e
}

func (s *BaseScanner) SendMessage(m observer.Message) {
	s.EventProducer.SendMessage(m)
}

func (s *BaseScanner) RemoveMessageListener(l observer.EventListener) {
	s.EventProducer.AddEventListener(l)
}

func (s *BaseScanner) AddEventListener(l observer.EventListener) {
	s.EventProducer.AddEventListener(l)
}

func NewScanner(source io.Source, producer observer.EventProducer) Scanner {
	return &BaseScanner{
		source		  	: source,
		EventProducer	: producer,
	}
}
