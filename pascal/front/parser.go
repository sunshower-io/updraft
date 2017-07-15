package front

import (
	"time"
	"strings"
	"gitlab.com/sunshower.io/updraft/common/ir"
	"gitlab.com/sunshower.io/updraft/common/observer"
	"gitlab.com/sunshower.io/updraft/middle/core"
	"gitlab.com/sunshower.io/updraft/pascal/tokens"
	"gitlab.com/sunshower.io/updraft/common/frontend"
	"gitlab.com/sunshower.io/updraft/common/compiler"
	ccore "gitlab.com/sunshower.io/updraft/common/core"
)

type RecursiveDescentPascalParser struct {
	frontend.Parser

	errorHandler 		compiler.ErrorHandler
	symbolTables  		ir.SymbolTableStack
	symbolFactory 		ir.SymbolTableFactory
}

func (p *RecursiveDescentPascalParser) GetSymbolTables() ir.SymbolTableStack {
	return p.symbolTables
}

func (p *RecursiveDescentPascalParser) Parse(
		ccore.CompilationResult,
) error {

	startTime := time.Now()

	var token core.Token

	for {
		token = p.GetNextToken()
		switch token.(type) {
		case *core.EofToken:
			goto DONE
		}

		tokenType := token.GetType()

		switch tokenType {
		case core.ERROR_TOKEN:
			p.errorHandler.Flag(
				compiler.PARSING,
				token,
				token.GetValue(),
			)
		case tokens.IDENTIFIER:
			tname := strings.ToLower(token.GetText())

			symbol, er := p.symbolTables.Resolve(tname)
			if er != nil {
				symbol, _ = p.symbolTables.EnterLocal(tname)
			}

			symbol.AddLine(&ir.Line{
				Number: token.GetLineNumber(),
			})
		default:
			p.SendMessage(core.NewTokenMessage(token))
		}
	}
	
DONE:
	endTime := time.Since(startTime)


	p.SendMessage(observer.CreateEvent(
		observer.PARSER_SUMMARY,
		&frontend.ParserSummary{
			LineNumber		: token.GetLineNumber(),
			ErrorCount		: p.GetErrorCount(),
			Time			: endTime,
			SymbolTables 	: p.GetSymbolTables(),
		}))

	return nil
}

func (p *RecursiveDescentPascalParser) GetNextToken() core.Token {
	return p.Parser.GetNextToken()
}

func (p *RecursiveDescentPascalParser) Initialize() error {
	p.symbolTables = p.symbolFactory.CreateStack()
	return nil
}




func FromSource(
		compiler compiler.Compiler,
		scanner frontend.Scanner,
		eventDispatcher observer.EventProducer,

) frontend.Parser {
	return NewPascalParser(
		compiler,
		scanner,
		eventDispatcher,
	)
}
