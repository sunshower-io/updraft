package tokens

import (
	"gitlab.com/sunshower.io/updraft/common/io"
	"gitlab.com/sunshower.io/updraft/middle/core"
	"unicode"
)

type PascalStringToken struct {
	*core.BaseToken
}

func (p *PascalStringToken) Extract() error {
	textBuffer := ""
	valueBuffer := ""

	source := p.Source

	ch, er := source.NextCharacter()
	textBuffer += "'"

	for {
		if unicode.IsSpace(ch) {
			ch = ' '
		}
		if !(ch == '\'' || er == io.EOF) {
			chs := string(ch)
			textBuffer += chs
			valueBuffer += chs
			ch, er = source.NextCharacter()
		}
		if ch == '\'' {
			for {
				pch, er := source.Peek()
				if er != nil {
					break
				}
				if ch == '\'' && pch == '\'' {
					textBuffer += "'"
					valueBuffer += string(ch)
					ch, er = source.NextCharacter()
					ch, er = source.NextCharacter()
				} else {
					break
				}
			}
		}

		if er != nil || ch == '\'' {
			break
		}
	}

	if ch == '\'' {
		textBuffer += "'"
		ch, er = source.NextCharacter()
	}
	p.Value = valueBuffer
	p.Text = textBuffer
	return er
}

func NewPascalString(s io.Source) (core.Token, error) {

	pt := &PascalStringToken{
		BaseToken: core.NewToken(s, STRING).(*core.BaseToken),
	}
	pt.Extract()
	return pt, nil
}
