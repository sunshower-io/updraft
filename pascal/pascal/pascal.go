package pascal

import (
	gio "io"
	"github.com/sunshower-io/updraft/pascal/front"
	"github.com/sunshower-io/updraft/common/compiler"
	fio "github.com/sunshower-io/updraft/front/parser"
	"github.com/sunshower-io/updraft/common/core"
    "github.com/sunshower-io/updraft/common"
)


type PascalCompiler struct {
	compiler.Compiler
}

func (c *PascalCompiler) Compile() core.CompilationResult {
	result := core.NewCompilationResult()
	c.GetParser().Parse(result)
	return result
}

func NewPascal(
	source gio.Reader,
) compiler.Compiler {

	baseCompiler := new(compiler.AbstractCompiler)

	lexingDispatcher := baseCompiler.GetDispatcher(
		common.LEXING,
	)
	
	
	parsingDispatcher := baseCompiler.GetDispatcher(
		common.PARSING,
	)

	stream := fio.NewSource(
		source,
		lexingDispatcher,
	)
	
	
	scanner := front.NewScanner(
		stream,
		lexingDispatcher,
	)
	
	
	baseCompiler.Source = stream
	
	pascalCompiler := &PascalCompiler{
		Compiler : baseCompiler,
	}
	
	parser := front.FromSource(
		pascalCompiler,
		scanner,
		parsingDispatcher,
	)
	
	
	
	
	
	baseCompiler.Parser = parser
	
	return pascalCompiler
	
	
}



/**

type Pascal struct {
	source      io.Source
	parser      cparser.Parser
	backend     backend.Backend
	symbolTable ir.SymbolTable
	//code        ir.Code

	sourceListeners  []observer.EventListener
	scannerListeners []observer.EventListener
}

func (p *Pascal) AddSourceListener(l observer.EventListener) {
	if p.sourceListeners == nil {
		p.sourceListeners = make([]observer.EventListener, 0)
	}
	p.sourceListeners = append(p.sourceListeners, l)
}

func (p *Pascal) AddScannerListener(l observer.EventListener) {

	if p.scannerListeners == nil {
		p.scannerListeners = make([]observer.EventListener, 0)
	}
	p.scannerListeners = append(p.scannerListeners, l)
}

func (p *Pascal) Run() error {
	p.decorateSource(p.source)
	p.decorateScanner(p.parser)
	err := p.parser.Parse()
	code := p.parser.GetCode()
	symbolTable := p.parser.GetSymbolTable()
	symbolTables := p.parser.GetSymbolTables()
	println(new(ir.CrossReferencer).Print(symbolTables))
	p.backend.Process(code, symbolTable)
	return err
}

func NewPascal(
	source io.Source,
	opts PascalOptions,
) *Pascal {

	parser := front.FromSource(source)

	backend := factory.NewBackend(factory.INTERPRETER)
	backend.AddEventListener(new(
		common.BackendMessageListener,
	))

	p := &Pascal{
		source:  source,
		parser:  parser,
		backend: backend,
	}
	return p
}



func (p *Pascal) decorateSource(source io.Source) {
	source.AddEventListener(new(
		common.SourceMessageListener,
	))

	for _, l := range p.sourceListeners {
		source.AddEventListener(l)
	}
}

func (p *Pascal) decorateScanner(parser cparser.Parser) {

	parser.AddEventListener(new(
		common.ParserMessageListener,
	))

	parser.AddEventListener(new(
		cparser.ParserMessageListener,
	))
	for _, l := range p.scannerListeners {
		parser.GetScanner().AddEventListener(l)
	}
}

*/
