package front

import (
	"github.com/sunshower-io/updraft/common/ir"
	"github.com/sunshower-io/updraft/common/frontend"
	"github.com/sunshower-io/updraft/front/parser"
	"github.com/sunshower-io/updraft/common/observer"
	"github.com/sunshower-io/updraft/common/compiler"
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
