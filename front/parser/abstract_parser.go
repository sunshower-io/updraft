package parser

import (
	"github.com/sunshower-io/updraft/common/ir"
	"github.com/sunshower-io/updraft/common/observer"
	core "github.com/sunshower-io/updraft/middle/core"
	"github.com/sunshower-io/updraft/common/frontend"
)

type AbstractParser struct {

	/**

	 */
	frontend.Parser

	/**

	 */
	observer.EventProducer

	/**
	  Scanner
	*/

	scanner frontend.Scanner

	/**
	  Core
	*/

	code *core.Code

	/**

	 */
	symbolTable ir.SymbolTable
}

func (f *AbstractParser) GetNextToken() core.Token {
	t, er := f.scanner.NextToken()
	if er != nil {
		panic(er)
	}
	
	return t
}

func (f *AbstractParser) SendMessage(m observer.Message) {
	f.EventProducer.SendMessage(m)
}

func (f *AbstractParser) RemoveMessageListener(l observer.EventListener) {
	f.EventProducer.RemoveMessageListener(l)
}

func (f *AbstractParser) AddEventListener(l observer.EventListener) {
	f.EventProducer.AddEventListener(l)
}

func (f *AbstractParser) GetErrorCount() int {
	return 0
}

func (f *AbstractParser) GetCode() core.Code {
	return f.code
}

func (f *AbstractParser) GetSymbolTable() ir.SymbolTable {
	return f.symbolTable
}

func (f *AbstractParser) GetScanner() frontend.Scanner {
	return f.scanner
}

func NewAbstractParser(
		scanner frontend.Scanner,
		producer observer.EventProducer,
) frontend.Parser {
	parser := new(AbstractParser)
	parser.scanner = scanner
	parser.EventProducer = producer
	return parser
}
