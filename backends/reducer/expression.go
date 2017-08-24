package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type ExpressionReducer struct {
    StatementReducer
}


func (r ExpressionReducer) Apply(
        node ir.IntermediateNode, 
        ctx common.OperationContext,
) interface{} {
    
    reducer := ctx.ResolveFor(r, node.GetType())
    return reducer.Apply(node, ctx)
}
