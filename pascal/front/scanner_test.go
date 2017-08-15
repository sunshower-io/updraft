package front

import (
    "testing"
    "strings"
    "github.com/stretchr/testify/assert"
    "github.com/sunshower-io/updraft/front/parser"
    "github.com/sunshower-io/updraft/common/observer"
)

func TestSourceProducesAllCharacters(t *testing.T) {
    
    source := parser.NewSource(
        strings.NewReader("4+6"),
        observer.NewEventProducer(),
    )
    r, _ := source.CurrentCharacter()
    assert.Equal(t, r, rune('4'))
    r, _ = source.NextCharacter()
    assert.Equal(t, r, rune('+'))
    r, _ = source.NextCharacter()
    assert.Equal(t, r, rune('6'))
    
}

func TestLexingBinaryOperationWorks(t *testing.T) {
    source := parser.NewSource(
        strings.NewReader("two * four"), 
        observer.NewEventProducer(),
    )
 
    scanner := NewScanner(
        source, 
        observer.NewEventProducer(),
    )
    
    i := 0
    for {
        token := scanner.CurrentToken()
        if token != nil {
            println(token.GetText())
        }
        if i > 10 {
            return
        }
        token, _ = scanner.NextToken()

        i++
    }
        
        
    
    
    
}
