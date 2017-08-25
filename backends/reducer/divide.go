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
   
    lhsfloat := isFloat(lhs)
    rhsfloat := isFloat(rhs)
    
    
    if lhsfloat && rhsfloat {
        return lhs.(float64) / rhs.(float64)
    } else if lhsfloat {
        return lhs.(float64) / float64(rhs.(int64))
    } else if rhsfloat {
        return float64(lhs.(int64)) / rhs.(float64)
    }
    
    return lhs.(int64) / rhs.(int64)
}
