package tokens

import (
	"bytes"
	"github.com/sunshower-io/updraft/common/io"
	"github.com/sunshower-io/updraft/middle/core"
	"unicode"
    "math"
)


type PascalNumberToken struct {
	*core.BaseToken
}

func (p *PascalNumberToken) Extract() error {
    
    b := new(bytes.Buffer)
    source := p.Source
    
    return p.extractNumber(source, b)
}

func (p *PascalNumberToken) extractNumber(
        source io.Source, 
        buffer *bytes.Buffer,
) error {
    
    var (
        whole string
        fractional string
        expononents string 
        exponentSign = '+'
        sawSpread = false
        current rune
        valueType = INTEGER
    )
    
    
    
    whole, er := p.extractUnsigned(source, buffer)
    
    if er != nil {
        return er
    }
    
    current, er = source.CurrentCharacter()
    
    if current == '.' {
        ch, er := source.Peek()
        if ch == '.' {
            sawSpread = true
        } else {
            valueType = REAL
            buffer.WriteRune(current)
            current, er = source.NextCharacter()
            
            fractional, er = p.extractUnsigned(source, buffer)
            
            if er != nil {
                return er
            }
        }
    }
    
    current, er = source.CurrentCharacter()
    
    if !sawSpread && (current == 'E' || current == 'e') {
        valueType = REAL
        buffer.WriteRune(current)
        current, er = source.NextCharacter()
        
        if current == '+' || current == '-' {
            buffer.WriteRune(current)
            exponentSign = current
            current, er = source.NextCharacter()
        }
        
        expononents, er = p.extractUnsigned(source, buffer)
        
    }
    
    
    if valueType == INTEGER {
        intValue := p.computeIntegralValue(whole)
        if p.Type != core.ERROR_TOKEN {
            p.Value = intValue
            p.Type = INTEGER 
        }
    }
    
    if valueType == REAL {
        floatValue := p.computeFloatingValue(
            whole, 
            fractional, 
            expononents, 
            exponentSign,
        )
        
        if p.Type != core.ERROR_TOKEN {
            p.Value = floatValue
        }
        p.Type = REAL
        
    }
    
    
    return er
}


func (p*PascalNumberToken) computeFloatingValue(
        whole, 
        fractional, 
        exponents string, 
        sign rune,
) interface{} {
    
    
   
    var value float64 = 0.0
    
    
    var digits bytes.Buffer
    
    digits.WriteString(whole)
    
    exponentValue := p.computeIntegralValue(exponents)
    
    if sign == '-' {
        exponentValue = -exponentValue
    }
    
    if fractional != "" {
        exponentValue -= int64(len(fractional))
        digits.WriteString(fractional)
    }
    
    if IntAbs(exponentValue + int64(len(whole))) > 11 {
        p.Type = core.ERROR_TOKEN
        p.Value = RANGE_REAL
        return 0.0
    }
    
   
    ds := digits.String()
    for i := 0; i < len(ds) ; i++ {
        value = 10 * value + float64(int(ds[i] - '0'))
    }
    
    if exponentValue != 0 {
        value += math.Pow(10, float64(exponentValue))
    }
    
    return value
}


func IntAbs(v int64) int64 {
    if v < 0 {
        return -v
    }
    return v
}


func (p*PascalNumberToken) computeIntegralValue(i string) int64  {
    if i == "" {
        return 0
    }
    var (
        value int64 = 0
        previous int64 = -1 
        index = 0
    )
   
    for {
        
        if index >= len(i) || value < previous {
            break
        }
        
        previous = value 
        value = 10 * value + int64(i[index])  - '0'
        index++
    }
    
    if value >= previous {
        return value
    } else {
        p.Type = core.ERROR_TOKEN
        p.Value = RANGE_INTEGER
        return 0
    }
    
    
    
    
    
    
}


func (p *PascalNumberToken) extractUnsigned(
        source io.Source,
        textBuffer *bytes.Buffer,
) (string, error) {
    
    current, er := source.CurrentCharacter()
    
    if !unicode.IsDigit(current) {
        p.Type = core.ERROR_TOKEN
        p.Value = INVALID_NUMBER
    } 
    
    var (
        buffer bytes.Buffer 
    )
    
    for {
        if !unicode.IsDigit(current) {
            break
        }
        
        buffer.WriteRune(current)
        textBuffer.WriteRune(current)
        current, er = source.NextCharacter()
    }
    
    return buffer.String(), er
}


func NewPascalNumber(s io.Source) (core.Token, error) {
	pt := &PascalNumberToken{
		BaseToken: core.NewToken(s, STRING).(*core.BaseToken),
	}
	pt.Extract()
	return pt, nil
}
