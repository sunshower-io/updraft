package tokens

import (
	"github.com/stretchr/testify/assert"
	"github.com/sunshower-io/updraft/front/parser"
	"strings"
	"testing"
	"github.com/sunshower-io/updraft/common/observer"
)

func TestPascalTokenExtractsIdentifierCorrectly(t *testing.T) {
	source := parser.NewSource(
		strings.NewReader("coolbeans"),
		observer.NewEventProducer(),
	)

	token, _ := NewPascalToken(IDENTIFIER, source)
	assert.Equal(t, token.GetText(), "coolbeans")
}

func TestConsumingSingleDigitIntWorks(t *testing.T) {
	source := parser.NewSource(
		strings.NewReader("100"),
		observer.NewEventProducer(),
	)
	r, e := NewPascalNumber(source)
	assert.NoError(t, e)
	assert.Equal(t, r.GetValue(), 100)
}

func TestConsumingIntWorks(t *testing.T) {
	source := parser.NewSource(
		strings.NewReader("100"),
		observer.NewEventProducer(),
	)
	r, e := NewPascalNumber(source)
	assert.NoError(t, e)
	assert.Equal(t, r.GetValue(), 100)
}
