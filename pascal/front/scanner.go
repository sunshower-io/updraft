package front

import (
	"io"
	"unicode"
	"gitlab.com/sunshower.io/updraft/middle/core"
	"gitlab.com/sunshower.io/updraft/pascal/tokens"
	"gitlab.com/sunshower.io/updraft/common/frontend"
	"gitlab.com/sunshower.io/updraft/common/comments"
	io2 "gitlab.com/sunshower.io/updraft/common/io"
	"gitlab.com/sunshower.io/updraft/common/observer"
)

type PascalScanner struct {
	frontend.Scanner
}

type pascalExtractor struct {
	observer.EventProducer
	scanner PascalScanner
}

func (e *pascalExtractor) skipWhitespace(
		scanner frontend.Scanner,
) (core.Token, error) {
	
	var (
		er error
		ch rune
	)
	for ch, er = scanner.CurrentCharacter();
			er == nil && (unicode.IsSpace(ch) || ch == '{'); {
		if ch == '{' {
			comment := ""
			for ch, er = scanner.CurrentCharacter(); ch != '}'; {
				comment += string(ch)
				if er == io.EOF {
					return tokens.NewError(
						scanner.Source(),
						tokens.UNEXPECTED_EOF,
						"}",
					), nil
				}
				ch, er = scanner.NextCharacter()
			}
			if ch == '}' {
				ch, er = scanner.NextCharacter()
			}
			scanner.SendMessage(comments.CommentScanned(comment))
		} else {
			ch, er = scanner.NextCharacter()
		}
	}

	return nil, er
}

func (e *pascalExtractor) Extract(
		scanner frontend.Scanner,
) (core.Token, error) {
	
	var token core.Token
	source := scanner.Source()
	token, er := e.skipWhitespace(scanner)
	
	if token != nil {
		return token, nil
	}
	
	ch, er := scanner.CurrentCharacter()


	if er == io.EOF {
		return core.EOF(source), nil
	} else if unicode.IsLetter(ch) {
		token, er = tokens.NewPascalToken(tokens.IDENTIFIER, source)
	} else if unicode.IsDigit(ch) {
		token, er = tokens.NewPascalNumber(source)
	} else if ch == '\'' {
		token, er = tokens.NewPascalString(source)
	} else if tokens.IsSymbol(ch) {
		token, er = tokens.NewSymbol(source)
	} else {
		token = tokens.NewError(
			source,
			tokens.INVALID_CHARACTER,
			string(ch),
		)
		source.NextCharacter()
	}

	return token, er
}


func NewScanner(
		source io2.Source,
		producer observer.EventProducer,
) frontend.Scanner {
	scanner := frontend.NewScanner(source, producer)
	extractor := new(pascalExtractor)
	extractor.EventProducer = producer
	scanner.SetExtractor(extractor)
	return scanner
}