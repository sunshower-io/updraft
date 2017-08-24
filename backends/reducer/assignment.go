package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type AssignmentReducer struct {
    StatementReducer
}

func (a AssignmentReducer) Apply(
        node ir.IntermediateNode,
        ctx  common.OperationContext,
) interface{} {
    
    children := node.GetChildren()
    target := children[0]
    expression := children[1]
    
    expressionReducer := ctx.ResolveFor(a, ir.EXPRESSION)
    
    value := expressionReducer.Apply(expression, ctx)
    
    symbol := target.GetAttribute(ir.ID).(ir.Symbol)
    symbol.SetAttribute(ir.DATA_VALUE, value)
    ctx.IncrementOperations()
    return nil
    
}
