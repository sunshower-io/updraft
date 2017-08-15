package frontend

import (
    "time"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/common/observer"
    core "github.com/sunshower-io/updraft/middle/core"
    ccore "github.com/sunshower-io/updraft/common/core"
    "github.com/sunshower-io/updraft/common/errors"
)

type ParserSummary struct {
    LineNumber      int
    ErrorCount      int
    Time            time.Duration
    SymbolTables    ir.SymbolTableStack
}





type Parser interface {
    observer.EventProducer
    
    
    GetScanner()                        Scanner
   
    
    GetCode()                           core.Code
    
    GetSymbolTable()                    ir.SymbolTable
    
    GetSymbolTables()                   ir.SymbolTableStack
    
    GetErrorCount()                     int
    
    GetErrorHandler()                   errors.ErrorHandler
    
    GetNextToken()                      core.Token
    
    GetCurrentToken()                   core.Token
    
    Initialize()                        error
    
    Parse()                             ccore.CompilationResult  
    
    
    GetExecutionModelFactory()          ir.ExecutionModelFactory
}
