package tokens

import (
	"bytes"
	"github.com/sunshower-io/updraft/common/io"
	"github.com/sunshower-io/updraft/middle/core"
	"unicode"
)


type PascalNumberToken struct {
	*core.BaseToken
}

func (p *PascalNumberToken) Extract() error {

	b := new(bytes.Buffer)
	source := p.Source

	s, er := p.unsignedInteger(b, source)

	sawSpanOp := false
	if p.Type == core.ERROR_TOKEN || er != nil {
		return er
	}

	ch, er := source.CurrentCharacter()

	//var (
	//    fraction string
	//    exponent rune
	//    exponentPart string
	//)
	//
	if ch == '.' {
		pch, _ := source.Peek()
		if pch == '.' {
			sawSpanOp = true
		} else {
			p.Type = REAL
			b.WriteRune(ch)
			ch, er = source.NextCharacter()
			_, er = p.unsignedInteger(b, source)
			if p.Type == core.ERROR_TOKEN {
				return er
			}
		}
	}
	ch, er = source.CurrentCharacter()

	if !sawSpanOp && (ch == 'e' || ch == 'E') {
		p.Type = REAL
		b.WriteRune(ch)
	}
	ch, er = source.NextCharacter()

	if ch == '+' || ch == '-' {
		b.WriteRune(ch)
		ch, er = source.NextCharacter()
	}

	_, er = p.unsignedInteger(b, source)

	p.Text = s

	p.Type = INTEGER
	//p.Value, er = strconv.Atoi(s)
	p.Value = 100

	return er
}

func (p *PascalNumberToken) unsignedInteger(buffer *bytes.Buffer, source io.Source) (string, error) {

	ch, er := source.CurrentCharacter()
	var buf bytes.Buffer
	if (er != nil && !(er == io.EOF)) || !unicode.IsDigit(ch) {
		p.Type = core.ERROR_TOKEN
		p.Value = INVALID_NUMBER
	}

	for {
		if !unicode.IsDigit(ch) || er != nil {
			break
		}
		buffer.WriteRune(ch)
		buf.WriteRune(ch)
		ch, er = source.NextCharacter()

	}
	return buf.String(), er
}

func NewPascalNumber(s io.Source) (core.Token, error) {
	pt := &PascalNumberToken{
		BaseToken: core.NewToken(s, STRING).(*core.BaseToken),
	}
	pt.Extract()
	return pt, nil
}
