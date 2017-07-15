package tokens

import (
	"gitlab.com/sunshower.io/updraft/common/io"
	"gitlab.com/sunshower.io/updraft/middle/core"
)

type PascalSymbolToken struct {
	*core.BaseToken
}

func NewSymbol(s io.Source) (core.Token, error) {
	pt := &PascalSymbolToken{
		BaseToken: core.NewToken(s, LABEL).(*core.BaseToken),
	}
	pt.Extract()
	return pt, nil
}

func (t *PascalSymbolToken) Extract() error {

	source := t.Source
	ch, er := source.CurrentCharacter()

	t.Text = string(ch)
	switch ch {
	case
		'+', '-', '*', '/',
		',', ';', '\'', '=',
		'(', ')', '[', ']',
		'{', '}', '^':
		ch, er = source.NextCharacter()
	case ':':
		ch, er = source.NextCharacter()
		switch ch {
		case '=':
			t.Text += string('=')
			source.NextCharacter()
		}
	case '<':
		ch, er = source.NextCharacter()
		switch ch {
		case '=', '>':
			t.Text += string(ch)
			source.NextCharacter()
		}

	case '>':
		ch, er = source.NextCharacter()
		switch ch {
		case '=':
			t.Text += string(ch)
			source.NextCharacter()
		}
	case '.':
		ch, er = source.NextCharacter()
		switch ch {
		case '.':
			t.Text += string(ch)
			source.NextCharacter()
		}

	default:
		ch, er = source.NextCharacter()
		t.Type = core.ERROR_TOKEN
		t.Value = INVALID_CHARACTER
	}

	t.Type = SPECIAL_TOKENS[t.Text]

	return er
}
