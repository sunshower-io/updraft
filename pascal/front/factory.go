package front

import (
	"gitlab.com/sunshower.io/updraft/common/ir"
	"gitlab.com/sunshower.io/updraft/common/frontend"
	"gitlab.com/sunshower.io/updraft/front/parser"
	"gitlab.com/sunshower.io/updraft/common/observer"
	"gitlab.com/sunshower.io/updraft/common/compiler"
)

func NewPascalParser(
		compiler 		compiler.Compiler,
		scanner 		frontend.Scanner,
		eventProducer 	observer.EventProducer,
		
) frontend.Parser {
	p := &RecursiveDescentPascalParser{
		errorHandler: &PascalErrorHandler{
			Compiler: compiler,
		},
		Parser: parser.NewAbstractParser(
			&PascalScanner{
				Scanner: scanner,
			},
			eventProducer,
		),
	}
	p.symbolFactory = ir.DefaultSymbolTableFactory
	p.Initialize()
	return p
}
