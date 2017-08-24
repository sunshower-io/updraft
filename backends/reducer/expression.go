package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
)

type ExpressionReducer struct {
    StatementReducer
}


func (r ExpressionReducer) Apply(
        node ir.IntermediateNode, 
        ctx ReducerContext,
) interface{} {
    
    reducer := ctx.ResolveFor(r, node.GetType())
    return reducer.Apply(node, ctx)
}
