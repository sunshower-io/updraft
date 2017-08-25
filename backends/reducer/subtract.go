package reducer

import (
    "github.com/sunshower-io/updraft/backends/common"
    "github.com/sunshower-io/updraft/common/ir"
)

type SubtractOperation struct {
    common.Operation
}

func (o SubtractOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    reducer := o.Operation
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx)
    rhs := reducer.Apply(rhsNode, ctx)
    
    
    if isFloat(lhs) || isFloat(rhs) {
        return lhs.(float64) - rhs.(float64)
    }
    
    return lhs.(int64) - rhs.(int64)
}
