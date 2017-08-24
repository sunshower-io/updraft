package reducer

import (
    "time"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/common/observer"
    "github.com/sunshower-io/updraft/backends/common"
)


var (
    add = AddOperation{}
    
)



type Reducer struct {
    common.Backend
    Parent              common.Operation
    ErrorHandler        common.RuntimeErrorHandler
    
    
    
    ctx                 common.OperationContext
    executionModel      ir.ExecutionModel
    
    symbolTables        ir.SymbolTableStack
    
    ErrorCount          uint 
    InstructionCount    uint 
}


func (r *Reducer) IncrementOperations() {
    r.ErrorCount++
}

func (r *Reducer) ResolveFor(
    common.Operation,
    ir.IntermediateNodeType,
) common.Operation {
    return nil
}


func (r *Reducer) Resolve(
    common.Operation,
    ir.IntermediateNode,
) common.Operation {
    return nil
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
    
    statementReducer.Apply(root, r)
    
    
    
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



