package reducer

import (
    "github.com/sunshower-io/updraft/common/ir"
    "github.com/sunshower-io/updraft/backends/common"
)

type DivideOperation struct {
    common.Operation 
}

func (o DivideOperation) Apply(
        node ir.IntermediateNode,
        ctx common.OperationContext,
) interface{} {
    reducer := o.Operation
    lhsNode := node.Get(0)
    rhsNode := node.Get(1)
    
    lhs := reducer.Apply(lhsNode, ctx)
    rhs := reducer.Apply(rhsNode, ctx)
    
    
    if isFloat(lhs) || isFloat(rhs) {
        return lhs.(float64) / rhs.(float64)
    }
    
    return lhs.(int64) / rhs.(int64)
}
