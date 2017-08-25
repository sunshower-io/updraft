package reducer

import (
    "time"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/common/observer"
    "github.com/sunshower-io/updraft/backends/common"
    "fmt"
)


var (
    add = AddOperation{}
    
)



type Reducer struct {
    common.Backend
    
    executionModel      ir.ExecutionModel
    
    symbolTables        ir.SymbolTableStack
    RootOperation       StatementReducer 
    ErrorHandler        common.RuntimeErrorHandler
    
    ErrorCount          uint 
    InstructionCount    uint 
}

func (r *Reducer) GetSymbolTables() ir.SymbolTableStack {
    return r.symbolTables
}


func (r *Reducer) IncrementOperations() {
    r.ErrorCount++
}

func (r *Reducer) ResolveFor(
    parent common.Operation,
    nodeType ir.IntermediateNodeType,
) common.Operation {
    
    switch nodeType {
    case ir.SCOPE:
        return CompoundReducer{parent}
    case ir.ASSIGN:
        return AssignmentReducer{parent}
    case ir.EXPRESSION:
        return ExpressionReducer{parent}
    case ir.INTEGER:
        return PrimitiveReducer{}
    case ir.ADD:
        return &AddOperation{parent}
    case ir.NEGATE:
        return NegationOperation{parent}

    case ir.MULTIPLY:
        return &AddOperation{parent}
    case ir.DIVIDE:
        return &AddOperation{parent}
    case ir.VARIABLE:
        return VariableOperation{parent}
    }
    panic(fmt.Sprintf("No reducer %s", nodeType))
}


func (r *Reducer) Resolve(
    parent common.Operation,
    node ir.IntermediateNode,
) common.Operation {
   
    switch node.GetType() {
    case ir.SCOPE:
        return CompoundReducer{parent}
    case ir.ASSIGN:
        return AssignmentReducer{parent}
    case ir.EXPRESSION:
        return ExpressionReducer{parent}
    case ir.INTEGER:
        return PrimitiveReducer{}
    case ir.NEGATE:
        return NegationOperation{parent}

    case ir.MULTIPLY:
        return &AddOperation{parent}

    case ir.ADD:
        return &AddOperation{parent}
    case ir.DIVIDE:
        return &AddOperation{parent}
    case ir.VARIABLE:
        return VariableOperation{parent}
    }
    panic(fmt.Sprintf("No reducer %s", node.GetType()))
}





func (r *Reducer) Process(
        model ir.ExecutionModel, 
        symboltables ir.SymbolTableStack,
) error {
    
    
    r.symbolTables = symboltables
    r.executionModel = model
    
    startTime := time.Now()
    
    
    statementReducer := StatementReducer{
        ErrorHandler: r.ErrorHandler,
    }
    
    root := model.GetRoot()
    if root != nil {
        statementReducer.Apply(root, r)
    }
    
    
    
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



func NewReductionExecutionEngine(
        executionModel ir.ExecutionModel, 
        stack ir.SymbolTableStack,
) common.Backend  {
    return &Reducer {
        Backend: common.NewBaseBackend(),
        ErrorHandler: common.NewRuntimeErrorHandler(),
        symbolTables: stack,
        executionModel: executionModel,
    }
}
