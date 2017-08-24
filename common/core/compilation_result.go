package core

import "github.com/sunshower-io/updraft/common/ir"

type CompilationResult interface {
    
    GetSymbolTables()   ir.SymbolTableStack
    GetExecutionModel() ir.ExecutionModel
}


type AbstractCompilationResult struct {
    CompilationResult
    ExecutionModel  ir.ExecutionModel
    SymbolTables    ir.SymbolTableStack
}

func (r *AbstractCompilationResult) GetExecutionModel() ir.ExecutionModel {
    return r.ExecutionModel
}

func (r *AbstractCompilationResult) GetSymbolTables() ir.SymbolTableStack {
    return r.SymbolTables
}

func NewCompilationResult(
        model ir.ExecutionModel,
        symbolTables ir.SymbolTableStack,
) CompilationResult {
    return &AbstractCompilationResult{
        ExecutionModel: model,
        SymbolTables: symbolTables,
    }
}