package core

import (
	"fmt"
	"github.com/sunshower-io/updraft/common/io"
	"github.com/sunshower-io/updraft/common/observer"
)

type TokenType struct {
	Name  string
	Value string
}

func (t TokenType) String() string {
	if t.Name == t.Value {
		return fmt.Sprintf("Token('%s')", t.Value)
	}
	return fmt.Sprintf("Token('%s':'%s')", t.Name, t.Value)
}

func CreateTT(name string) TokenType {
	return NewTokenType(name, name)
}

func SymbolTokenType(
	name, value string,
	assignment map[string]TokenType,
) TokenType {
	result := NewTokenType(name, value)
	assignment[value] = result
	return result
}

func ReservedWord(
	name string,
	set map[string]TokenType) TokenType {
	return SymbolTokenType(name, name, set)
}

func NewTokenType(name, value string) TokenType {
	return TokenType{
		Name:  name,
		Value: value,
	}
}

var (
	ERROR_TOKEN = NewTokenType("ERROR", "error")
)

func NewTokenMessage(t Token) *TokenMessage {
	return &TokenMessage{
		Position:   t.GetPosition(),
		LineNumber: t.GetLineNumber(),
		Type:       t.GetType(),
		Text:       t.GetText(),
		Value:      t.GetValue(),
		Message: &observer.BaseEvent{
			Body:  t,
			Topic: observer.TOKEN,
		},
	}
}

type TokenMessage struct {
	observer.Message

	Position int

	LineNumber int

	Type TokenType

	Text string

	Value interface{}
}

func (m *TokenMessage) Format() string {
	return fmt.Sprintf(">>> %s line=%03d, pos=%2d, text=\"%s\"\n>>>        value=%s",
		m.Type,
		m.LineNumber,
		m.Position,
		m.Text,
		m.Value,
	)
}

type Token interface {
	
	
	GetType() TokenType

	GetText() string

	GetPosition() int

	GetLineNumber() int

	GetValue() interface{}

	GetSource() io.Source

	Extract() error
}

type BaseToken struct {
	Token

	Source io.Source

	Type TokenType

	Value interface{}

	position int

	lineNumber int

	Text string
	
}

func (b *BaseToken) String() string {
	return fmt.Sprintf(
		"TOKEN{text: %s, line: %d, col: %d, type: %s, value: '%s'}",
		b.Text,
		b.lineNumber,
		b.position,
		b.Type,
		b.Value,
	)
}

func (b *BaseToken) GetValue() interface{} {
	return b.Value
}

func (b *BaseToken) GetText() string {
	return b.Text
}

func (b *BaseToken) GetType() TokenType {
	return b.Type
}

func (b *BaseToken) GetLineNumber() int {
	return b.lineNumber
}

func (b *BaseToken) Extract() error {
	r, err := b.Source.CurrentCharacter()
	if err != nil {
		return err
	}
	b.Text = string(r)
	b.Value = nil
	_, err = b.Source.NextCharacter()
	if err != nil {
		return err
	}

	return nil
}

func (t *BaseToken) GetPosition() int {
	return t.position
}


func CreateToken(
		s io.Source,
		t TokenType,
		text  string,
		value interface{},
) Token {
	result := &BaseToken{
		Text        : text,
		Type		: t,
		Source		: s,
		position	: s.GetPosition(),
		lineNumber	: s.GetLineNumber(),
		Value	    : value,
	}
	return result
}

func NewToken(s io.Source, t TokenType) Token {
	result := &BaseToken{
		Type:       t,
		Source:     s,
		position:   s.GetPosition(),
		lineNumber: s.GetLineNumber(),
	}
	return result
}

func FromSource(s io.Source, t Token) Token {
	result := &BaseToken{
		Source     : s,
		position   : s.GetPosition(),
		lineNumber : s.GetLineNumber(),
	}
	result.Extract()
	return result
}
