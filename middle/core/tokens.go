package core

import "gitlab.com/sunshower.io/updraft/common/io"

type EofToken struct {
	*BaseToken
}

func (e *EofToken) Extract() error {
	return nil
}

func EOF(source io.Source) Token {
	token := new(EofToken)
    token.BaseToken = FromSource(source, token).(*BaseToken)
	return token
}
