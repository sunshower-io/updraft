package tokens

import (
    "strings"
    "testing"
	"github.com/stretchr/testify/assert"
	"github.com/sunshower-io/updraft/front/parser"
	"github.com/sunshower-io/updraft/common/observer"
)


func TestPascalTokenExtractsPlusSymbolCorrectly(t *testing.T) {
    
    source := parser.NewSource(
        strings.NewReader("+"),
        observer.NewEventProducer(),
    )
    
    token, er := NewSymbol(source)
    
    assert.Nil(t, er)
   
    assert.Equal(t, token.GetType(), PLUS)
}


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
		strings.NewReader("1014"),
		observer.NewEventProducer(),
	)
	r, e := NewPascalNumber(source)
	assert.NoError(t, e)
	assert.Equal(t, r.GetValue(), int64(1014))
}

func TestConsumingLargeIntWithExponentWorks(t *testing.T) {
    source := parser.NewSource(
        strings.NewReader("22e6"),
        observer.NewEventProducer(),
    )
    _, e := NewPascalNumber(source)
    assert.NoError(t, e)
}


func TestConsumingLargeInt(t *testing.T) {
    source := parser.NewSource(
        strings.NewReader("10023241345"),
        observer.NewEventProducer(),
    )
    r, e := NewPascalNumber(source)
    assert.NoError(t, e)
    assert.Equal(t, r.GetValue(), int64(10023241345))
}


func TestConsumingIntWorks(t *testing.T) {
	source := parser.NewSource(
		strings.NewReader("0"),
		observer.NewEventProducer(),
	)
	r, e := NewPascalNumber(source)
	assert.NoError(t, e)
	assert.Equal(t, r.GetValue(), int64(0))
}
