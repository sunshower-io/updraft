package reducer



import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type MultiplyOperation struct {
    StatementReducer
}

func (o MultiplyOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    reducer := o.StatementReducer
    
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx)
    rhs := reducer.Apply(rhsNode, ctx)
    
    
    if isFloat(lhs) || isFloat(rhs) {
        return lhs.(float64) * rhs.(float64)
    }
    
    return lhs.(int64) * rhs.(int64)
}
