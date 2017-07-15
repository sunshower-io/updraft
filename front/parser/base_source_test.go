package parser

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"unicode/utf8"
	"github.com/stretchr/testify/assert"
	sio "gitlab.com/sunshower.io/updraft/common/io"
	"gitlab.com/sunshower.io/updraft/middle/core"
	"gitlab.com/sunshower.io/updraft/common/observer"
)

func TestEOFToken(t *testing.T) {
	value := ""
	reader := strings.NewReader(value)
	source := NewSource(
		reader,
		observer.NewEventProducer(),
	).(*BaseSource)

	token := core.EOF(source)
	assert.NotNil(t, token)

}

func TestPeekPeeks(t *testing.T) {

	value := "Hello, 世界"
	reader := strings.NewReader(value)
	source := NewSource(
		reader,
		observer.NewEventProducer(),
	).(*BaseSource)
	r, _ := source.Peek()

	assert.Equal(t, r, 'e')
	r, _ = source.CurrentCharacter()
	assert.Equal(t, r, 'H')
	r, _ = source.NextCharacter()
	assert.Equal(t, r, 'e')
	r, _ = source.CurrentCharacter()
	assert.Equal(t, r, 'e')
	r, _ = source.Peek()
	assert.Equal(t, r, 'l')
}

func TestPeekDoesNotIncrementLine(t *testing.T) {

	value := "Hello, 世界"
	reader := strings.NewReader(value)
	source := NewSource(
		reader,
		observer.NewEventProducer(),
	).(*BaseSource)
	source.CurrentCharacter()

	fstPos := source.currentPosition
	r, ch := source.Peek()
	assert.NoError(t, ch)
	assert.Equal(t, r, 'e')
	assert.Equal(t, fstPos, source.currentPosition)
}

func TestSourceWorksForUTF8(t *testing.T) {
	value := "Hello, 世界"
	reader := strings.NewReader(value)
	source := NewSource(
		reader,
		observer.NewEventProducer(),
	).(*BaseSource)

	last := 0
	for i, err := source.CurrentCharacter(); err == nil; i, err = source.nextCharacter() {
		r, size := utf8.DecodeRuneInString(value[last:])
		last += size
		println(fmt.Sprintf("mine: <%c> theirs: <%c>", i, r))
	}
}

func TestSourceReturnsEOFForEmptyString(t *testing.T) {
	value := ""
	reader := strings.NewReader(value)

	source := NewSource(
		reader,
		observer.NewEventProducer(),
	).(*BaseSource)

	_, err := source.Peek()
	assert.Equal(t, err, io.EOF)
}

func TestSourceProgression(t *testing.T) {
	value := "abcde"
	reader := strings.NewReader(value)

	source := NewSource(
		reader,
		observer.NewEventProducer(),
	).(*BaseSource)
	
	
	r, _ := source.currentCharacter()

	assert.EqualValues(t, r, 'a')

	r, _ = source.nextCharacter()
	assert.Equal(t, r, 'b')

	r, _ = source.nextCharacter()
	assert.Equal(t, r, 'c')

	r, _ = source.nextCharacter()
	assert.Equal(t, r, 'd')

	r, _ = source.currentCharacter()
	assert.Equal(t, r, 'd')

	r, _ = source.currentCharacter()
	assert.Equal(t, r, 'd')

	r, _ = source.nextCharacter()
	assert.Equal(t, r, 'e')

	r, _ = source.nextCharacter()
	assert.Equal(t, r, sio.EOL)
}

func TestPeekReturnsEOLForEOL(t *testing.T) {
	value := "\n"
	reader := strings.NewReader(value)
	source := NewSource(
		reader,
		observer.NewEventProducer(),
	)
	r, _ := source.Peek()
	assert.Equal(t, r, sio.EOL)

}

func TestReadLineReadsFirstLine(t *testing.T) {
	value := "abcde\nefgj"
	reader := strings.NewReader(value)
	source := NewSource(
		reader,
		observer.NewEventProducer(),
	).(*BaseSource)
	source.readLine()
	assert.Equal(t, source.line, "abcde")
}

func TestReadMultipleLinesWorks(t *testing.T) {
	value := "abcde\nefgj"
	reader := strings.NewReader(value)
	source := NewSource(
		reader,
		observer.NewEventProducer(),
	).(*BaseSource)
	source.readLine()
	assert.Equal(t, source.line, "abcde")
	source.readLine()
	assert.Equal(t, source.line, "efgj")
}



func TestReadingLinesProducesCorrectLineEvents(t *testing.T) {
	value := "abcde\nefgj\n     afdafadf\n\n\n  frapper \n"
	reader := strings.NewReader(value)
	
	source := NewSource(
		reader,
		observer.NewEventProducer(),
	).(*BaseSource)
	
	

	listener := &newlineCountingListener{}
	source.AddEventListener(listener)

	for {
		
		if source.IsEOF() {
			break
		}
		source.NextCharacter()
	}
	
	assert.Equal(t, listener.count, 6)
}


type newlineCountingListener struct {
	count int
}



func(s *newlineCountingListener) Id() string {
	return "newlines"
}

func(s *newlineCountingListener) ListensFor(m observer.Message) bool {
	return m.TopicId() == observer.SOURCE_LINE
}

func (s *newlineCountingListener) OnMessage(_ observer.Message) {
	s.count++
}