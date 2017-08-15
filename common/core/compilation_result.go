package core

import "github.com/sunshower-io/updraft/common/ir"

type CompilationResult interface {
    
    GetExecutionModel() ir.ExecutionModel
}


type AbstractCompilationResult struct {
    CompilationResult
    ExecutionModel ir.ExecutionModel
}

func (r *AbstractCompilationResult) GetExecutionModel() ir.ExecutionModel {
    return r.ExecutionModel
}

func NewCompilationResult(
        model ir.ExecutionModel,
) CompilationResult {
    return &AbstractCompilationResult{
        ExecutionModel: model,
    }
}