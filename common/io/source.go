package io

import (
	"io"
	"gitlab.com/sunshower.io/updraft/common/observer"
)

const (
	EOL = '\n'
)

var (
	EOF = io.EOF
)

type SourcePosition struct {
	Line       string
	LineNumber int
}

type Source interface {
	observer.EventProducer

	/**
	Get the current line number
	 */
	GetLineNumber() int

	/**
	Determine if this source has more characters
	 */
	IsEOF() bool

	
	/**
	Get the position (current column) of the source
	 */
	GetPosition() int

	/**

	return the current character
	 */
	CurrentCharacter() (rune, error)

	/**

	 */
	NextCharacter() (rune, error)

	/**

	 */
	Peek() (rune, error)
}
