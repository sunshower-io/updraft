package reducer

import (
    "time"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/common/observer"
    "github.com/sunshower-io/updraft/common/backend"
    "github.com/sunshower-io/updraft/backends/common"
)


var (
    add = AddOperation{}
    
)

type ReducerContext struct {
    common.OperationContext
} 


func (r *ReducerContext) ResolveFor(
        operation backend.Operation, 
        nodeType ir.IntermediateNodeType,
) backend.Operation {
    
    
    
    switch nodeType {
    case ir.ADD:
        return add
    }
    
    return nil
}


func (r *ReducerContext) Resolve(
        operation backend.Operation, 
        node ir.IntermediateNode,
) backend.Operation {
    return nil 
}


type Reducer struct {
    backend.Backend
    Parent              backend.Operation
    ErrorHandler        backend.RuntimeErrorHandler
    
    
    
    executionModel      ir.ExecutionModel
    
    symbolTables        ir.SymbolTableStack
    
    ErrorCount          uint 
    InstructionCount    uint 
}



func (r *Reducer) Process(
        model ir.ExecutionModel, 
        symboltables ir.SymbolTableStack,
) error {
    
    
    r.symbolTables = symboltables
    r.executionModel = model
    
    startTime := time.Now()
    
    
    statementReducer := StatementReducer{}
    
    root := model.GetRoot()
    
    statementReducer.Apply(root)
    
    
    
    r.SendMessage(&observer.BaseEvent{
        Topic: observer.INTERPRETER_SUMMARY,
        Body: common.Summary {
            OperationCount: r.InstructionCount,
            ErrorCount: r.ErrorCount,
            ElapsedTime: time.Since(startTime),
        },
    })
    
    return nil
}



