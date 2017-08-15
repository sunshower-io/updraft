package front

import (
    "time"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/common/observer"
    "github.com/sunshower-io/updraft/middle/core"
    "github.com/sunshower-io/updraft/pascal/tokens"
    "github.com/sunshower-io/updraft/common/frontend"
    "github.com/sunshower-io/updraft/common/compiler"
    ccore "github.com/sunshower-io/updraft/common/core"
    "github.com/sunshower-io/updraft/pascal/elements"
    "github.com/sunshower-io/updraft/common/errors"
    "github.com/sunshower-io/updraft/common"
)

type RecursiveDescentPascalParser struct {
    frontend.Parser
    
    errorHandler          errors.ErrorHandler
    symbolTables          ir.SymbolTableStack
    symbolFactory         ir.SymbolTableFactory
    executionModelFactory ir.ExecutionModelFactory
}

func (p *RecursiveDescentPascalParser) GetErrorHandler() errors.ErrorHandler {
    return p.errorHandler
}

func (p *RecursiveDescentPascalParser) GetSymbolTables() ir.SymbolTableStack {
    return p.symbolTables
}

func (p *RecursiveDescentPascalParser) GetExecutionModelFactory() ir.ExecutionModelFactory {
    return p.executionModelFactory
}

func (p *RecursiveDescentPascalParser) Parse() ccore.CompilationResult {
    
    startTime := time.Now()
    
    var (
        token          core.Token
        root           ir.IntermediateNode
        executionModel ir.ExecutionModel
    )
    executionModel = p.executionModelFactory.NewExecutionModel()
    
    token = p.GetNextToken()
    
    tokenType := token.GetType()
    
    switch tokenType {
    case core.ERROR_TOKEN:
        p.errorHandler.Flag(
            common.PARSING,
            token,
            token.GetValue(),
        )
    case tokens.BEGIN:
        statement := elements.NewStatementParser(p)
        root, _ = statement.Parse(token)
        executionModel.SetRoot(root)
        //tname := strings.ToLower(token.GetText())
        //
        //symbol, er := p.symbolTables.Resolve(tname)
        //if er != nil {
        //	symbol, _ = p.symbolTables.EnterLocal(tname)
        //}
        //
        //symbol.AddLine(&ir.Line{
        //	Number: token.GetLineNumber(),
        //})
    default:
        p.SendMessage(core.NewTokenMessage(token))
    }

    endTime := time.Since(startTime)
    
    executionModel.SetRoot(root)
    
    p.SendMessage(observer.CreateEvent(
        observer.PARSER_SUMMARY,
        &frontend.ParserSummary{
            LineNumber:   token.GetLineNumber(),
            ErrorCount:   p.GetErrorCount(),
            Time:         endTime,
            SymbolTables: p.GetSymbolTables(),
        }))
   
   
    return ccore.NewCompilationResult(executionModel)
    
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
