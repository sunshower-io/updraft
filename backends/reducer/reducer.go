package reducer

import (
    "fmt"
    "time"
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/common/observer"
    "github.com/sunshower-io/updraft/backends/common"
)




var (
    add         = AddOperation{}
    subtract    = SubtractOperation{}
    divide      = DivideOperation{}
    modulo      = ModuloOperation{}
    multiply    = MultiplyOperation{}
    negate      = NegationOperation{}
    or          = OrOperation{}
    not         = NotOperation{}
    and         = AndOperation{}
    xor         = XorOperation{}
    compound    = CompoundReducer{}
    eq          = EqualityReducer{}
    statement   = StatementReducer{}
    assignment  = AssignmentReducer{}
    expression  = ExpressionReducer{}
    primitive   = PrimitiveReducer{}
    noop        = NoOp{}
    variable    = VariableOperation{}
    neq         = InequalityReducer{}
)

func initialize(op common.Operation) {
    add.Operation           = op
    subtract.Operation      = op
    divide.Operation        = op
    modulo.Operation        = op
    multiply.Operation      = op
    negate.Operation        = op
    or.Operation            = op
    not.Operation           = op
    and.Operation           = op
    xor.Operation           = op
    assignment.Operation    = op
    expression.Operation    = op
    compound.Operation      = op
    statement.Operation     = op
    variable.Operation      = op
    eq.Operation            = op
    neq.Operation           = op
}

func resolve(
        nodeType ir.IntermediateNodeType,
) common.Operation {
    
    switch nodeType {
    case ir.SCOPE:
        return compound
    case ir.ASSIGN:
        return assignment
    case ir.EXPRESSION:
        return expression
    case ir.INTEGER, ir.FLOAT, ir.BOOLEAN:
        return primitive 
    case ir.ADD:
        return add 
    case ir.NEGATE:
        return negate 
    case ir.MULTIPLY:
        return multiply 
    case ir.DIVIDE, ir.FLOAT_DIVIDE, ir.INTEGER_DIVIDE:
        return divide 
    case ir.VARIABLE:
        return variable 
    case ir.NO_OP:
        return noop
    case ir.SUBTRACT:
        return subtract
    case ir.MODULO:
        return modulo
    case ir.OR:
        return or
    case ir.AND:
        return and
    case ir.NOT:
        return not
    case ir.EQUAL_TO:
        return eq
    case ir.NOT_EQUAL_TO:
        return neq
    }
    panic(fmt.Sprintf("No reducer %s", nodeType))
    
}

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
    return resolve(nodeType)
}


func (r *Reducer) Resolve(
    parent common.Operation,
    node ir.IntermediateNode,
) common.Operation {
    return resolve(node.GetType())
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
    initialize(&statement)
    
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
