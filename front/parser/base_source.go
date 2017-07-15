package parser

import (
	"bufio"
    gio "io"
    "unicode/utf8"
	"github.com/sunshower-io/updraft/common/io"
	"github.com/sunshower-io/updraft/common/observer"
)

type BaseSource struct {
    
    observer.EventProducer

	eof             bool
	line            string
	lineNumber      int
	currentPosition int
	reader          *bufio.Reader
}





func (s *BaseSource) GetLineNumber() int {
	return s.lineNumber
}




func (s *BaseSource) GetPosition() int {
	return s.currentPosition
}






func (s *BaseSource) NextCharacter() (rune, error) {
	return s.nextCharacter()
}





func (s *BaseSource) CurrentCharacter() (rune, error) {
	return s.currentCharacter()
}




func (s *BaseSource) Peek() (rune, error) {
	s.currentCharacter()
	if s.eof {
		return 0, io.EOF
	}

	nextPosition := s.currentPosition + 1

	if nextPosition < len(s.line) {
		i, _ := utf8.DecodeRuneInString(s.line[nextPosition:])
		return i, nil
	}
	return io.EOL, nil
}






func (s *BaseSource) currentCharacter() (rune, error) {
	if s.currentPosition == -2 {
		er := s.readLine()
		if er != nil {
			s.eof = true
			return 0, io.EOF
		}
		return s.nextCharacter()
	} else if s.eof {
		return 0, io.EOF
	} else if s.currentPosition == -1 || s.currentPosition == len(s.line) {
		return io.EOL, nil
	} else if s.currentPosition > len(s.line) {
		er := s.readLine()
		if er != nil {
			s.eof = true
			return 0, io.EOF
		}
		return s.nextCharacter()
	} else {
		return s.next(), nil
	}
}





func (s *BaseSource) next() rune {
	i, l := utf8.DecodeRuneInString(s.line[s.currentPosition:])
	if l > 1 {
		s.currentPosition += l - 1
	}
	return i
}





func (s *BaseSource) nextCharacter() (rune, error) {
	s.currentPosition += 1
	return s.currentCharacter()
}





func (s *BaseSource) readLine() error {
	line, _, err := s.reader.ReadLine()
	if err != nil {
		s.eof = true
		return io.EOF
	}
	s.currentPosition = -1
	s.lineNumber += 1

	l := string(line[:])

	s.SendMessage(observer.CreateEvent(
		observer.SOURCE_LINE,
		&io.SourcePosition{
			Line:       l,
			LineNumber: s.lineNumber,
		}))

	s.line = l
	return nil
}





func (s *BaseSource) IsEOF() bool {
	return s.eof
}





func NewSource(
        reader gio.Reader,
        producer observer.EventProducer,
) io.Source {
 
 

	return &BaseSource {
       EventProducer    : producer,
		reader          : bufio.NewReader(reader),
		lineNumber      : 0,
		currentPosition : -2,
	}
 
}
